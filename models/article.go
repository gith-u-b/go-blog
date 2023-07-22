package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
}


func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface {}) (count int, err error){
	err = db.Model(&Article{}).Where(maps).Count(&count).Error

	if err != nil {
		return 0, xerrors.Errorf("%v", err)
	}

	return
}

func GetArticles(pageNum int, pageSize int, maps interface {}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error

	if err != nil {
		return articles, xerrors.Errorf("%v", err)
	}

	return articles, nil
}

func GetArticle(id int) (*Article, error){
	var article Article
	err := db.Where("id = ?", id).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, xerrors.Errorf("%v", err)
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, xerrors.Errorf("%v", err)
	}

	return &article, nil
}

func EditArticle(id int, data interface {}) error {
	err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error

	if err != nil{
		return xerrors.Errorf("%v", err)
	}


	return nil
}

func AddArticle(data map[string]interface {}) error {
	db.Create(&Article {
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})

	return nil
}

func DeleteArticle(id int) error {
	err := db.Where("id = ?", id).Delete(Article{}).Error

	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	return nil
}

func CleanAllArticle() error {
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error

	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	return nil
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}
