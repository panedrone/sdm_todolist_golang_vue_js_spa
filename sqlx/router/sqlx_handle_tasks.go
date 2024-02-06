package router

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/shared/datetime"
	"sdm_demo_todolist/shared/request"
	"sdm_demo_todolist/shared/resp"
	"sdm_demo_todolist/sqlx/dbal"
	"sdm_demo_todolist/sqlx/dbal/dto"
)

var (
	tDao = dbal.NewTasksDao()
)

// TaskCreate
//
//	@Summary	create task
//	@Tags		Tasks
//	@Id			TaskCreate
//	@Accept		json
//	@Param		json	body	request.NewTask	true	"task data"
//	@Success	201
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/projects/{p_id}/tasks [post]
//	@Param		p_id	path	integer	true	"project id"
func TaskCreate(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	var inTask request.NewTask
	if err := ctx.ShouldBindJSON(&inTask); err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
		return
	}
	t := dto.Task{}
	t.PId = uri.PId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	t.TDate = datetime.NowLocalString()
	if err := tDao.CreateTask(ctx, &t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

// TasksReadByProject
//
//	@Summary	get project tasks
//	@Tags		Tasks
//	@Id			TasksReadByProject
//	@Produce	json
//	@Success	200	{object}	[]dto.TaskLi		"project tasks"
//	@Failure	500
//	@Security	none
//	@Router		/projects/{p_id}/tasks [get]
//	@Param		p_id	path	integer	true	"project id"
func TasksReadByProject(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	tasks, err := tDao.GetGroupTasks(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.RespondWithJSON(ctx, http.StatusOK, tasks)
}

// TaskRead
//
//	@Summary	get task
//	@Tags		Tasks
//	@Id			TaskRead
//	@Produce	json
//	@Success	200	{object}	dto.Task	"task data"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/tasks/{t_id} [get]
//	@Param		t_id	path	integer	true	"task id"
func TaskRead(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	task, err := tDao.ReadTask(ctx, uri.TId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.RespondWithJSON(ctx, http.StatusOK, task)
}

// TaskUpdate
//
//	@Summary	update task
//	@Tags		Tasks
//	@Id			TaskUpdate
//	@Accept		json
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/tasks/{t_id} [put]
//	@Param		t_id	path	integer		true	"task id"
//	@Param		json	body	dto.Task	true	"task data"
func TaskUpdate(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	var req dto.Task
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if _, err := datetime.Parse(req.TDate); err != nil {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Date format expected '%s': %s", datetime.TimeFormat, err.Error()))
		return
	}
	if len(req.TSubject) == 0 {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Subject required"))
		return
	}
	if req.TPriority <= 0 {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Invalid Priority: %d", req.TPriority))
		return
	}
	t, err := tDao.ReadTask(ctx, uri.TId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	t.TSubject = req.TSubject
	t.TPriority = req.TPriority
	t.TDate = req.TDate
	t.TComments = req.TComments
	if _, err = tDao.UpdateTask(ctx, t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
}

// TaskDelete
//
//	@Summary	delete task
//	@Tags		Tasks
//	@Id			TaskDelete
//	@Success	204
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/tasks/{t_id} [delete]
//	@Param		t_id	path	integer	true	"task id"
func TaskDelete(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	t := dto.Task{TId: uri.TId}
	if _, err := tDao.DeleteTask(ctx, &t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
