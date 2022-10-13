package pb_service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hmr_grpc_blockchain/pb"
	"time"
)

func InitBlock(nonce int64, previousHash [32]byte, transactions []*pb.Transaction) *pb.Block {
	return &pb.Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: fmt.Sprintf("%x", previousHash),
		Transactions: transactions,
	}
}

func BlockHash(b *pb.Block) [32]byte {
	mj, _ := json.Marshal(b)
	h := sha256.Sum256(mj)
	return h
}
