package daos

import (
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
)

type TaskDAO struct{}

func (t *TaskDAO) GetTasksScanning(fromID uint, limit int) ([]models.Task, error) {
	tasks := []models.Task{}
	err := models.Database().
		Where(`
				id >= ?
				AND (
					status = ?
					OR
					status = ?
				)
			`, fromID, models.TaskStatusPending, models.TaskStatusRetry).
		Limit(limit).
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskDAO) GetLastIdScanning() (models.Task, error) {
	task := models.Task{}
	err := models.Database().
		Select("id").
		Where(`
			status = ?
			OR
			status = ?
			`, models.TaskStatusPending, models.TaskStatusRetry).
		Order("id desc").
		First(&task).Error

	if err != nil {
		return task, err
	}
	return task, nil
}

// []int{10, 11}
func (t *TaskDAO) MultiUpdateStatusByID(ids []uint, status models.TaskStatus, tx *gorm.DB) error {
	// err := models.Database().Table("tasks").
	err := tx.Table("tasks").
		Where("id IN (?)", ids).
		Updates(map[string]interface{}{"status": status}).
		Error
	return err
}

func (t *TaskDAO) Update(task *models.Task, tx *gorm.DB) error {
	err := tx.Save(task).Error
	return err
}

func (t *TaskDAO) New(task *models.Task, tx *gorm.DB) error {
	err := tx.Create(task).Error
	return err
}

func (t *TaskDAO) DeleteAll(query string, tx *gorm.DB) error {
	// This comment only update delete_at
	// err := tx.Where(query).Delete(models.Task{}).Error
	err := tx.Unscoped().Where(query).Delete(models.Task{}).Error
	return err
}
