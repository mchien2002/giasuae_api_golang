package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctg *entities.Category) error
	UpdateCategory(ctg *entities.Category) error
	DeleteCategory(ctg *entities.Category) error
	FindAllCategory() []entities.Category
	FindByID(id int) entities.Category
	FilterCategory(value ...interface{}) entities.Category
}

type categoryConnection struct {
	connection *gorm.DB
}

// FilterCategory implements CategoryRepository
func (*categoryConnection) FilterCategory(value ...interface{}) entities.Category {
	panic("unimplemented")
}

// DeleteCategory implements CategoryRepository
func (*categoryConnection) DeleteCategory(ctg *entities.Category) error {
	panic("unimplemented")
}

// FindAllCategory implements CategoryRepository
func (db *categoryConnection) FindAllCategory() []entities.Category {
	var categories []entities.Category
	db.connection.Find(&categories)
	return categories
}

// FindByID implements CategoryRepository
func (db *categoryConnection) FindByID(id int) entities.Category {
	var category entities.Category
	db.connection.First(&category, id)
	return category
}

// InsertCategory implements CategoryRepository
func (db *categoryConnection) InsertCategory(ctg *entities.Category) error {
	err := db.connection.Save(&ctg)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateCategory implements CategoryRepository
func (db *categoryConnection) UpdateCategory(ctg *entities.Category) error {
	err := db.connection.Save(&ctg)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbConn,
	}
}
