package client

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/trace"
)

func Test_client_AllGroups(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				trace: tt.fields.trace,
			}
			got, err := c.AllGroups(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AllGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.AllGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CreateGroup(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx   context.Context
		group interface{}
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
			got, err := c.CreateGroup(tt.args.ctx, tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CreateGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetGroup(t *testing.T) {
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
			got, err := c.GetGroup(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SaveGroup(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx   context.Context
		id    string
		group interface{}
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
			got, err := c.SaveGroup(tt.args.ctx, tt.args.id, tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SaveGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.SaveGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SetGroupState(t *testing.T) {
	type fields struct {
		trace *trace.Client
	}
	type args struct {
		ctx   context.Context
		id    string
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
			got, err := c.SetGroupState(tt.args.ctx, tt.args.id, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SetGroupState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.SetGroupState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_DeleteGroup(t *testing.T) {
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
			if err := c.DeleteGroup(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("client.DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
