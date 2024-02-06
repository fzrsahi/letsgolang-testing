package repository

import "GolandUnitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
