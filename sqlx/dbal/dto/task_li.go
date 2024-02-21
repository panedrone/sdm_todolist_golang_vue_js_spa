package dto

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

// TaskLi `task` list item: [-] p_id, [-] t_comments
type TaskLi struct {
	TId       int64  `json:"t_id" db:"t_id"`
	TPriority int64  `json:"t_priority" db:"t_priority"`
	TDate     string `json:"t_date" db:"t_date"`
	TSubject  string `json:"t_subject" db:"t_subject"`
}
