package grpc

import (
	"context"
	clientPb "github.com/sslab-instapay/instapay-go-client/proto/client"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
	"github.com/sslab-instapay/instapay-go-client/model"
)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *clientPb.AgreeRequestsMessage) (*clientPb.Result, error) {
	// 동의한다는 메시지를 전달
	channelPayments := in.ChannelPayments

	for _, channelPayment := range channelPayments.ChannelPayments {

		channel, err := repository.GetChannelById(channelPayment.ChannelId)

		// update 채널 status 및 locked Balance
		channel.Status = model.PRE_UPDATE
		channel.LockedBalance -= channelPayment.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Println(err)
		}
		// PaymentData 삽입
		_, err = repository.InsertPaymentData(model.PaymentData{PaymentNumber: in.PaymentNumber, ChannelId: channelPayment.ChannelId, Amount: channelPayment.Amount})
		if err != nil{
			log.Println(err)
		}
	}
	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *clientPb.UpdateRequestsMessage) (*clientPb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelPayments := in.ChannelPayments

	for _, channelPayment := range channelPayments.ChannelPayments {
		// 페이먼트 데이터 삽입
		if result, _ := repository.FindPaymentData(model.PaymentData{ PaymentNumber: in.PaymentNumber, ChannelId: channelPayment.ChannelId, Amount: channelPayment.Amount}); !result{
			repository.InsertPaymentData(model.PaymentData{ PaymentNumber: in.PaymentNumber, ChannelId: channelPayment.ChannelId, Amount: channelPayment.Amount})
		}
		channel, err := repository.GetChannelById(channelPayment.ChannelId)
		channel.Status = model.POST_UPDATE
		channel.MyBalance += channelPayment.Amount
		channel.LockedBalance += channelPayment.Amount

		_, err = repository.UpdateChannel(channel)
		if err != nil {
			log.Println("Something is wrong")
			return &clientPb.Result{}, err
		}
	}
	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.Result, error) {
	log.Println("----ConfirmPayment Request Receive----")
	paymentDatas, err := repository.GetPaymentDatasByPaymentNumber(in.PaymentNumber)
	if err != nil {
		return &clientPb.Result{}, err
	}

	for _, paymentData := range paymentDatas {
		channel, err := repository.GetChannelById(paymentData.ChannelId)
		if err != nil {
			log.Println("Error! ConfirmPayment")
			return &clientPb.Result{}, err
		}

		channel.Status = model.IDLE
		_, err = repository.UpdateChannel(channel)
		if err != nil{
			log.Println("Error! ConfirmPayment")
			return &clientPb.Result{}, err
		}
	}
	log.Println("----ConfirmPayment Request End----")

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}
