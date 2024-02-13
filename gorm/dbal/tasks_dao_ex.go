package dbal

import (
	"context"
	"gorm.io/gorm"
	"sdm_demo_todolist/gorm/models"
)

// Hand coded additions

// 1. "Task" as "Model" and "TaskLi" as Result, no "Select" --> "SELECT t_id, t_date, t_subject, t_priority FROM ...

//      ----- this is the best one so far -----

func (dao *TasksDao) ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	queryModel := &models.Task{
		TaskBase: models.TaskBase{PId: pId},
	}
	err = dao.ds.Session(ctx).Model(queryModel).
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(queryModel). // https://gist.github.com/WangYihang/7d43d70db432ff8f3a0a88425bfca7f2
		Order("t_date, t_id").Find(&res).Error

	return
}

// 2. "TaskLi" for both "Model" and Result. Requires "Select" --> SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id

func (dao *TasksDao) _ReadProjectTasks2(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	model := &models.TaskLi{PId: pId}
	err = dao.ds.Session(ctx).Model(model).
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(model).
		Order("t_date, t_id").Find(&res).Error

	return
}

// 3. Using "Table". Requires "Select" --> SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id

func (dao *TasksDao) _ReadProjectTasks3(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	err = dao.ds.Session(ctx).Table("tasks").
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(&models.TaskLi{PId: pId}).
		Order("t_date, t_id").Find(&res).Error

	return
}

// 4. The case "direct TaskLi": no "Table", no "Model". Requires "Select" --> SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id

func (dao *TasksDao) _ReadProjectTasks4(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	err = dao.ds.Session(ctx).
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(&models.TaskLi{PId: pId}).
		Order("t_date, t_id").Find(&res).Error

	return
}

// 5. Using "Preload" for "educational purposes". Requires "Select".

func (dao *TasksDao) _ReadProjectTasks5(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	var queryModel = &models.ProjectWithTasks{
		Project: models.Project{PId: pId},
	}
	err = dao.ds.Session(ctx).Model(queryModel).Preload(models.RefProjectTasks,
		func(db *gorm.DB) *gorm.DB {
			// Use "Select" because "Preload" default issues "SELECT * FROM ..."
			return db.Select("t_id", "p_id", "t_date", "t_subject", "t_priority").
				Order("t_date, t_id")
		}).
		Where(queryModel).Take(queryModel).Error
	if err == nil {
		res = queryModel.RefTasks
	}
	return
}
