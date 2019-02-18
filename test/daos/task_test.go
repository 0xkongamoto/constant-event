package daos_test

import (
	"fmt"
	"testing"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"

	_ "github.com/go-sql-driver/mysql"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TaskDAOTestSuite struct {
	suite.Suite
	taskDAO *daos.TaskDAO
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (task *TaskDAOTestSuite) SetupTest() {
	// init daos
	if db := models.Database(); db == nil {
		panic("Db cannot connect")
	}
}

func (suite *TaskDAOTestSuite) TestAddNewTask() {
	newTask := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTask, tx)
		return err
	})

	suite.Nil(errTx)
	delete := newTask.Deleted
	suite.Equal(false, delete)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.DeleteAll(fmt.Sprintf("id = %v", newTask.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *TaskDAOTestSuite) TestUpdateTask() {
	newTask := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTask, tx)
		return err
	})

	suite.Nil(errTx)
	newContractName := "contract2"
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		newTask.ContractName = newContractName
		err := suite.taskDAO.Update(newTask, tx)
		return err
	})

	suite.Nil(errTx)
	suite.Equal(newTask.ContractName, newContractName)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.DeleteAll(fmt.Sprintf("id = %v", newTask.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *TaskDAOTestSuite) TestMultiUpdateStatusByID() {
	newTask1 := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTask1, tx)
		return err
	})

	suite.Nil(errTx)

	newTask2 := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTask2, tx)
		return err
	})

	suite.Nil(errTx)

	arrID := []uint{newTask1.ID, newTask2.ID}

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.MultiUpdateStatusByID(arrID, models.TaskStatusSuccess, tx)
		return err
	})

	suite.Nil(errTx)

	task1 := models.Task{}
	task2 := models.Task{}
	models.Database().Table("tasks").Where(`id=?`, newTask1.ID).First(&task1)
	models.Database().Table("tasks").Where(`id=?`, newTask2.ID).First(&task2)

	suite.Equal(task1.Status, models.TaskStatusSuccess)
	suite.Equal(task2.Status, models.TaskStatusSuccess)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d)", newTask1.ID, newTask2.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *TaskDAOTestSuite) TestGetLastIdScanning() {
	newTaskRetry := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusRetry,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTaskRetry, tx)
		return err
	})

	newTaskPending := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTaskPending, tx)
		return err
	})

	suite.Nil(errTx)

	task, errFind := suite.taskDAO.GetLastIdScanning()

	suite.Nil(errFind)

	suite.Equal(task.ID, newTaskPending.ID)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d)", newTaskRetry.ID, newTaskPending.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *TaskDAOTestSuite) TestGetTasksScanning() {
	newTaskSuccess := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusSuccess,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTaskSuccess, tx)
		return err
	})
	suite.Nil(errTx)

	newTaskRetry := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusRetry,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTaskRetry, tx)
		return err
	})
	suite.Nil(errTx)

	newTaskPending := &models.Task{
		TaskType:        models.TaskTypeConstant,
		Action:          models.TaskActionSell,
		Status:          models.TaskStatusPending,
		MasterAddress:   "0x1234",
		ContractAddress: "0x4567",
		ContractName:    "contract1",
		Data:            "{}",
		Deleted:         false,
	}
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.New(newTaskPending, tx)
		return err
	})

	suite.Nil(errTx)

	tasks, errFind := suite.taskDAO.GetTasksScanning(0, 10)

	suite.Nil(errFind)

	suite.Equal(2, len(tasks))

	for _, task := range tasks {
		if task.ID != newTaskRetry.ID && task.ID != newTaskPending.ID {
			suite.Equal(true, false)
		}
	}

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.taskDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d, %d)", newTaskSuccess.ID, newTaskRetry.ID, newTaskPending.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func TestTaskDAOTestSuite(t *testing.T) {
	suite.Run(t, new(TaskDAOTestSuite))
}
