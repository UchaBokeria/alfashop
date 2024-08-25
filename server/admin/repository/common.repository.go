package repository

import (
	"fmt"
	"main/pkg/storage"
)

func FindByID[T any](id int) *T {
	modeler := new(T)
    if err := storage.DB.Where("id = ?", id).Last(modeler).Error; err != nil {
        // return err
		fmt.Print("some error of ", err.Error())
    }
	return modeler
}

func Create[T any](params interface{}) error {
	if result := storage.DB.Create(params); result.Error != nil {
		return result.Error
	}
	return nil
}

func Update[T any](ID int, params interface{}) error {
	entity := new(T)
	match := storage.DB.Last(&entity, uint(ID))

	if match.Error != nil {
		return match.Error
	} else if result := storage.DB.Model(entity).Updates(params); result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete[T any](conds interface{}) error {
	entity := new(T)
	storage.DB.First(entity, conds)
	if result := storage.DB.Delete(&entity); result.Error != nil {
		return result.Error
	}
	return nil
}
