package repository

import (
	"context"
	"sync"

	"qredet.com/wallet-service/internal/wallet"
)

// InMemoryRepository stocke les wallets dans une map en mémoire.
type InMemoryRepository struct {
	mu      sync.RWMutex
	wallets map[string]*wallet.Wallet
}

// NewInMemoryRepository crée un repo vide.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		wallets: make(map[string]*wallet.Wallet),
	}
}

// Seed permet de pré-peupler le repo (utile dans les tests).
func (r *InMemoryRepository) Seed(w *wallet.Wallet) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// On stocke une copie pour éviter les mutations externes
	clone := *w
	r.wallets[w.ID] = &clone
}

// GetByID retourne le wallet par son ID.
func (r *InMemoryRepository) GetByID(_ context.Context, id string) (*wallet.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	w, ok := r.wallets[id]
	if !ok {
		return nil, wallet.ErrWalletNotFound
	}
	// On retourne une copie pour éviter les races conditions
	clone := *w
	return &clone, nil
}

// Update écrase le wallet existant avec les nouvelles données.
func (r *InMemoryRepository) Update(_ context.Context, w *wallet.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.wallets[w.ID]; !ok {
		return wallet.ErrWalletNotFound
	}
	clone := *w
	r.wallets[w.ID] = &clone
	return nil
}