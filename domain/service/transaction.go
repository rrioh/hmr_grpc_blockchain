package service

import "hmr_grpc_blockchain/pb"

func InitTransaction(sender, receiver string, value float64) *pb.Transaction {
	return &pb.Transaction{
		SenderAddress:   sender,
		ReceiverAddress: receiver,
		Value:           value,
	}
}
