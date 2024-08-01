package goao

import (
	"bytes"
	"encoding/json"
	"fmt"

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

func NewClient(muURL, cuURL string, bundler *goar.Bundler) *Client {
	return &Client{
		muCli: gentleman.New().URL(muURL),
		cuCli: gentleman.New().URL(cuURL),

		bundler: bundler,
	}
}

func (c *Client) Send(processId, data string, tags []goarSchema.Tag) (res schema.ResponseMu, err error) {
	tags = append(tags, goarSchema.Tag{Name: "Data-Protocol", Value: schema.DataProtocol})
	tags = append(tags, goarSchema.Tag{Name: "Variant", Value: schema.Variant})
	tags = append(tags, goarSchema.Tag{Name: "Type", Value: schema.TypeMessage})
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
	defer resp.Close()

	if resp.StatusCode != 202 {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	err = json.Unmarshal(resp.Bytes(), &res)
	return
}

func (c *Client) Eval(processId, code string) (res schema.ResponseMu, err error) {
	return c.Send(
		processId, code,
		[]goarSchema.Tag{
			goarSchema.Tag{Name: "Action", Value: "Eval"},
		},
	)
}
