package wallet

import (
	"context"
	"fmt"

	"qredet.com/wallet-service/internal/money"
)

// WalletService contient la logique métier des transferts.
type WalletService struct {
	repo WalletRepository
}

// NewWalletService est le constructeur. Il injecte la dépendance.
func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

// Transfer transfère `amount` centimes du wallet `fromID` vers `toID`.
func (s *WalletService) Transfer(ctx context.Context, fromID, toID string, amount money.Money) error {
	if amount.Int64() <= 0 {
		return money.ErrInvalidAmount
	}

	//Récupération du wallet émetteur
	from, err := s.repo.GetByID(ctx, fromID)
	if err != nil {
		return fmt.Errorf("getting sender wallet: %w", err)
	}

	//Récupération du wallet destinataire
	to, err := s.repo.GetByID(ctx, toID)
	if err != nil {
		return fmt.Errorf("getting recipient wallet: %w", err)
	}

	//Vérification des fonds
	if from.Balance.Int64() < amount.Int64() {
		return ErrInsufficientFunds
	}

	//Débit : nouvelle balance = ancienne - montant
	from.Balance = money.Money(from.Balance.Int64() - amount.Int64())

	//nouvelle balance = ancienne + montant
	to.Balance = money.Money(to.Balance.Int64() + amount.Int64())

	//Persistance de l'émetteur mis à jour
	if err := s.repo.Update(ctx, from); err != nil {
		return fmt.Errorf("updating sender wallet: %w", err)
	}

	//Persistance du destinataire mis à jour
	if err := s.repo.Update(ctx, to); err != nil {
		return fmt.Errorf("updating recipient wallet: %w", err)
	}

	return nil
}