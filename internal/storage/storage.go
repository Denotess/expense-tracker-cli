package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Expense struct {
	ID          int64     `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Store struct {
	Path string
}

func DefaultPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "expense-tracker", "expenses.json"), nil
}

func New(path string) *Store {
	return &Store{Path: path}
}

func (s *Store) Load() ([]Expense, error) {
	var expenses []Expense
	file, err := os.ReadFile(s.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Expense{}, nil
		}
		return nil, err
	}
	if len(file) == 0 {
		return []Expense{}, nil
	}
	if err = json.Unmarshal(file, &expenses); err != nil {
		return nil, err
	}
	return expenses, nil

}

func (s *Store) Save(expenses []Expense) error {
	dir := filepath.Dir(s.Path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	file, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(s.Path, file, 0o644); err != nil {
		return err
	}
	return nil

}

func (s *Store) Add(amount int64, description string) (Expense, error) {
	expenses, err := s.Load()
	if err != nil {
		return Expense{}, err
	}
	id := int64(len(expenses) + 1)
	if amount <= 0 {
		return Expense{}, errors.New("amount must be larger than 0")
	}
	if strings.TrimSpace(description) == "" {
		return Expense{}, errors.New("description must not be empty")
	}
	now := time.Now()
	expenses = append(expenses, Expense{ID: id, Amount: amount, Description: description, Date: now})
	if err := s.Save(expenses); err != nil {
		return Expense{}, err
	}
	newExpense := Expense{ID: id, Amount: amount, Description: description, Date: now}
	return newExpense, nil
}

func (s *Store) Delete(id int64) error {
	expenses, err := s.Load()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("invalid input")
	}
	found := false
	expensesNew := make([]Expense, 0, len(expenses))
	for _, expense := range expenses {
		if expense.ID != id {
			expensesNew = append(expensesNew, expense)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("expense not found")
	}
	if err := s.Save(expensesNew); err != nil {
		return err
	}
	return nil
}

func (s *Store) List() ([]Expense, error) {
	expenses, err := s.Load()
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *Store) Update(id, amount int64, description string) error {
	expenses, err := s.Load()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("id must be > 0")
	}
	expensesNew := make([]Expense, 0, len(expenses))

	found := false
	for _, expense := range expenses {
		if expense.ID == id {
			expensesNew = append(expensesNew, Expense{ID: expense.ID, Amount: amount, Description: description, Date: expense.Date})
			found = true
			continue
		}
		expensesNew = append(expensesNew, expense)
	}
	if !found {
		return errors.New("expense not found")
	}
	if err = s.Save(expensesNew); err != nil {
		return err
	}
	return nil
}

func (s *Store) Summary(month int) (int64, error) {
	expenses, err := s.Load()
	if err != nil {
		return 0, err
	}
	if month > 12 || month < 0 {
		return 0, errors.New("month must be 0 or 1 - 12")
	}
	m := time.Month(month)
	var sum int64
	for _, expense := range expenses {
		if month != 0 && expense.Date.Month() != m {
			continue
		}
		sum += expense.Amount
	}
	return sum, nil
}
