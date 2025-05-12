package main

import (
	"SplitWise/internal/domain/entity"
	"SplitWise/internal/repository"
	"SplitWise/internal/service"
	"fmt"
)

func main() {
	fmt.Println("🚀 Hello, Go!")
	fmt.Println("Sample Program")
	fmt.Println("Hello World")

	userRepo := &repository.UserRepository{Users: make(map[string]*entity.User)}
	expRepo := &repository.ExpenseRepository{Expenses: make(map[string]*entity.Expense)}
	groupRepo := &repository.GroupRepository{Groups: make(map[string]*entity.ExpenseGroup)}
	service := &service.ExpenseService{UserRepo: userRepo, GroupRepo: groupRepo, ExpenseRepo: expRepo}

	u1 := userRepo.RegisterUser("u1", "Alice", "alice@example.com")
	u2 := userRepo.RegisterUser("u2", "Bob", "bob@example.com")
	u3 := userRepo.RegisterUser("u3", "Charlie", "charlie@example.com")

	service.CreateExpenseGroup("g1", "Trip", []string{u1.ID, u2.ID, u3.ID})

	splits := map[string]float64{u1.ID: 100, u2.ID: 150, u3.ID: 250}
	exp := service.CreateExpense("e1", "Dinner", u1.ID, 500, splits)

	service.ShareExpense(exp.ID)
	service.AddContribution(exp.ID, u2.ID, 150)
	service.AddContribution(exp.ID, u1.ID, 100)
	service.AddContribution(exp.ID, u3.ID, 250)
}
