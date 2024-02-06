package router

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/shared/request"
	"sdm_demo_todolist/shared/resp"
	"sdm_demo_todolist/sqlx/dbal"
	"sdm_demo_todolist/sqlx/dbal/dto"
)

var (
	pDao = dbal.NewProjectsDao()
)

// ProjectCreate
//
//	@Summary	create project
//	@Tags		Projects
//	@Id			ProjectCreate
//	@Accept		json
//	@Success	201
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/projects [post]
//	@Param		json	body	request.Project	true	"project data"
func ProjectCreate(ctx *gin.Context) {
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if err := pDao.CreateProject(ctx, &dto.Project{PName: req.PName}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

// ProjectsReadAll
//
//	@Summary	get project list
//	@Tags		Projects
//	@Id			ProjectsReadAll
//	@Produce	json
//	@Success	200	{object}	[]dto.ProjectLi	"project list"
//	@Failure	500
//	@Security	none
//	@Router		/projects [get]
func ProjectsReadAll(ctx *gin.Context) {
	projects, err := pDao.GetProjects(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.RespondWithJSON(ctx, http.StatusOK, projects)
}

// ProjectRead
//
//	@Summary	get project
//	@Tags		Projects
//	@Id			ProjectRead
//	@Produce	json
//	@Success	200	{object}	dto.Project	"project data"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/projects/{p_id} [get]
//	@Param		p_id	path	integer	true	"project id"
func ProjectRead(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	pr, err := pDao.ReadProject(ctx, uri.PId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.RespondWithJSON(ctx, http.StatusOK, pr)
}

// ProjectUpdate
//
//	@Summary	update project
//	@Tags		Projects
//	@Id			ProjectUpdate
//	@Accept		json
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/projects/{p_id} [put]
//	@Param		p_id	path	integer			true	"project id"
//	@Param		json	body	request.Project	true	"project data"
func ProjectUpdate(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if _, err := pDao.UpdateProject(ctx, &dto.Project{PId: uri.PId, PName: req.PName}); err != nil {
		resp.Abort500(ctx, err)
	}
}

// ProjectDelete
//
//	@Summary	delete project
//	@Tags		Projects
//	@Id			ProjectDelete
//	@Success	204
//	@Failure	400
//	@Failure	500
//	@Security	none
//	@Router		/projects/{p_id} [delete]
//	@Param		p_id	path	integer	true	"project id"
func ProjectDelete(ctx *gin.Context) {
	var uri request.ProjectUri
	if err := request.BindUri(ctx, &uri); err != nil {
		return
	}
	if _, err := pDao.DeleteProject(ctx, &dto.Project{PId: uri.PId}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
