package system

import (
	"errors"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"gorm.io/gorm"
)

func NewDictionaryService(db *gorm.DB) *DictionaryService {
	return &DictionaryService{
		db: db,
	}
}

type DictionaryService struct {
	db *gorm.DB
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (s *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(s.db.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = s.db.Create(&sysDictionary).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (s *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = s.db.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = s.db.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	if sysDictionary.SysDictionaryDetails != nil {
		return s.db.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *model.SysDictionary
//@return: err error

func (s *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	err = s.db.Where("id = ?", sysDictionary.ID).First(&dict).Error
	if err != nil {
		zap_logger.Debug(err.Error())
		return errors.New("查询字典数据失败")
	}
	if dict.Type != sysDictionary.Type {
		if !errors.Is(s.db.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = s.db.Model(&dict).Updates(sysDictionaryMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary model.SysDictionary

func (s *DictionaryService) GetSysDictionary(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = s.db.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (s *DictionaryService) GetSysDictionaryInfoList() (list interface{}, err error) {
	var sysDictionarys []system.SysDictionary
	err = s.db.Find(&sysDictionarys).Error
	return sysDictionarys, err
}
