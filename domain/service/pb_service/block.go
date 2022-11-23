package pb_service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hmr_grpc_blockchain/pb"
	"time"
)

/*
Block構造体を初期化する。Blockchain構造体を初期化する際の1ブロック目、またマイニングでブロックを追加する際などに利用する。
*/
func InitBlock(nonce int64, previousHash [32]byte, transactions []*pb.Transaction) *pb.Block {
	return &pb.Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: fmt.Sprintf("%x", previousHash),
		Transactions: transactions,
	}
}

/*
Blockのハッシュ値を求める。マイニングの際のPOWなどで利用する。
*/
func BlockHash(b *pb.Block) [32]byte {
	mj, _ := json.Marshal(b)
	h := sha256.Sum256(mj)
	return h
}
