package dbal

import (
	"context"
	"gorm.io/gorm"
	"sdm_demo_todolist/gorm/models"
)

// Hand coded additions

func (dao *TasksDao) _ReadProjectTasks1(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	err = dao.ds.Session(ctx).Table("tasks").
		Select("t_id", "t_date", "t_subject", "t_priority").
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks4(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT * FROM `tasks` WHERE `tasks`.`p_id` = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).Model(&models.TaskLi{}).
		Where("`tasks`.`p_id` = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks5(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).Model(&models.TaskLi{}).
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks6(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT * FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks7(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT `t_id`,`t_date`,`t_subject`,`t_priority`
	// FROM `tasks` WHERE p_id = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).
		Select("t_id", "t_date", "t_subject", "t_priority").
		Where("p_id = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) __ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	// SELECT `tasks`.`t_id`,`tasks`.`p_id`,`tasks`.`t_priority`,`tasks`.`t_date`,`tasks`.`t_subject`
	// FROM `tasks` join projects p on p.p_id=`tasks`.`p_id`
	// WHERE `tasks`.`p_id` = 2 ORDER BY t_date, t_id
	err = dao.ds.Session(ctx).
		Joins("join projects p on p.p_id=`tasks`.`p_id`"). // join avoids "SELECT * ..."
		Where("`tasks`.`p_id` = ?", pId).
		Order("t_date, t_id").Find(&res).Error
	return
}

func (dao *TasksDao) _ReadProjectTasks(ctx context.Context, pId int64) (res []*models.TaskLi, err error) {
	var project models.TaskProject
	err = dao.ds.Session(ctx).Model(&models.Project{}).
		Preload(models.RefTasks, func(db *gorm.DB) *gorm.DB {
			return db.Select("t_id", "t_date", "t_subject", "t_priority"). // Preload defaults is "SELECT * ..."
											Order("t_date, t_id")
		}).
		Where("p_id = ?", pId).Take(&project).Error
	if err != nil {
		return
	}
	res = project.Tasks
	return
}

func (dao *TasksDao) ReadProjectTasks(ctx context.Context, pId int64) (res []*models.ProjectTaskLi, err error) {
	// SELECT `tasks`.`t_id`,`tasks`.`p_id`,`tasks`.`t_priority`,`tasks`.`t_date`,`tasks`.`t_subject`,`Project`.`p_id` AS `Project__p_id`
	// FROM `tasks` INNER JOIN `projects` `Project` ON `tasks`.`p_id` = `Project`.`p_id`
	// WHERE tasks.p_id=2 ORDER BY tasks.t_date, tasks.t_id
	err = dao.ds.Session(ctx).
		Joins(models.FkProject). // join avoids "SELECT * ..."
		Where("tasks.p_id=?", pId).
		Order("tasks.t_date, tasks.t_id").Find(&res).Error
	return
}
