package wallet_test

import (
	"context"
	"errors"
	"testing"

	"qredet.com/wallet-service/internal/money"
	"qredet.com/wallet-service/internal/repository"
	"qredet.com/wallet-service/internal/wallet"
)

// helper : crée un service avec un InMemoryRepository pré-peuplé
func setupService(wallets ...*wallet.Wallet) *wallet.WalletService {
	repo := repository.NewInMemoryRepository()
	for _, w := range wallets {
		repo.Seed(w)
	}
	return wallet.NewWalletService(repo)
}

// --- Test 1 : Transfert réussi ---
func TestTransfer_Success(t *testing.T) {
	alice := &wallet.Wallet{ID: "alice", Owner: "Alice", Balance: money.Money(100000)} // 100.000 FCFA
	bob := &wallet.Wallet{ID: "bob", Owner: "Bob", Balance: money.Money(5000)}        // 5.000 FCFA

	svc := setupService(alice, bob)

	amount, _ := money.NewMoney(3000) // 3.000 FCFA
	err := svc.Transfer(context.Background(), "alice", "bob", amount)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

// --- Test 2 : Fonds insuffisants ---
func TestTransfer_InsufficientFunds(t *testing.T) {
	alice := &wallet.Wallet{ID: "alice", Owner: "Alice", Balance: money.Money(1000)} // 1.000 FCFA
	bob := &wallet.Wallet{ID: "bob", Owner: "Bob", Balance: money.Money(5000)}

	svc := setupService(alice, bob)

	amount, _ := money.NewMoney(5000) // 5.000 FCFA > solde alice
	err := svc.Transfer(context.Background(), "alice", "bob", amount)

	if !errors.Is(err, wallet.ErrInsufficientFunds) {
		t.Fatalf("expected ErrInsufficientFunds, got: %v", err)
	}
}

// --- Test 3 : Wallet source introuvable ---
func TestTransfer_SenderNotFound(t *testing.T) {
	bob := &wallet.Wallet{ID: "bob", Owner: "Bob", Balance: money.Money(5000)}
	svc := setupService(bob)

	amount, _ := money.NewMoney(1000)
	err := svc.Transfer(context.Background(), "unknown", "bob", amount)

	if !errors.Is(err, wallet.ErrWalletNotFound) {
		t.Fatalf("expected ErrWalletNotFound, got: %v", err)
	}
}

// --- Test 4 : DI — FailingRepository ---
func TestTransfer_WithFailingRepository(t *testing.T) {
	failRepo := repository.NewFailingRepository()
	svc := wallet.NewWalletService(failRepo)

	amount, _ := money.NewMoney(1000)
	err := svc.Transfer(context.Background(), "alice", "bob", amount)

	// On s'attend à une erreur liée au storage, wrappée par le service
	if !errors.Is(err, repository.ErrStorageUnavailable) {
		t.Fatalf("expected ErrStorageUnavailable, got: %v", err)
	}
}

// --- Test 5 : Montant invalide ---
func TestTransfer_InvalidAmount(t *testing.T) {
	_, err := money.NewMoney(-500)
	if !errors.Is(err, money.ErrInvalidAmount) {
		t.Fatalf("expected ErrInvalidAmount, got: %v", err)
	}
}