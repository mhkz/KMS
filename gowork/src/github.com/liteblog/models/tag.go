package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
/**
 * gorm 的 Callbacks 可以将回调方法定义为模型结构的指针， 在创建，更新， 查询，删除是被调用
 * 如果任何回调返回错误， gorm将停止未来操作并回滚所有更改
 * gorm 支持的回调方法
 * 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
 * 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
 * 删除：BeforeDelete、AfterDelete
 * 查询：AfterFind
 */
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags [] Tag) {
	// 在同个models包下， db *gorm.DB是可以直接使用的
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	// 函数显示的声明了返回值，因此可以直接 return
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag {
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}


func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface {}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

http://test.restoplus.cn/wechat/pay/fixedRefund?brandId=b92323e7c0274617b755dc47c4c7c657&shopId=d0405171498d4555865c5e8fd0f10d15&total=53600&
refund=53600&transaction_id=4200000281201903140474266909&mchid=1391463202&id=1