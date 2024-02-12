package dbal

import (
	"context"
	"gorm.io/gorm"
	"sdm_demo_todolist/gorm/models"
)

// Hand coded additions

func (dao *TasksDao) ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {

	// SELECT `tasks`.`t_id`,`tasks`.`t_priority`,`tasks`.`t_date`,`tasks`.`t_subject` FROM `tasks` WHERE p_id = 1 ORDER BY t_date, t_id

	where := &models.Task{
		TaskBase: models.TaskBase{PId: pId},
	}

	// err = dao.ds.Session(ctx). // "Table" and "Model" are not needed if "TaskLi" has "TableName".
	// err = dao.ds.Session(ctx).Table("tasks"). // Use "Table" or "Model" if "TaskLi" has no "TableName".
	err = dao.ds.Session(ctx).Model(where).
		// Select("t_id", "t_date", "t_subject", "t_priority"). // "Select" is not needed!!!
		Where(where).
		Order("t_date, t_id").Find(&res).Error

	return
}

func (dao *TasksDao) _ReadProjectTasks2(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {

	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id

	err = dao.ds.Session(ctx).Model(&models.TaskLi{}).
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error

	return
}

func (dao *TasksDao) _ReadProjectTasks3(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {

	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id

	err = dao.ds.Session(ctx).
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error

	return
}

func (dao *TasksDao) _ReadProjectTasks4(ctx context.Context, pId int64) (res []*models.ProjectTaskLi, err error) {

	// SELECT `t_id`,`p_id`,`t_date`,`t_subject`,`t_priority` FROM `tasks` WHERE `tasks`.`p_id` = 5 ORDER BY t_date, t_id

	var project = &models.ProjectWithTasks{
		Project: models.Project{PId: pId},
	}

	err = dao.ds.Session(ctx).Model(project).Preload(models.RefProjectTasks, func(db *gorm.DB) *gorm.DB {

		// Use "Select" because "Preload" default yields " SELECT * FROM ..."

		return db.Select("t_id", "p_id", "t_date", "t_subject", "t_priority").
			Order("t_date, t_id")

	}).
		Where(project). // https://gist.github.com/WangYihang/7d43d70db432ff8f3a0a88425bfca7f2

		Take(&project).Error

	if err != nil {
		return
	}

	res = project.RefTasks

	return
}
