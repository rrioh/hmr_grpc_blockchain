package pb_service

import (
	"errors"
	"fmt"
	"hmr_grpc_blockchain/pb"
	"strings"
)

/*
Blockchain構造体を操作する各種メソッドを内包したインターフェース。
BlockchainHandler構造体が実装している。
*/
type BlockchainHandlerInterface interface{}

/*
BlockchainHandlerInterfaceを実装している構造体。
Blockchain構造体を内包し、Blockchainに対して各種処理を行うレシーバとなる。
*/
type BlockchainHandler struct {
	Blockchain *pb.Blockchain
}

/*
Blockchain構造体の初期化を行う。
一つ目のブロックはナンス値0、前ブロックのハッシュ値が空のバイト列、トランザクションがnilとなる。
*/
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

/*
Blockchainの保有する未取り込みのトランザクションプールにトランザクションを追加する。
振込人アドレス、引き出し人アドレス、送金額を指定する。
*/
func (bch *BlockchainHandler) AddTransaction(sender, receiver string, value float64) {
	transaction := InitTransaction(sender, receiver, value)
	bch.Blockchain.TransactionPool = append(bch.Blockchain.TransactionPool, transaction)
}

/*
マイニングを行い、ブロックを作成する。
ブロック作成に伴い、トランザクションプール内のトランザクションは当該ブロックに全て取り込まれ、プールは空になる。
トランザクションプールが空の場合はマイニングは失敗し、エラーを返却する。
*/
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

/*
対象のアドレスが保有する金額を計算する。
Blockchain内のトランザクションをはじめから追跡し、最新ブロックに至るまでの取引後の残高を返却する。
*/
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

/*
Blockchainが有効かどうかを判定する。
各ブロックが保有する前ブロックのハッシュ値と、実際の前ブロックのハッシュ値が異なる場合は無効判定となる。
*/
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
