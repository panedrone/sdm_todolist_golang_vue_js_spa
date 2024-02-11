package models

// Hand coded additions

const TasksFK = "Tasks"

type ProjectEx struct {
	// Project
	PId int64 `json:"p_id" gorm:"column:p_id;primaryKey;autoIncrement"`
	// to test Preload:
	Tasks []*TaskLi `gorm:"foreignKey:PId;references:PId"`
}

func (t *ProjectEx) TableName() string {
	return "projects"
}

type Task struct {
	TaskBase
	// to test AutoMigrate:
	Project Project `gorm:"foreignKey:PId;references:PId"`
}

func (t *TaskLi) TableName() string {
	return "tasks"
}
