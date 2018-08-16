package client

import (
	"context"
	"log"
	"testing"

	hue "github.com/ninnemana/huego"
	"github.com/ninnemana/huego/client"
)

var (
	cl hue.Client
)

func TestMain(m *testing.M) {

	var err error
	cl, err = client.New()
	if err != nil {
		log.Fatalf("failed to create new hue.Client: %v", err)
	}

	m.Run()
}

func TestAllBridges(t *testing.T) {
	ctx := context.Background()

	results, err := cl.AllBridges(ctx, &hue.AllBridgeParams{})
	if err != nil {
		t.Errorf("failed to get all briges: %v", err)
		return
	}

	for _, res := range results {
		bridge, ok := res.(*hue.Bridge)
		if !ok {
			t.Errorf("expected '*hue.Bridge' got '%T'", res)
			continue
		}

		t.Log(bridge.GetName())
	}
}
