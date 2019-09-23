package grpc

import (
	"context"
	pb "github.com/sslab-instapay/instapay-go-client/proto"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
	"github.com/sslab-instapay/instapay-go-client/model"
)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *pb.AgreeRequestsMessage) (*pb.Result, error) {
	// 동의한다는 메시지를 전달
	channelIds := in.GetChannelIds()

	for channelId := range channelIds.ChannelIds {
		channel, err := repository.GetChannelById(int64(channelId))

		// update 채널 status 및 locked Balance
		channel.Status = model.PRE_UPDATE
		channel.LockedBalance = in.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &pb.Result{}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *pb.UpdateRequestsMessage) (*pb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelIds := in.GetChannelIds()

	for channelId := range channelIds.ChannelIds {

		channel, err := repository.GetChannelById(int64(channelId))

		// update 채널 status 및 locked Balance
		channel.Status = model.POST_UPDATE
		channel.MyBalance += in.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &pb.Result{}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *pb.ConfirmRequestsMessage) (*pb.Result, error) {
	//페이먼트를 ~~
	return &pb.Result{}, nil
}
