package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	State         int    `json:"state"`
	TagId         int    `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"Content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	Tag           Tag    `json:"tag"`
}

//func ExistArticleByID(id int) bool {
//	var article Article
//	db.Select("id").Where("id = ?", id).First(&article)
//	if article.ID > 0 {
//		return true
//	}
//	return false
//}
//
//func GetArticleTotal(maps interface{}) (count int) {
//	db.Model(&Article{}).Where(maps).Count(&count)
//	return
//}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (tag *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

//func GetArticle(id int) (article Article) {
//	db.Where("id = ?", id).First(&article)
//	db.Where("id = ?", article.TagId).First(&article.Tag)
//	//db.Model(&article).Related(&article.Tag)
//	return
//}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

//func GetArticles(PageNum int, PageSize int, maps interface{}) (article []Article) {
//	db.Where(maps).Offset(PageNum).Limit(PageSize).Find(&article)
//	return
//}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

func AddArticle(data map[string]interface{}) error {
	db.Create(&Article{
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return nil
}

func EditArticle(id int, maps interface{}) error {
	db.Model(&Article{}).Where("id = ?", id).Update(maps)
	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}
