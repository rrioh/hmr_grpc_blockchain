package service

import (
	"fmt"
	"hmr_grpc_blockchain/domain/model"
	"time"
)

func InitBlock(nonce int64, previousHash [32]byte, transactions []*model.Transaction) *model.Block {
	return &model.Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: fmt.Sprintf("%x", previousHash),
		Transactions: transactions,
	}
}
