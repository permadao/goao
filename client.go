package goao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/permadao/goao/schema"
	"github.com/permadao/goar"
	goarSchema "github.com/permadao/goar/schema"
	"gopkg.in/h2non/gentleman.v2"
)

type Client struct {
	muCli *gentleman.Client
	cuCli *gentleman.Client

	bundler *goar.Bundler
}

func NewClient(muURL, cuURL string, signer interface{}) (*Client, error) {
	muCli := gentleman.New().URL(muURL)
	cuCli := gentleman.New().URL(cuURL)

	bundler, err := goar.NewBundler(signer)
	if err != nil {
		return nil, err
	}

	return &Client{
		muCli: muCli,
		cuCli: cuCli,

		bundler: bundler,
	}, nil
}

func (c *Client) Send(processId, data, msgType string, tags []goarSchema.Tag) (res schema.ResponseMu, err error) {
	tags = append(tags, goarSchema.Tag{Name: "Data-Protocol", Value: schema.DataProtocol})
	tags = append(tags, goarSchema.Tag{Name: "Variant", Value: schema.Variant})
	tags = append(tags, goarSchema.Tag{Name: "Type", Value: msgType})
	tags = append(tags, goarSchema.Tag{Name: "SDK", Value: schema.SDK})

	item, err := c.bundler.CreateAndSignItem([]byte(data), processId, "", tags)
	if err != nil {
		return
	}

	req := c.muCli.Post()
	req.SetHeader("content-type", "application/octet-stream")
	req.SetHeader("accept", "application/json")
	req.Body(bytes.NewBuffer(item.Binary))

	resp, err := req.Send()
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("invalid server response: %d", resp.StatusCode)
		return
	}

	err = json.Unmarshal(resp.Bytes(), &res)
	return
}

func (c *Client) Eval(processId, code string) (res schema.ResponseMu, err error) {
	return c.Send(
		processId, code, schema.TypeMessage,
		[]goarSchema.Tag{
			goarSchema.Tag{Name: "Action", Value: "Eval"},
		},
	)
}

func (c *Client) Result(processId, messageId string) (res schema.ResponseCu, err error) {
	req := c.cuCli.Get()
	req.AddPath(fmt.Sprintf("/result/%v", messageId))
	req.AddQuery("process-id", processId)

	resp, err := req.Send()
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("invalid server response: %d", resp.StatusCode)
		return
	}

	// golang http not handle Temporary Redirect(307)
	if resp.StatusCode == http.StatusTemporaryRedirect {
		loc := resp.Header.Get("Location")
		resp.Close()
		resp, err = c.cuCli.Request().URL(loc).Send()
		if err != nil {
			return
		}
		if !resp.Ok {
			err = fmt.Errorf("invalid server response: %d", resp.StatusCode)
			return
		}
	}

	defer resp.Close()
	err = json.Unmarshal(resp.Bytes(), &res)
	return
}

func (c *Client) Spawn(processName, appName, module, scheduler string) (res schema.ResponseMu, err error) {
	return c.Send(
		"", strconv.Itoa(int(time.Now().UnixNano())), schema.TypeProcess,
		[]goarSchema.Tag{
			{Name: "Name", Value: processName},
			{Name: "App-Name", Value: appName},
			{Name: "Module", Value: module},
			{Name: "Scheduler", Value: scheduler},
		},
	)
}

func (c *Client) DryRun(processId, data string, tags []goarSchema.Tag) (res schema.ResponseCu, err error) {
	tags = append(tags, goarSchema.Tag{Name: "Data-Protocol", Value: schema.DataProtocol})
	tags = append(tags, goarSchema.Tag{Name: "Variant", Value: schema.Variant})
	tags = append(tags, goarSchema.Tag{Name: "Type", Value: schema.TypeMessage})
	tags = append(tags, goarSchema.Tag{Name: "SDK", Value: schema.SDK})

	item := struct {
		Id     string           `json:"Id"`
		Target string           `json:"Target"`
		Owner  string           `json:"Owner"`
		Data   string           `json:"Data"`
		Tags   []goarSchema.Tag `json:"Tags"`
		Anchor string           `json:"Anchor"`
	}{
		Id:     "0000000000000000000000000000000000000000001",
		Target: processId,
		Owner:  "0000000000000000000000000000000000000000001",
		Data:   data,
		Tags:   tags,
	}
	payload, err := json.Marshal(item)
	if err != nil {
		return
	}

	req := c.cuCli.Post()
	req.AddPath("/dry-run")
	req.AddQuery("process-id", processId)
	req.JSON(payload)
	resp, err := req.Send()
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("invalid server response: %d", resp.StatusCode)
		return
	}

	if resp.StatusCode == http.StatusTemporaryRedirect {
		redirectURL := resp.Header.Get("Location")
		resp.Close()
		req = gentleman.New().URL(redirectURL).Post()
		req.JSON(payload)
		resp, err = req.Send()
		if err != nil {
			return
		}
		if !resp.Ok {
			err = fmt.Errorf("invalid server response: %d", resp.StatusCode)
			return
		}
	}

	defer resp.Close()
	err = json.Unmarshal(resp.Bytes(), &res)
	return
}
