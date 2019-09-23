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
	channelPayments := in.ChannelPayments

	for _, channelPayment := range channelPayments.ChannelPayments {

		channel, err := repository.GetChannelById(channelPayment.ChannelId)

		// update 채널 status 및 locked Balance
		channel.Status = model.PRE_UPDATE
		channel.LockedBalance += in.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Fatal(err)
		}
		// PaymentData 삽입
		repository.InsertPaymentData(model.PaymentData{PaymentNumber: in.PaymentNumber, ChannelId: channelPayment.ChannelId, Amount: in.Amount})
	}
	return &pb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *pb.UpdateRequestsMessage) (*pb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelPayments := in.ChannelPayments

	for _, channelPayment := range channelPayments.ChannelPayments {

		channel, err := repository.GetChannelById(channelPayment.ChannelId)
		channel.Status = model.POST_UPDATE
		channel.MyBalance += in.Amount
		channel.LockedBalance -= in.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Println("Something is wrong")
			return &pb.Result{}, err
		}
	}
	return &pb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *pb.ConfirmRequestsMessage) (*pb.Result, error) {
	paymentDatas, err := repository.GetPaymentDatasByPaymentId(in.PaymentNumber)
	if err != nil{
		return &pb.Result{}, err
	}

	for _, paymentData := range paymentDatas{
		channel, err := repository.GetChannelById(paymentData.ChannelId)
		if err != nil{
			return &pb.Result{}, err
		}

		channel.Status = model.IDLE
		repository.UpdateChannel(channel)
	}

	return &pb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}
