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

//type TaskProjectFK struct {
//	PId int64 `gorm:"column:p_id;not null"`
//}
//
//func (t *TaskProjectFK) TableName() string {
//	return "projects"
//}

type ProjectTaskLi struct {
	TaskLi
	ProjectId int64 `gorm:"column:p_id;not null"`
	// Project   TaskProjectFK `gorm:"ForeignKey:ProjectId;references:PId"`
}

const (

	// "RefTasks" is the name of field "ProjectWithTasks.RefTasks".
	// It is used for "Preload" test.

	RefProjectTasks = "RefTasks"
)

type ProjectWithTasks struct {
	Project
	RefTasks []*ProjectTaskLi `gorm:"ForeignKey:ProjectId;references:PId"`
}

func (t *ProjectWithTasks) TableName() string {
	return "projects"
}
