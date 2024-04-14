package model

import (
	"errors"
	"go-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null " json:"name"`
}

// 查询分类

func CheckCategory(categoryName string) int {
	var cate Category
	db.Select("id").Where("name = ?", categoryName).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED // 2001 检查不通过，重名
	}
	return errmsg.SUCCESS // 200  检查通过，无重名
}

// 新增分类

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 查询单个分类信息

func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// SearchCategory 搜索分类
func SearchCategory(name string, pageSize int, pageNum int) ([]Category, int, int64) {
	var cateList []Category
	var err error
	var total int64

	err = db.
		Select("category.id, name").Where("name LIKE ?",
		name+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateList).Error
	//单独计数
	db.Model(&cateList).Where("name LIKE ?",
		name+"%",
	).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return cateList, errmsg.SUCCESS, total
}

// 查询分类列表

func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64 // 统计总数
	// Limit 表示一页有多少个用户，Offset 是固定写法
	err = db.Find(&cate).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	// gorm 还需要判断一下没有记录 err != gorm.ErrRecordNotFound 这个是旧方法
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, 0
	}
	return cate, total
}

// 编辑分类

func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// 删除分类

func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
