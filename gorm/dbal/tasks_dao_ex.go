package dbal

import (
	"context"
	"gorm.io/gorm"
	"sdm_demo_todolist/gorm/models"
)

// Hand coded additions

func (dao *TasksDao) _ReadProjectTasks1(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// "Table" and "direct TaskLi" require "Select":
	//		err = dao.ds.Session(ctx).Table("tasks"). // --> "SELECT * FROM ..."
	//		err = dao.ds.Session(ctx). // --> "SELECT * FROM ..."
	// "Model" is OK with no "Select".
	model := &models.Task{
		TaskBase: models.TaskBase{PId: pId},
	}
	err = dao.ds.Session(ctx).Model(model). // --> "SELECT t_id, t_date, t_subject, t_priority FROM ..."
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(model). // // https://gist.github.com/WangYihang/7d43d70db432ff8f3a0a88425bfca7f2
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks2(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id
	model := &models.TaskLi{PId: pId}
	err = dao.ds.Session(ctx).Model(model).
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where(model).Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks3(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).
		// Select("t_id", "t_date", "t_subject", "t_priority").
		Where("p_id = ?", pId).Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// Implemented with "Preload" for educational purposes. The best one is "_ReadProjectTasks1".
	// SELECT `t_id`,`p_id`,`t_date`,`t_subject`,`t_priority` FROM `tasks` WHERE `tasks`.`p_id` = 5 ORDER BY t_date, t_id
	var model = &models.ProjectWithTasks{
		Project: models.Project{PId: pId},
	}
	err = dao.ds.Session(ctx).Model(model.Project).Preload(models.RefProjectTasks,
		func(db *gorm.DB) *gorm.DB {
			// Use "Select" because "Preload" default issues "SELECT * FROM ..."
			return db.Select("t_id", "p_id", "t_date", "t_subject", "t_priority").
				Order("t_date, t_id")
		}).
		Where(model).Take(model).Error
	if err == nil {
		res = model.RefTasks
	}
	return
}
