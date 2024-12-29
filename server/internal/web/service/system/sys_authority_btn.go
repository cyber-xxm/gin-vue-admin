package system

import (
	"errors"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	sysResp "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"gorm.io/gorm"
)

func NewAuthorityBtnService(db *gorm.DB) *AuthorityBtnService {
	return &AuthorityBtnService{
		db: db,
	}
}

type AuthorityBtnService struct {
	db *gorm.DB
}

func (s *AuthorityBtnService) GetAuthorityBtn(req system2.SysAuthorityBtnReq) (res sysResp.SysAuthorityBtnRes, err error) {
	var authorityBtn []system.SysAuthorityBtn
	err = s.db.Find(&authorityBtn, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	res.Selected = selected
	return res, err
}

func (s *AuthorityBtnService) SetAuthorityBtn(req system2.SysAuthorityBtnReq) (err error) {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var authorityBtn []system.SysAuthorityBtn
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityId:      req.AuthorityId,
				SysMenuID:        req.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	fErr := s.db.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("此按钮正在被使用无法删除")
}
