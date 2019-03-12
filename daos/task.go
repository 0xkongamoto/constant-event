package daos

import (
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
)

// TaskDAO : struct
type TaskDAO struct {
}

func (t *TaskDAO) GetTasksScanning(fromID uint, limit int) ([]wm.Task, error) {
	tasks := []wm.Task{}
	err := models.Database().
		Where(`
				id >= ?
				AND status = ?
			`, fromID, wm.TaskStatusPending).
		Limit(limit).
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetLastIdScanning : ...
func (t *TaskDAO) GetLastIdScanning() (wm.Task, error) {
	task := wm.Task{}
	err := models.Database().
		Select("id").
		Where(`
			status = ?
			`, wm.TaskStatusPending).
		Order("id desc").
		First(&task).Error

	if err != nil {
		return task, err
	}
	return task, nil
}

func (t *TaskDAO) Update(task *wm.Task, tx *gorm.DB) error {
	err := tx.Save(task).Error
	return err
}

func (t *TaskDAO) New(task *wm.Task, tx *gorm.DB) error {
	err := tx.Create(task).Error
	return err
}
