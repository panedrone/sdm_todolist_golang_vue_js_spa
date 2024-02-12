package models

// Hand coded additions

const RefTasks = "Tasks" // this is not FK

const FkProject = "Project"

// Task to test AutoMigrate
type Task struct {
	TaskBase
	Project Project `gorm:"foreignKey:PId;references:PId"`
}

// TaskProject to test Joins and Preload
type TaskProject struct {
	PId   int64     `gorm:"column:p_id;primaryKey"`
	Tasks []*TaskLi `gorm:"foreignKey:PId;references:PId"`
}

func (t *TaskProject) TableName() string {
	return "projects"
}

// ProjectTaskLi to test Joins
type ProjectTaskLi struct {
	TaskLi
	Project TaskProject `gorm:"foreignKey:PId;references:PId"`
}

func (t *ProjectTaskLi) TableName() string {
	return "tasks"
}
