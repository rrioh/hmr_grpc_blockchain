package pb_service

import (
	"errors"
	"fmt"
	"hmr_grpc_blockchain/pb"
	"strings"
)

type BlockchainHandlerInterface interface{}

type BlockchainHandler struct {
	Blockchain *pb.Blockchain
}

type Blockchain pb.Blockchain

func InitBlockchain(address string, difficulty int) *BlockchainHandler {
	bch := &BlockchainHandler{
		Blockchain: &pb.Blockchain{
			Address:    address,
			Difficulty: int32(difficulty),
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

	prevHash := BlockHash(bc.Chain[len(bc.Chain)-1])
	b := InitBlock(0, prevHash, bc.TransactionPool)
	b = bch.proofOfWork(b)
	bch.addBlock(b)

	return nil
}

func (bch *BlockchainHandler) GetAddressBalance(address string) float64 {
	bc := bch.Blockchain

	balance := 0.0
	for _, chain := range bc.Chain {
		for _, transaction := range chain.Transactions {
			if transaction.SenderAddress == address {
				balance -= transaction.Value
			}
			if transaction.ReceiverAddress == address {
				balance += transaction.Value
			}
		}
	}

	return balance
}

func (bch *BlockchainHandler) IsValidChain() bool {
	chain := bch.Blockchain.Chain
	for i := 0; i < len(chain)-1; i++ {
		if fmt.Sprintf("%x", BlockHash(chain[i])) != chain[i+1].PreviousHash {
			return false
		}
	}

	return true
}

// private

func (bch *BlockchainHandler) proofOfWork(b *pb.Block) *pb.Block {
	bc := bch.Blockchain

	pow := strings.Repeat("0", int(bc.Difficulty))

	for {
		hash := fmt.Sprintf("%x", BlockHash(b))
		if strings.HasPrefix(hash, pow) {
			return b
		}

		b.Nonce += 1
	}
}

// do not reset TransactionPool before this func
func (bch *BlockchainHandler) addBlock(b *pb.Block) {
	bc := bch.Blockchain

	bc.Chain = append(bc.Chain, b)

	bc.TransactionPool = []*pb.Transaction{}
}
