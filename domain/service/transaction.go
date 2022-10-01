package service

import "hmr_grpc_blockchain/domain/model"

func InitTransaction(sender, receiver string, value float64) *model.Transaction {
	return &model.Transaction{
		SenderAddress:   sender,
		ReceiverAddress: receiver,
		Value:           value,
	}
}
