package main

import (
	"errors"
	"fmt"
)

// ============================================
// S — Single Responsibility Principle
// ============================================

// User отвечает за данные пользователя.
type User struct {
	Name  string
	Email string
}

// ============================================
// I — Interface Segregation Principle
// ============================================

// UserRepository описывает только операции с пользователями.
type UserRepository interface {
	Save(user User) error
}

// Notifier описывает только отправку уведомлений.
type Notifier interface {
	Send(user User, message string) error
}

// ============================================
// D — Dependency Inversion Principle
// ============================================

// UserService зависит от интерфейсов,
// а не от конкретных реализаций.
type UserService struct {
	repository UserRepository
	notifier   Notifier
}

func NewUserService(
	repository UserRepository,
	notifier Notifier,
) *UserService {
	return &UserService{
		repository: repository,
		notifier:   notifier,
	}
}

func (s *UserService) Register(user User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}

	if err := s.repository.Save(user); err != nil {
		return err
	}

	return s.notifier.Send(user, "Welcome!")
}

// ============================================
// O — Open/Closed Principle
// ============================================

// Реализация сохранения в памяти.
type MemoryRepository struct {
	users []User
}

func (r *MemoryRepository) Save(user User) error {
	r.users = append(r.users, user)

	fmt.Println("User saved:", user.Name)

	return nil
}

// Реализация уведомлений по электронной почте.
type EmailNotifier struct{}

func (n *EmailNotifier) Send(user User, message string) error {
	fmt.Printf(
		"Email sent to %s: %s\n",
		user.Email,
		message,
	)

	return nil
}

// Новую реализацию можно добавить,
// не изменяя UserService.
type SMSNotifier struct{}

func (n *SMSNotifier) Send(user User, message string) error {
	fmt.Printf(
		"SMS sent to %s: %s\n",
		user.Name,
		message,
	)

	return nil
}

// ============================================
// main
// ============================================

func main() {
	repository := &MemoryRepository{}
	notifier := &EmailNotifier{}

	service := NewUserService(repository, notifier)

	user := User{
		Name:  "Nick",
		Email: "nick@example.com",
	}

	if err := service.Register(user); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("--- Using SMS ---")

	smsNotifier := &SMSNotifier{}
	serviceWithSMS := NewUserService(repository, smsNotifier)

	user2 := User{
		Name:  "Alex",
		Email: "alex@example.com",
	}

	if err := serviceWithSMS.Register(user2); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// User saved: Nick
// Email sent to nick@example.com: Welcome!
// --- Using SMS ---
// User saved: Alex
// SMS sent to Alex: Welcome!
