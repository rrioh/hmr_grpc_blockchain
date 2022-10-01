package service

import (
	"errors"
	"fmt"
	"hmr_grpc_blockchain/domain/model"
	"strings"
)

type BlockchainHandlerInterface interface{}

type BlockchainHandler struct {
	Blockchain *model.Blockchain
}

func InitBlockchain(difficulty int) *BlockchainHandler {
	bch := &BlockchainHandler{
		Blockchain: &model.Blockchain{
			Difficulty: difficulty,
		},
	}

	// first block
	genesisBlock := InitBlock(0, [32]byte{}, nil)
	bch.addBlock(genesisBlock)
	return bch
}

func (bch *BlockchainHandler) AddTransaction(sender, receiver string, value float64) {
	transaction := InitTransaction(sender, receiver, value)
	bch.Blockchain.TransactionPool = append(bch.Blockchain.TransactionPool, transaction)
}

func (bch *BlockchainHandler) Mining() error {
	bc := bch.Blockchain

	if len(bc.TransactionPool) == 0 {
		return errors.New("transaction does not exist")
	}

	prevHash := bc.Chain[len(bc.Chain)-1].Hash()
	b := InitBlock(0, prevHash, bc.TransactionPool)
	b = bch.proofOfWork(b)
	bch.addBlock(b)

	return nil
}

func (bch *BlockchainHandler) IsValidChain() bool {
	chain := bch.Blockchain.Chain
	for i := 0; i < len(chain)-1; i++ {
		if fmt.Sprintf("%x", chain[i].Hash()) != chain[i+1].PreviousHash {
			return false
		}
	}

	return true
}

// private

func (bch *BlockchainHandler) proofOfWork(b *model.Block) *model.Block {
	bc := bch.Blockchain

	pow := strings.Repeat("0", bc.Difficulty)

	for {
		hash := fmt.Sprintf("%x", b.Hash())
		if strings.HasPrefix(hash, pow) {
			return b
		}

		b.Nonce += 1
	}
}

// do not reset TransactionPool before this func
func (bch *BlockchainHandler) addBlock(b *model.Block) {
	bc := bch.Blockchain

	bc.Chain = append(bc.Chain, b)

	bc.TransactionPool = []*model.Transaction{}
}
