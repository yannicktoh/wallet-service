package repository

import (
	"context"
	"errors"

	"qredet.com/wallet-service/internal/wallet"
)

var ErrStorageUnavailable = errors.New("storage unavailable")

// FailingRepository est un stub qui échoue systématiquement.
type FailingRepository struct{}

func NewFailingRepository() *FailingRepository {
	return &FailingRepository{}
}

func (r *FailingRepository) GetByID(_ context.Context, _ string) (*wallet.Wallet, error) {
	return nil, ErrStorageUnavailable
}

func (r *FailingRepository) Update(_ context.Context, _ *wallet.Wallet) error {
	return ErrStorageUnavailable
}