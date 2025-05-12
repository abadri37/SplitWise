package repository

import (
	"SplitWise/internal/domain/entity"
	"SplitWise/internal/logger"
	"sync"
)

type ExpenseRepository struct {
	Expenses map[string]*entity.Expense
	Mutex    sync.RWMutex
}

func (r *ExpenseRepository) Save(expense *entity.Expense) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Expenses[expense.ID] = expense
	logger.GetLogger("ExpenseRepository").Infof("Saved Expense %s", expense.ID)
}

func (r *ExpenseRepository) GetExpense(expenseId string) (*entity.Expense, bool) {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	expense, ok := r.Expenses[expenseId]
	return expense, ok
}
