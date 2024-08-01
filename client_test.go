package goao

import (
	"testing"

	"github.com/permadao/goar"
)

var tClient *Client

func init() {
	signer, err := goar.NewSignerFromPath("testKey.json")
	if err != nil {
		panic(err)
	}
	bundler, err := goar.NewBundler(signer)
	if err != nil {
		panic(err)
	}
	tClient = NewClient(
		"https://mu.ao-testnet.xyz",
		"https://cu.ao-testnet.xyz",
		bundler)
}

func TestSend(t *testing.T) {
	// res, err := tClient.Send(
	// 	"ya9XinY0qXeYyf7HXANqzOiKns8yiXZoDtFqUMXkX0Q",
	// 	"",
	// 	[]schema.Tag{
	// 		schema.Tag{Name: "Action", Value: "Claim"},
	// 	})
	// assert.NoError(t, err)
	// fmt.Println(res)
}

func TestEval(t *testing.T) {
	// res, err := tClient.Eval(
	// 	"ya9XinY0qXeYyf7HXANqzOiKns8yiXZoDtFqUMXkX0Q",
	// 	"a=1")
	// assert.NoError(t, err)
	// fmt.Println(res)
}