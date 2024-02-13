package models

// Hand coded additions

// Task can be used in "AutoMigrate"
type Task struct {
	TaskBase
	Project Project `gorm:"foreignKey:PId;references:PId"`
}

func (t *TaskLi) TableName() string {
	return "tasks"
}

const (
	RefProjectTasks = "RefTasks"
)

type ProjectWithTasks struct {
	Project
	RefTasks []*TaskLi `gorm:"ForeignKey:PId;references:PId"`
}
