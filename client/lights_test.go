package client

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/ninnemana/huego"

	"cloud.google.com/go/trace"
)

func Test_client_AllLights(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.HostKey{}, "http://192.168.86.133"),
					hue.UserKey{},
					os.Getenv("HUE_USER"),
				),
			},
		},
		{
			name: "no host",
			args: args{
				ctx: context.WithValue(
					context.Background(),
					hue.UserKey{},
					os.Getenv("HUE_USER"),
				),
			},
			wantErr: true,
		},
		{
			name: "no user",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.HostKey{}, "http://192.168.86.133"),
					hue.UserKey{},
					"",
				),
			},
			wantErr: true,
		},
		{
			name: "invalid host",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.HostKey{}, "http://192.168.86.13"),
					hue.UserKey{},
					os.Getenv("HUE_USER"),
				),
			},
			wantErr: true,
		},
		{
			name: "invalid user",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.HostKey{}, "http://192.168.86.133"),
					hue.UserKey{},
					"user",
				),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{}
			_, err := c.AllLights(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AllLights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_NewLights(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.UserKey{}, os.Getenv("HUE_USER")),
					hue.HostKey{},
					"http://192.168.86.133",
				),
			},
		},
		{
			name: "invalid host",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.UserKey{}, os.Getenv("HUE_USER")),
					hue.HostKey{},
					"http://192.168.86.13",
				),
			},
			wantErr: true,
		},
		{
			name: "invalid user",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.UserKey{}, "HUE_USER"),
					hue.HostKey{},
					"http://192.168.86.133",
				),
			},
			wantErr: true,
		},
		{
			name: "no user",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.UserKey{}, ""),
					hue.HostKey{},
					"http://192.168.86.133",
				),
			},
			wantErr: true,
		},
		{
			name: "no host",
			args: args{
				ctx: context.WithValue(
					context.WithValue(context.Background(), hue.UserKey{}, os.Getenv("HUE_USER")),
					hue.HostKey{},
					"",
				),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{}
			_, err := c.NewLights(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.NewLights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("client.NewLights() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_client_SearchLights(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx       context.Context
		deviceIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			if err := c.SearchLights(tt.args.ctx, tt.args.deviceIDs); (err != nil) != tt.wantErr {
				t.Errorf("client.SearchLights() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetLight(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			got, err := c.GetLight(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetLight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetLight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_RenameLight(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx  context.Context
		id   string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			got, err := c.RenameLight(tt.args.ctx, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RenameLight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.RenameLight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_LightState(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx   context.Context
		id    int
		state interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			got, err := c.LightState(tt.args.ctx, tt.args.id, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LightState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.LightState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Toggle(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			got, err := c.Toggle(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Toggle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_DeleteLight(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			if err := c.DeleteLight(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("client.DeleteLight() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
