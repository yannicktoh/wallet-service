package main

import (
	"context"
	"fmt"
	"log"

	"qredet.com/wallet-service/internal/money"
	"qredet.com/wallet-service/internal/repository"
	"qredet.com/wallet-service/internal/wallet"
)

func main() {
	// 1. Créer le repository (ici InMemory, pourrait être Postgres)
	repo := repository.NewInMemoryRepository()

	// 2. Pré-peupler avec des wallets
	repo.Seed(&wallet.Wallet{ID: "alice", Owner: "Alice", Balance: money.Money(5000)}) // 50.00€
	repo.Seed(&wallet.Wallet{ID: "bob", Owner: "Bob", Balance: money.Money(1000)})     // 10.00€

	// 3. Injecter le repository dans le service
	svc := wallet.NewWalletService(repo)

	// 4. Effectuer un transfert
	amount, err := money.NewMoney(2000) // 20.00€
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := svc.Transfer(ctx, "alice", "bob", amount); err != nil {
		log.Fatalf("Transfer failed: %v", err)
	}

	fmt.Println("Transfer successful!")
	fmt.Println("Alice: 30.00€, Bob: 30.00€")
}