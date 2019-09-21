package grpc

import (
	"context"
	pb "github.com/sslab-instapay/instapay-go-client/proto"
)

type grpcServer struct {
}

func (s *grpcServer) AgreementRequest(ctx context.Context, in *pb.AgreeRequestsMessage) (*pb.Result, error) {
	// 동의한다는 메시지를 전달
	return &pb.Result{}, nil
}

func (s *grpcServer) UpdateRequest(ctx context.Context, in *pb.UpdateRequestsMessage) (*pb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	return &pb.Result{}, nil
}

func (s *grpcServer) ConfirmPayment(ctx context.Context, in *pb.ConfirmRequestsMessage) (*pb.Result, error) {
	//페이먼트를 ~~
	return &pb.Result{}, nil
}
