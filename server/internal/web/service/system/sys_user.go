package system

import (
	"errors"
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"

	"time"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db:               db,
		menuService:      NewMenuService(db),
		authorityService: NewAuthorityService(db),
	}
}

type UserService struct {
	db               *gorm.DB
	menuService      *MenuService
	authorityService *AuthorityService
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

func (s *UserService) Register(u system2.SysUser) (userInter system2.SysUser, err error) {
	var user system2.SysUser
	if !errors.Is(s.db.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = s.db.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (s *UserService) Login(u *system2.SysUser) (userInter *system2.SysUser, err error) {
	if nil == s.db {
		return nil, fmt.Errorf("db not init")
	}

	var user system2.SysUser
	err = s.db.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		s.menuService.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (s *UserService) ChangePassword(u *system2.SysUser, newPassword string) (userInter *system2.SysUser, err error) {
	var user system2.SysUser
	if err = s.db.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = s.db.Save(&user).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (s *UserService) GetUserInfoList(info system.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := s.db.Model(&system2.SysUser{})
	var userList []system2.SysUser

	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (s *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {

	assignErr := s.db.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system2.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	var authority system2.SysAuthority
	err = s.db.Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return err
	}
	var authorityMenu []system2.SysAuthorityMenu
	var authorityMenuIDs []string
	err = s.db.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
	if err != nil {
		return err
	}

	for i := range authorityMenu {
		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
	}

	var authorityMenus []system2.SysBaseMenu
	err = s.db.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
	if err != nil {
		return err
	}
	hasMenu := false
	for i := range authorityMenus {
		if authorityMenus[i].Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return errors.New("找不到默认路由,无法切换本角色")
	}

	err = s.db.Model(&system2.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (s *UserService) SetUserAuthorities(adminAuthorityID, id uint, authorityIds []uint) (err error) {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var user system2.SysUser
		TxErr := tx.Where("id = ?", id).First(&user).Error
		if TxErr != nil {
			zap_logger.Debug(TxErr.Error())
			return errors.New("查询用户数据失败")
		}
		TxErr = tx.Delete(&[]system2.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system2.SysUserAuthority
		for _, v := range authorityIds {
			e := s.authorityService.CheckAuthorityIDAuth(adminAuthorityID, v)
			if e != nil {
				return e
			}
			useAuthority = append(useAuthority, system2.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Model(&user).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (s *UserService) DeleteUser(id int) (err error) {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system2.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system2.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (s *UserService) SetUserInfo(req system2.SysUser) error {
	return s.db.Model(&system2.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (s *UserService) SetSelfInfo(req system2.SysUser) error {
	return s.db.Model(&system2.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfSetting
//@description: 设置用户配置
//@param: req datatypes.JSON, uid uint
//@return: err error

func (s *UserService) SetSelfSetting(req models.JSONMap, uid uint) error {
	return s.db.Model(&system2.SysUser{}).Where("id = ?", uid).Update("origin_setting", req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (s *UserService) GetUserInfo(uuid uuid.UUID) (user system2.SysUser, err error) {
	var reqUser system2.SysUser
	err = s.db.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	s.menuService.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (s *UserService) FindUserById(id int) (user *system2.SysUser, err error) {
	var u system2.SysUser
	err = s.db.Where("id = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (s *UserService) FindUserByUuid(uuid string) (user *system2.SysUser, err error) {
	var u system2.SysUser
	if err = s.db.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ResetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (s *UserService) ResetPassword(ID uint) (err error) {
	err = s.db.Model(&system2.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}