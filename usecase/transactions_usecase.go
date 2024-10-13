package usecase

import (
	"coeffee/domain"
	"context"
	"time"
)

type transactionUseCase struct {
	TransactionsRepository domain.TransactionsRepository
	contextTimeout    time.Duration
}

func NewTransactionsUseCase(prodrepo domain.TransactionsRepository, timeout time.Duration) domain.TransactionsUseCase {
	return &transactionUseCase{
		TransactionsRepository: prodrepo,
		contextTimeout:    timeout,
	}
}

func (uc *transactionUseCase) CreateTransaction(transaction domain.Transactions) (domain.Transactions, error) {
	_, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	transaction, err := uc.TransactionsRepository.CreateTransaction(transaction)
	if err != nil {
		return domain.Transactions{}, err
	}

	return transaction, nil
}

func (uc *transactionUseCase) GetTransactions(userId string, role string) ([]domain.Transactions, error) {
	_, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	transactions, err := uc.TransactionsRepository.GetTransactions( userId, role)
	if err != nil {
		return []domain.Transactions{}, err
	}

	return transactions, nil
}