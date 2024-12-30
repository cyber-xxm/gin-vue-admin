package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/example"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"gorm.io/gorm"
)

func NewCustomerService(db *gorm.DB, authorityService *service.AuthorityService) *CustomerService {
	return &CustomerService{
		db:               db,
		authorityService: authorityService,
	}
}

type CustomerService struct {
	db               *gorm.DB
	authorityService *service.AuthorityService
}

var CustomerServiceApp = new(CustomerService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (s *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = s.db.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (s *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = s.db.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (s *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = s.db.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (s *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = s.db.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (s *CustomerService) GetCustomerInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := s.db.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := s.authorityService.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
