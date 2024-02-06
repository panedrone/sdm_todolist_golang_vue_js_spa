package router

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/no_orm/dbal"
	"sdm_demo_todolist/no_orm/dbal/dto"
	"sdm_demo_todolist/shared/datetime"
	"sdm_demo_todolist/shared/request"
	"sdm_demo_todolist/shared/resp"
)

var (
	tDao = dbal.NewTasksDao()
)

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

func TasksReadByProject(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	tasks, err := tDao.ReadByProject(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.RespondWithJSON(ctx, http.StatusOK, tasks)
}

func TaskRead(ctx *gin.Context) {
	var uri request.TaskUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
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
