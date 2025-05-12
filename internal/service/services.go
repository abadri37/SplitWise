package service

import (
	"SplitWise/internal/domain/entity"
	"SplitWise/internal/logger"
	"SplitWise/internal/observer"
	"SplitWise/internal/repository"
	"fmt"
	"time"
)

type ExpenseService struct {
	UserRepo    *repository.UserRepository
	GroupRepo   *repository.GroupRepository
	ExpenseRepo *repository.ExpenseRepository
}

func (s *ExpenseService) CreateExpenseGroup(groupId, name string, userIds []string) *entity.ExpenseGroup {
	group := &entity.ExpenseGroup{
		ID:      groupId,
		Name:    name,
		Members: make(map[string]*entity.User),
	}
	for _, id := range userIds {
		if user, ok := s.UserRepo.GetUser(id); ok {
			group.Members[id] = user
		}
	}
	s.GroupRepo.RegisterGroup(group)
	logger.GetLogger("ExpenseService").Infof("Created group with users:  %s", groupId)
	return group
}

func (s *ExpenseService) CreateExpense(id, title, createdBy string, total float64, splits map[string]float64) *entity.Expense {
	expense := &entity.Expense{
		ID:           id,
		Title:        title,
		CreatedBy:    createdBy,
		State:        entity.Created,
		TotalAmount:  total,
		CreatedAt:    time.Now(),
		Observers:    []observer.Observer{},
		History:      []entity.ExpenseHistory{},
		Contributors: make(map[string]*entity.Contribution),
	}

	for uid, amount := range splits {
		expense.Contributors[uid] = &entity.Contribution{UserID: uid, Amount: amount, Paid: 0}
		expense.Observers = append(expense.Observers, &observer.UserNotifier{UserID: uid})
	}
	expense.History = append(expense.History, entity.ExpenseHistory{Timestamp: time.Now(), Message: "Expense Created"})
	s.ExpenseRepo.Save(expense)
	logger.GetLogger("ExpenseService").Infof("Created Expense %s by : %s", id, createdBy)
	return expense
}

func (s *ExpenseService) ShareExpense(expenseId string) {
	if expense, ok := s.ExpenseRepo.GetExpense(expenseId); ok {
		expense.State = entity.Pending
		expense.History = append(expense.History, entity.ExpenseHistory{Timestamp: time.Now(), Message: "Expense shared and pending contributions"})
		s.ExpenseRepo.Save(expense)
		for _, observer := range expense.Observers {
			observer.Notify(fmt.Sprintf("Expense %s shared and pending contributions", expense.ID))
		}
		logger.GetLogger("ExpenseService").Infof("Shared Expense %s", expenseId)
	}
}

func (s *ExpenseService) AddContribution(expenseId, userId string, amount float64) {
	if expense, ok := s.ExpenseRepo.GetExpense(expenseId); ok {
		if contribution, exists := expense.Contributors[userId]; exists {
			contribution.Paid += amount
			expense.History = append(expense.History, entity.ExpenseHistory{Timestamp: time.Now(), Message: fmt.Sprintf("Thanks for your contribution of %f to %s", amount, expenseId)})
			for _, obs := range expense.Observers {
				if un, ok := obs.(*observer.UserNotifier); ok && un.UserID == userId {
					obs.Notify(fmt.Sprintf("Thanks for your contribution of %.2f to expense %s", amount, expense.Title))
				}
			}
			logger.GetLogger("ExpenseService").Infof("User %s Contributed %.2f to expense %s", userId, amount, expenseId)
			settled := true
			for _, contribution := range expense.Contributors {
				if contribution.Paid < contribution.Amount {
					settled = false
					break
				}
			}
			if settled {
				expense.State = entity.Settled
				now := time.Now()
				expense.SettledAt = &now
				expense.History = append(expense.History, entity.ExpenseHistory{Timestamp: now, Message: "Expense settled"})
				for _, observer := range expense.Observers {
					observer.Notify(fmt.Sprintf("Expense %s settled", expense.ID))
				}
				logger.GetLogger("ExpenseService").Infof("Expense %s settled", expenseId)
			}
			s.ExpenseRepo.Save(expense)
		}
	}
}
