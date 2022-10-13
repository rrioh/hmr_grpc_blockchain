package main

import (
	"context"
	"errors"
	"hmr_grpc_blockchain/domain/model"
	"hmr_grpc_blockchain/domain/service"
	"hmr_grpc_blockchain/domain/service/pb_service"
	"hmr_grpc_blockchain/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BlockchainServer struct {
	pb.UnimplementedBlockchainServiceServer
	BlockchainHandler *pb_service.BlockchainHandler
	Config            *model.ServerConfig
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	blockchainServer := &BlockchainServer{}

	config, err := service.NewServerConfig()
	if err != nil {
		log.Fatal(err)
	}
	blockchainServer.Config = config

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
		bch := pb_service.InitBlockchain(bs.Config.Blockchain.Address, bs.Config.Blockchain.Difficulty)
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

func (bs *BlockchainServer) Mining(ctx context.Context, req *pb.MiningRequest) (*pb.MiningResponse, error) {
	if bs.BlockchainHandler == nil {
		return nil, errors.New("blockchain does not exist")
	}

	bch := bs.BlockchainHandler
	err := bch.Mining()
	if err != nil {
		return nil, err
	}

	return &pb.MiningResponse{}, nil
}

func (bs *BlockchainServer) GetAddressBalance(ctx context.Context, req *pb.GetAddressBalanceRequest) (*pb.GetAddressBalanceResponse, error) {
	if bs.BlockchainHandler == nil {
		return nil, errors.New("blockchain does not exist")
	}

	address := req.Address

	bch := bs.BlockchainHandler
	balance := bch.GetAddressBalance(address)

	res := &pb.GetAddressBalanceResponse{
		Balance: balance,
	}

	return res, nil
}

func (bs *BlockchainServer) IsValidChain(ctx context.Context, req *pb.IsValidChainRequest) (*pb.IsValidChainResponse, error) {
	if bs.BlockchainHandler == nil {
		return nil, errors.New("blockchain does not exist")
	}

	bch := bs.BlockchainHandler
	isValid := bch.IsValidChain()

	res := &pb.IsValidChainResponse{
		IsValid: isValid,
	}

	return res, nil
}
