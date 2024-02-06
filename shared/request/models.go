package request

type ProjectUri struct {
	PId int64 `uri:"p_id" binding:"required"`
}

type Project struct {
	PName string `json:"p_name" binding:"required,lte=256"`
}

type TaskUri struct {
	TId int64 `uri:"t_id" binding:"required"`
}

type NewTask struct {
	TSubject string `json:"t_subject" binding:"required,lte=256"`
}
