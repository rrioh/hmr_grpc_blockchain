package pb_service

import "hmr_grpc_blockchain/pb"

/*
Transaction構造体を初期化する。トランザクションプールにトランザクションを新たに追加するときなどに使用する。
*/
func InitTransaction(sender, receiver string, value float64) *pb.Transaction {
	return &pb.Transaction{
		SenderAddress:   sender,
		ReceiverAddress: receiver,
		Value:           value,
	}
}
