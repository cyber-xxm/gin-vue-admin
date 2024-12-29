package system

import (
	"errors"
	"github.com/cyber-xxm/gin-vue-admin/global"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	sysResp "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"strconv"

	"gorm.io/gorm"
)

func NewAuthorityService(db *gorm.DB) *AuthorityService {
	return &AuthorityService{
		db:            db,
		casbinService: NewCasbinService(db),
		menuService:   NewMenuService(db),
	}
}

type AuthorityService struct {
	db            *gorm.DB
	casbinService *CasbinService
	menuService   *MenuService
}

var ErrRoleExistence = errors.New("存在相同角色id")

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

func (s *AuthorityService) CreateAuthority(auth system2.SysAuthority) (authority system2.SysAuthority, err error) {

	if err = s.db.Where("authority_id = ?", auth.AuthorityId).First(&system2.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}

	e := s.db.Transaction(func(tx *gorm.DB) error {

		if err = tx.Create(&auth).Error; err != nil {
			return err
		}

		auth.SysBaseMenus = system.DefaultMenu()
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := system.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return s.casbinService.AddPolicies(tx, rules)
	})

	return auth, e
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo system.SysAuthorityCopyResponse
//@return: authority system.SysAuthority, err error

func (s *AuthorityService) CopyAuthority(adminAuthorityID uint, copyInfo sysResp.SysAuthorityCopyResponse) (authority system2.SysAuthority, err error) {
	var authorityBox system2.SysAuthority
	if !errors.Is(s.db.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system2.SysAuthority{}
	menus, err := s.menuService.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system2.SysBaseMenu
	for _, v := range menus {
		intNum := v.MenuId
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = s.db.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}

	var btns []system2.SysAuthorityBtn

	err = s.db.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityId).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		err = s.db.Create(&btns).Error

		if err != nil {
			return
		}
	}
	paths := s.casbinService.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = s.casbinService.UpdateCasbin(adminAuthorityID, copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = s.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

func (s *AuthorityService) UpdateAuthority(auth system2.SysAuthority) (authority system2.SysAuthority, err error) {
	var oldAuthority system2.SysAuthority
	err = s.db.Where("authority_id = ?", auth.AuthorityId).First(&oldAuthority).Error
	if err != nil {
		zap_logger.Debug(err.Error())
		return system2.SysAuthority{}, errors.New("查询角色数据失败")
	}
	err = s.db.Model(&oldAuthority).Updates(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (s *AuthorityService) DeleteAuthority(auth *system2.SysAuthority) error {
	if errors.Is(s.db.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(s.db.Where("authority_id = ?", auth.AuthorityId).First(&system2.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(s.db.Where("parent_id = ?", auth.AuthorityId).First(&system2.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}

		if len(auth.SysBaseMenus) > 0 {
			if err = tx.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus); err != nil {
				return err
			}
			// err = db.Association("SysBaseMenus").Delete(&auth)
		}
		if len(auth.DataAuthorityId) > 0 {
			if err = tx.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId); err != nil {
				return err
			}
		}

		if err = tx.Delete(&system2.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}
		if err = tx.Where("authority_id = ?", auth.AuthorityId).Delete(&[]system2.SysAuthorityBtn{}).Error; err != nil {
			return err
		}

		authorityId := strconv.Itoa(int(auth.AuthorityId))

		if err = s.casbinService.RemoveFilteredPolicy(tx, authorityId); err != nil {
			return err
		}

		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (s *AuthorityService) GetAuthorityInfoList(authorityID uint) (list []system2.SysAuthority, err error) {
	var authority system2.SysAuthority
	err = s.db.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return nil, err
	}
	var authorities []system2.SysAuthority
	db := s.db.Model(&system2.SysAuthority{})
	if global.GVA_CONFIG.System.UseStrictAuth {
		// 当开启了严格树形结构后
		if *authority.ParentId == 0 {
			// 只有顶级角色可以修改自己的权限和以下权限
			err = db.Preload("DataAuthorityId").Where("authority_id = ?", authorityID).Find(&authorities).Error
		} else {
			// 非顶级角色只能修改以下权限
			err = db.Debug().Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error
		}
	} else {
		err = db.Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authorities).Error
	}

	for k := range authorities {
		err = s.findChildrenAuthority(&authorities[k])
	}
	return authorities, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (s *AuthorityService) GetStructAuthorityList(authorityID uint) (list []uint, err error) {
	var auth system2.SysAuthority
	_ = s.db.First(&auth, "authority_id = ?", authorityID).Error
	var authorities []system2.SysAuthority
	err = s.db.Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error
	if len(authorities) > 0 {
		for k := range authorities {
			list = append(list, authorities[k].AuthorityId)
			_, err = s.GetStructAuthorityList(authorities[k].AuthorityId)
		}
	}
	if *auth.ParentId == 0 {
		list = append(list, authorityID)
	}
	return list, err
}

func (s *AuthorityService) CheckAuthorityIDAuth(authorityID, targetID uint) (err error) {
	if !global.GVA_CONFIG.System.UseStrictAuth {
		return nil
	}
	authIDS, err := s.GetStructAuthorityList(authorityID)
	if err != nil {
		return err
	}
	hasAuth := false
	for _, v := range authIDS {
		if v == targetID {
			hasAuth = true
			break
		}
	}
	if !hasAuth {
		return errors.New("您提交的角色ID不合法")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: sa system.SysAuthority, err error

func (s *AuthorityService) GetAuthorityInfo(auth system2.SysAuthority) (sa system2.SysAuthority, err error) {
	err = s.db.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error

func (s *AuthorityService) SetDataAuthority(adminAuthorityID uint, auth system2.SysAuthority) error {
	var checkIDs []uint
	checkIDs = append(checkIDs, auth.AuthorityId)
	for i := range auth.DataAuthorityId {
		checkIDs = append(checkIDs, auth.DataAuthorityId[i].AuthorityId)
	}

	for i := range checkIDs {
		err := s.CheckAuthorityIDAuth(adminAuthorityID, checkIDs[i])
		if err != nil {
			return err
		}
	}

	var sa system2.SysAuthority
	s.db.Preload("DataAuthorityId").First(&sa, "authority_id = ?", auth.AuthorityId)
	err := s.db.Model(&sa).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (s *AuthorityService) SetMenuAuthority(auth *system2.SysAuthority) error {
	var sa system2.SysAuthority
	s.db.Preload("SysBaseMenus").First(&sa, "authority_id = ?", auth.AuthorityId)
	err := s.db.Model(&sa).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (s *AuthorityService) findChildrenAuthority(authority *system2.SysAuthority) (err error) {
	err = s.db.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = s.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (s *AuthorityService) GetParentAuthorityID(authorityID uint) (parentID uint, err error) {
	var authority system2.SysAuthority
	err = s.db.Where("authority_id = ?", authorityID).First(&authority).Error
	return *authority.ParentId, err
}
