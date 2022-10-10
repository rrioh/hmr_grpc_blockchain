package main

import (
	"context"
	"errors"
	"hmr_grpc_blockchain/domain/service"
	"hmr_grpc_blockchain/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	BLOCKCHAIN_ADDRESS = "BLOCKCHAIN_SERVER_ADDRESS"
	DIFFICULTRY        = 3
)

type BlockchainServer struct {
	pb.UnimplementedBlockchainServiceServer
	BlockchainHandler *service.BlockchainHandler
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	blockchainServer := &BlockchainServer{}

	pb.RegisterBlockchainServiceServer(grpcServer, blockchainServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}

	/*bch := service.InitBlockchain("BLOCKCHAIN_SERVER_ADDRESS", 3)

	bch.AddTransaction("A", "B", 6.3)
	bch.AddTransaction("C", "D", 5.2)
	fmt.Println("======= BEFORE FRIST MINING ========")
	fmt.Println(*bch)
	bch.Mining()
	fmt.Println("======= AFTER FRIST MINING ========")
	fmt.Println(*bch)

	bch.AddTransaction("E", "F", 12)
	bch.AddTransaction("G", "H", 1.3)
	fmt.Println("======= BEFORE SECOND MINING ========")
	fmt.Println(*bch)
	bch.Mining()
	fmt.Println("======= AFTER SECOND MINING ========")
	fmt.Println(*bch)

	bch.AddTransaction("H", "I", 54)
	bch.AddTransaction("J", "K", 13)
	fmt.Println("======= BEFORE THIRD MINING ========")
	fmt.Println(*bch)
	bch.Mining()
	fmt.Println("======= AFTER THIRD MINING ========")
	fmt.Println(*bch)

	fmt.Println("IS VALID?: ", bch.IsValidChain())*/
}

func (bs *BlockchainServer) GetBlockchain(ctx context.Context, req *pb.GetBlockchainRequest) (*pb.GetBlockchainResponse, error) {
	if bs.BlockchainHandler == nil {
		bch := service.InitBlockchain(BLOCKCHAIN_ADDRESS, DIFFICULTRY)
		bs.BlockchainHandler = bch
	}

	res := &pb.GetBlockchainResponse{
		Blockchain: bs.BlockchainHandler.Blockchain,
	}
	return res, nil
}

func (bs *BlockchainServer) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest) (*pb.AddTransactionResponse, error) {
	if bs.BlockchainHandler == nil {
		return nil, errors.New("blockchain does not exist")
	}

	senderAddress := req.SenderAddress
	receiverAddress := req.ReceiverAddress
	value := req.Value

	bch := bs.BlockchainHandler
	bch.AddTransaction(senderAddress, receiverAddress, value)

	return &pb.AddTransactionResponse{}, nil
}
