package model

import (
	"crypto/sha256"
	"encoding/json"
)

// model

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Nonce        int64          `json:"nonce"`
	PreviousHash string         `json:"previous_hash"`
	Transactions []*Transaction `json:"transactions"`
}

type Blockchain struct {
	Address         string         `json:"address"`
	Difficulty      int            `json:"difficulty"`
	Chain           []*Block       `json:"chain"`
	TransactionPool []*Transaction `json:"transaction_pool"`
}

type Transaction struct {
	SenderAddress   string  `json:"sender_address"`
	ReceiverAddress string  `json:"receiver_address"`
	Value           float64 `json:"value"`
}

// model method

func (b *Block) String() string {
	mj, _ := json.Marshal(*b)
	return string(mj)
}

func (b *Block) Hash() [32]byte {
	mj, _ := json.Marshal(*b)
	h := sha256.Sum256(mj)
	return h
}

func (bc *Blockchain) String() string {
	mj, _ := json.Marshal(*bc)
	return string(mj)
}

func (t *Transaction) String() string {
	mj, _ := json.Marshal(*t)
	return string(mj)
}
