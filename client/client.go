package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hmr_grpc_blockchain/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewBlockchainServiceClient(conn)

	reqGetBlockchain(client)

	reqAddTransaction(client, "AA", "BB", 6.34)
	reqGetBlockchain(client)
	reqAddTransaction(client, "BB", "AA", 3)
	reqGetBlockchain(client)
	reqMining(client)
	reqGetBlockchain(client)

	reqAddTransaction(client, "EE", "AA", 2.5)
	reqGetBlockchain(client)
	reqAddTransaction(client, "GG", "HH", 8)
	reqGetBlockchain(client)
	reqMining(client)
	reqGetBlockchain(client)

	reqGetAddressBalance(client, "BB")

	reqIsValidChain(client)
}

func reqGetBlockchain(client pb.BlockchainServiceClient) {
	res, err := client.GetBlockchain(context.Background(), &pb.GetBlockchainRequest{})
	if err != nil {
		log.Fatal(err)
	}

	bc := res.Blockchain
	m, _ := json.Marshal(&bc)
	fmt.Println("get response: ", string(m))
}

func reqAddTransaction(client pb.BlockchainServiceClient, sender, receiver string, value float64) {
	req := &pb.AddTransactionRequest{
		SenderAddress:   sender,
		ReceiverAddress: receiver,
		Value:           value,
	}
	_, err := client.AddTransaction(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("add transaction. sender: %s, receiver: %s, value: %.2f\n", sender, receiver, value)
}

func reqMining(client pb.BlockchainServiceClient) {
	_, err := client.Mining(context.Background(), &pb.MiningRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mining done")
}

func reqGetAddressBalance(client pb.BlockchainServiceClient, address string) float64 {
	req := &pb.GetAddressBalanceRequest{
		Address: address,
	}
	res, err := client.GetAddressBalance(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	balance := res.Balance
	fmt.Printf("balance of address %s is %.2f\n", address, balance)

	return balance
}

func reqIsValidChain(client pb.BlockchainServiceClient) bool {
	res, err := client.IsValidChain(context.Background(), &pb.IsValidChainRequest{})
	if err != nil {
		log.Fatal(err)
	}

	isValid := res.IsValid
	fmt.Println("is valid blockchain?:", isValid)

	return isValid
}
