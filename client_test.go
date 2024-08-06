package goao

import (
	"fmt"
	"github.com/everFinance/goether"
	goarSchema "github.com/permadao/goar/schema"
	"testing"

	"github.com/permadao/goar"
	"github.com/stretchr/testify/assert"
)

var (
	tClient   *Client
	eccClient *Client
)

func init() {
	signer, err := goar.NewSignerFromPath("testKey.json")
	if err != nil {
		panic(err)
	}

	tClient, err = NewClient(
		"https://mu.ao-testnet.xyz",
		"https://cu.ao-testnet.xyz",
		bundler)

	eccSigner, err := goether.NewSigner("4c3f9a1e5b234ce8f1ab58d82f849c0f70a4d5ceaf2b6e2d9a6c58b1f897ef0a")
	if err != nil {
		panic(err)
	}
	eccBundler, _ := goar.NewBundler(eccSigner)
	eccClient = NewClient(
		"https://mu.ao-testnet.xyz",
		"https://cu.ao-testnet.xyz",
		eccBundler)
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
	// 	"1+1")
	// assert.NoError(t, err)
	// fmt.Println(res)
}

func TestSpawn(t *testing.T) {
	// res, err := tClient.Spawn(
	// 	"test1", "goao-test",
	// 	schema.DefaultModule, schema.DefaultScheduler)
	// assert.NoError(t, err)
	// t.Log(res)
	//
	// // 0x address
	// res, err = eccClient.Spawn(
	// 	"test2", "goao-test2",
	// 	schema.DefaultModule, schema.DefaultScheduler)
	// assert.NoError(t, err)
	// t.Log(res)
}

func TestResult(t *testing.T) {
	res, err := tClient.Result("ya9XinY0qXeYyf7HXANqzOiKns8yiXZoDtFqUMXkX0Q", "5JtjkYy1hk0Zce5mP6gDWIOdt9rCSQAFX-K9jZnqniw")
	assert.NoError(t, err)
	fmt.Println(res)
}

func TestClient_DryRun(t *testing.T) {
	processId := "xU9zFkq3X2ZQ6olwNVvr1vUWIjc3kXTWr7xKQD6dh10"
	data := ""
	tags := []goarSchema.Tag{
		{Name: "Action", Value: "Info"},
	}
	res, err := tClient.DryRun(processId, data, tags)
	assert.NoError(t, err)
	t.Log(res)
}
