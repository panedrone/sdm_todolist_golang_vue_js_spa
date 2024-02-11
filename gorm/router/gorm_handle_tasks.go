package router

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	m "sdm_demo_todolist/gorm/models"
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
	var req request.NewTask
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	t := m.Task{}
	t.PId = uri.PId
	t.TSubject = req.TSubject
	t.TPriority = 1
	t.TDate = datetime.NowLocalString()
	if err := tDao.CreateTaskBase(ctx, &t.TaskBase); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func TaskRead(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	task, err := tDao.ReadTaskBase(ctx, uri.TId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, task)
}

func TasksReadByProject(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	// tasks, err := tDao.RawProjectTasks(ctx, uri.PId)
	tasks, err := tDao.ReadProjectTasks(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, tasks)
}

func TaskUpdate(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	dao := tDao
	t, err := dao.ReadTaskBase(ctx, uri.TId)
	if err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
		return
	}
	var req m.Task
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	_, err = datetime.Parse(req.TDate)
	if err != nil {
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
	t.TSubject = req.TSubject
	t.TPriority = req.TPriority
	t.TDate = req.TDate
	t.TComments = req.TComments
	if _, err = dao.UpdateTaskBase(ctx, t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
}

func TaskDelete(ctx *gin.Context) {
	var uri request.TaskUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	if _, err := tDao.DeleteTaskBase(ctx, &m.TaskBase{TId: uri.TId}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
