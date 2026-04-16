package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type RecurringType string

const (
	RecurringTypeDaily   RecurringType = "daily"
	RecurringTypeMonthly RecurringType = "monthly"
	RecurringTypeDates   RecurringType = "dates"
	RecurringTypeEvenOdd RecurringType = "evenOdd"
)

type Recurring struct {
	Type RecurringType `json:"type"`

	EveryNDays  int         `json:"every_n_days,omitempty"`
	DaysOfMonth []int       `json:"days_of_month,omitempty"`
	Dates       []time.Time `json:"dates,omitempty"`
	Odd         *bool       `json:"odd,omitempty"`
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsRecurring bool       `json:"is_recurring"`
	Recurring   *Recurring `json:"recurring"`
	Status      Status     `json:"status"`
	ExecutorIds []string   `json:"executor_ids"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (s Status) Valid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}

func (rec *Recurring) Valid() error {
	switch rec.Type {
	case RecurringTypeDaily:
		if rec.EveryNDays <= 0 {
			return fmt.Errorf("every_n_days must be > 0")
		}
	case RecurringTypeMonthly:
		if len(rec.DaysOfMonth) == 0 {
			return fmt.Errorf("days_of_month required")
		}
		for _, day := range rec.DaysOfMonth {
			if day < 1 || day > 30 {
				return fmt.Errorf("day %d must be between 1 and 30", day)
			}
		}
	case RecurringTypeDates:
		if len(rec.Dates) == 0 {
			return fmt.Errorf("dates required")
		}
	case RecurringTypeEvenOdd:
		if rec.Odd == nil {
			return fmt.Errorf("odd must be set (true=odd, false=even)")
		}
	default:
		return fmt.Errorf("invalid recurring type: %s", rec.Type)
	}

	return nil
}

func checkExecutors() bool {
	// ...
	return true
}
