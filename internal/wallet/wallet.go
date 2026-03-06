package wallet

import (
	"context"
	"errors"
	"qredet.com/wallet-service/internal/money"
)

// Wallet représente un portefeuille utilisateur.
type Wallet struct {
	ID      string
	Owner   string
	Balance money.Money
}

// Erreurs métier exposées publiquement.
var (
	ErrWalletNotFound    = errors.New("wallet not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// WalletRepository est l'interface de persistance.
type WalletRepository interface {
	// GetByID retourne le wallet correspondant à l'ID.
	// Retourne ErrWalletNotFound si absent.
	GetByID(ctx context.Context, id string) (*Wallet, error)

	// Update persiste les modifications d'un wallet.
	Update(ctx context.Context, w *Wallet) error
}