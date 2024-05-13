package repositories

import (
	"github.com/iqbalmahad/todolist-app-golang.git/models"
	"gorm.io/gorm"
)

type TodoRepoImpl struct {
	db *gorm.DB // GORM database connection
}

func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &TodoRepoImpl{db: db}
}

func (tr *TodoRepoImpl) Create(todo models.Todo) error {
	return tr.db.Create(&todo).Error
}

func (tr *TodoRepoImpl) Read(id string) (models.Todo, error) {
	var todo models.Todo
	if err := tr.db.Where("id = ?", id).First(&todo).Error; err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (tr *TodoRepoImpl) Update(todo models.Todo) error {
	return tr.db.Model(&todo).Updates(todo).Error
}

func (tr *TodoRepoImpl) Delete(id string) error {
	return tr.db.Where("id = ?", id).Delete(&models.Todo{}).Error
}

func (tr *TodoRepoImpl) Index() ([]models.Todo, error) {
	var todos []models.Todo
	if err := tr.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
