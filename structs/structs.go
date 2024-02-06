package structs

type Status string

const (
	Pending   Status = "pending"
	Completed Status = "completed"
	Overdue   Status = "overdue"
)

type Tasks struct {
	// The list of tasks
	ID     int    `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	Time   string `json:"time"`
	Status Status `json:"status"`
}

type TempTasks struct {
	Task   string `json:"task"`
	Time   string `json:"time"`
	Status Status `json:"status"`
}
