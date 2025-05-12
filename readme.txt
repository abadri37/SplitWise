📘 README.txt

🎯 Project Overview
=====================
This project is a simplified SplitWise clone built in Golang that allows users to manage shared expenses in groups. The architecture emphasizes clean code, separation of concerns, and maintainability.

✅ Key Features:
- Users can join groups
- Create shared expenses
- Notifiers update users when expenses are added
- Implements Observer pattern for notifications

---

🧱 SOLID Principles
=====================

1️⃣ **S - Single Responsibility Principle**  
   Each package and type has a focused responsibility:  
   - `UserRepository`, `ExpenseRepository`, and `GroupRepository` only handle DB logic.  
   - `ExpenseService` handles expense business logic.  
   - `UserNotifier` handles notification logic.

2️⃣ **O - Open/Closed Principle**  
   The system is open for extension but closed for modification:  
   - You can add new types of observers (e.g., `EmailNotifier`, `SMSNotifier`) without changing existing logic.

3️⃣ **L - Liskov Substitution Principle**  
   - Any implementation of the `Observer` interface can be used interchangeably.

4️⃣ **I - Interface Segregation Principle**  
   - The `Observer` interface is simple and focused (`Notify(string)`), not bloated with unnecessary methods.

5️⃣ **D - Dependency Inversion Principle**  
   - High-level modules (services) depend on abstractions (`interface`) instead of concrete implementations:
     - `ExpenseService` uses `UserRepository`, `GroupRepository`, and `ExpenseRepository` as injected dependencies via pointers.

---

🔁 Design Patterns
=====================

🧩 **Observer Pattern**
- Interface: `Observer`
- Concrete observer: `UserNotifier`
- Use case: Automatically notify users when a new expense is added to a group.

```go
type Observer interface {
    Notify(event string)
}

type UserNotifier struct {
    UserID string
}

func (n *UserNotifier) Notify(event string) {
    logger.GetLogger("Observer").Infof("Notify User %s: %s", n.UserID, event)
}

for _, obs := range expense.Observers {
    if un, ok := obs.(*observer.UserNotifier); ok && un.UserID == userId {
        obs.Notify(fmt.Sprintf("Thanks for your contribution of %.2f to expense %s", amount, expense.Title))
    }
}

🛠️ Repository Pattern

Abstracts data access logic from business logic.

Repositories like UserRepository, GroupRepository, and ExpenseRepository handle database operations.

🏗️ Dependency Injection

Dependencies like repositories are injected into the service layer using pointers.

type ExpenseService struct {
    UserRepo    *repository.UserRepository
    GroupRepo   *repository.GroupRepository
    ExpenseRepo *repository.ExpenseRepository
}

SplitWise/
├── internal/
│   ├── service/           // Business logic
│   ├── repository/        // Data access
│   ├── observer/          // Observer pattern implementation
│   └── logger/            // Custom logging
└── main.go                // Entry point

🧪 Testing
Each component (service, repository, observer) can be unit tested independently.

Mocks can be injected due to use of interfaces and DI.

🚀 Extensibility Ideas
Add more observers (e.g., SlackNotifier)

Introduce Strategy Pattern to split expenses differently

Use Factory Pattern for creating expense types

🔚 Conclusion
This codebase demonstrates clean architecture, core Go idioms, and well-established software engineering principles like SOLID and classic patterns like Observer and Repository. 🧑‍💻
