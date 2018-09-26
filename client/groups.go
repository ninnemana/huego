package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/huego"
	"github.com/pkg/errors"
)

func (c *client) AllGroups(ctx context.Context) ([]interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.groups.all")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s/groups", host, user)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	groups := make(map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		return nil, err
	}

	results := []interface{}{}
	for _, l := range groups {
		lmp, ok := l.(map[string]interface{})
		if !ok {
			continue
		}

		results = append(results, lmp)
	}

	return results, nil
}

func (c *client) CreateGroup(ctx context.Context, group interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetGroup(ctx context.Context, id string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SaveGroup(ctx context.Context, id string, group interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SetGroupState(ctx context.Context, id string, state interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteGroup(ctx context.Context, id string) error {
	return hue.ErrNotImplemented
}
