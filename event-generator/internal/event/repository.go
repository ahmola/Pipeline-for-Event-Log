package event

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
}

// Gin Base ORM Repository
type GormRepository[T any] struct {
	Repository[T]
	DB *gorm.DB
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.DB.Create(&entity).Error
}

type EventRepository struct {
	GormRepository[Event]
}
