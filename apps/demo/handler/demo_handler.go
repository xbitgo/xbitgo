package handler

import (
	"context"
	pb "xbitgo/proto/apps/demo"
)

// DemoHandlerImpl @IMPL[Demo]
type DemoHandlerImpl struct {
	pb.UnimplementedDemoServer
}

func NewDemoHandlerImpl() *DemoHandlerImpl {
	return &DemoHandlerImpl{}
}

func (impl *DemoHandlerImpl) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Msg: "Hello XBitGO",
	}, nil
}
