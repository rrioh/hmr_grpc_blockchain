package main

import (
	"fmt"
	"hmr_grpc_blockchain/domain/service"
)

func main() {
	bch := service.InitBlockchain(3)

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

	fmt.Println("IS VALID?: ", bch.IsValidChain())
}
