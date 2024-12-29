package system

import (
	"errors"
	"github.com/cyber-xxm/gin-vue-admin/global"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"gorm.io/gorm"
)

func NewBaseMenuService(db *gorm.DB) *BaseMenuService {
	return &BaseMenuService{
		db: db,
	}
}

type BaseMenuService struct {
	db *gorm.DB
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func (s *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	err = s.db.First(&system2.SysBaseMenu{}, "parent_id = ?", id).Error
	if err == nil {
		return errors.New("此菜单存在子菜单不可删除")
	}
	var menu system2.SysBaseMenu
	err = s.db.First(&menu, id).Error
	if err != nil {
		return errors.New("记录不存在")
	}
	err = s.db.First(&system2.SysAuthority{}, "default_router = ?", menu.Name).Error
	if err == nil {
		return errors.New("此菜单有角色正在作为首页，不可删除")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {

		err = tx.Delete(&system2.SysBaseMenu{}, "id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system2.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system2.SysBaseMenuBtn{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&system2.SysAuthorityBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system2.SysAuthorityMenu{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		return nil
	})

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return: err error

func (s *BaseMenuService) UpdateBaseMenu(menu system2.SysBaseMenu) (err error) {
	var oldMenu system2.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["active_name"] = menu.ActiveName
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = s.db.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system2.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				zap_logger.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&system2.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			zap_logger.Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&system2.SysBaseMenuBtn{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			zap_logger.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				zap_logger.Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				zap_logger.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = tx.Model(&oldMenu).Updates(upDateMap).Error
		if txErr != nil {
			zap_logger.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: menu system.SysBaseMenu, err error

func (s *BaseMenuService) GetBaseMenuById(id int) (menu system2.SysBaseMenu, err error) {
	err = s.db.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
