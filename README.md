# Wallet Service

A small Go service for transferring money between wallets.

## How to run

```bash
go run ./cmd/main.go
go test ./...
```

## Dependency Injection

`WalletService` depends on the `WalletRepository` **interface**, not on any concrete implementation. The concrete repository is passed via `NewWalletService(repo)` — the constructor.

This means:
- In tests → `InMemoryRepository` 
- In production → could be `PostgresRepository`
- For error testing → `FailingRepository`

The service code never changes, only the injected dependency does.

## Money Representation

Money is represented as `int64` **cents** (smallest unit).

- `float64` is **never used** for money due to IEEE 754 rounding errors.
- `NewMoney(cents)` validates that the amount is strictly positive.

## Known Limitations

- Transfers are **not atomic**: if `Update(from)` succeeds but `Update(to)` fails, funds are lost. A production system would use a database transaction.