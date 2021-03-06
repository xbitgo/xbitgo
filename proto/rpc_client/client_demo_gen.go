// Code generated by xbit. DO NOT EDIT.
package rpc_client

import (
	"context"
	"time"

	//"google.golang.org/grpc/metadata"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/xbitgo/components/prometheus"

	"xbitgo/proto/apps/demo"
)

var Demo = RPCService{Name: "demo"}

func GetDemoClient() *DemoClient {
	conn, ok := RpcConn[Demo.Name]
	if !ok {
		panic(errors.Errorf("No Register RPC Client[%s]", Demo.Name))
	}
	return NewDemoClient(conn)
}

type DemoClient struct {
	cli demo.DemoClient
}

func NewDemoClient(cc *grpc.ClientConn) *DemoClient {
	return &DemoClient{
		cli: demo.NewDemoClient(cc),
	}
}

func (s *DemoClient) Test(ctx context.Context, req *demo.TestRequest, timeout ...time.Duration) (*demo.TestResponse, error) {
	_timeout := 2 * time.Second
	if len(timeout) > 0 {
		_timeout = timeout[0]
	}
	ctx, cancel := context.WithTimeout(ctx, _timeout)
	defer cancel()
	st := time.Now()
	resp, err := s.cli.Test(ctx, req)
	prometheus.HistogramVec.Timing("rpc_client_Demo_Test", []string{"from", From, "ret", prometheus.RetLabel(err)}, st)
	return resp, err
}
