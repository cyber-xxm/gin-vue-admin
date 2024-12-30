package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/example"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/example"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewCustomerApi(db *gorm.DB) *CustomerApi {
	return &CustomerApi{
		customerService: service.NewCustomerService(db),
	}
}

type CustomerApi struct {
	customerService *service.CustomerService
}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (s *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = s.customerService.CreateExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (s *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.customerService.DeleteExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (s *CustomerApi) UpdateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.customerService.UpdateExaCustomer(&customer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
func (s *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindQuery(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := s.customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaCustomerResponse{Customer: data}, "获取成功", c)
}

// GetExaCustomerList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (s *CustomerApi) GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := s.customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
