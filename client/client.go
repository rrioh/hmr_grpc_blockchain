package main

import (
	"context"
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
}

func reqGetBlockchain(client pb.BlockchainServiceClient) {
	res, err := client.GetBlockchain(context.Background(), &pb.GetBlockchainRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get response: ", res)
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
