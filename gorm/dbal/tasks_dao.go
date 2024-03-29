package dbal

// Code generated by a tool. DO NOT EDIT.
// Additional custom methods can be implemented in a separate file like <this_file>_ex.go.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/gorm/dbal/models"
)

type TasksDao struct {
	ds DataStore
}

// (C)RUD: tasks
// Generated/AI values are passed to DTO/model.

func (dao *TasksDao) CreateTaskBase(ctx context.Context, item *models.TaskBase) error {
	return dao.ds.Create(ctx, "tasks", item)
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTaskBaseList(ctx context.Context) (res []*models.TaskBase, err error) {
	err = dao.ds.ReadAll(ctx, "tasks", &res)
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTaskBase(ctx context.Context, tId int64) (*models.TaskBase, error) {
	res := &models.TaskBase{}
	err := dao.ds.Read(ctx, "tasks", res, tId)
	if err == nil {
		return res, nil
	}
	return nil, err
}

// CR(U)D: tasks

func (dao *TasksDao) UpdateTaskBase(ctx context.Context, item *models.TaskBase) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Update(ctx, "tasks", item)
	return
}

// CRU(D): tasks

func (dao *TasksDao) DeleteTaskBase(ctx context.Context, item *models.TaskBase) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Delete(ctx, "tasks", item)
	return
}

func (dao *TasksDao) RawProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	sql := `select t_id, t_priority, t_date, t_subject from tasks where p_id =? 
		order by t_id`
	err = dao.ds.Select(ctx, sql, &res, pId)
	return
}
