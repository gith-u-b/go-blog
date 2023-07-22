package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag, err error) {
	err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error

	return tags, err
}

func GetTagTotal(maps interface {}) (count int, err error){
	err = db.Model(&Tag{}).Where(maps).Count(&count).Error

	return count, err
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error

	if err != nil{
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag

	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil

}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func AddTag(name string, state int, createdBy string) error {
	err := db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	}).Error

	return err
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}

//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}