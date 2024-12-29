package system

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func NewJwtService(db *gorm.DB) *JwtService {
	return &JwtService{
		db: db,
	}
}

type JwtService struct {
	db *gorm.DB
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (s *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = s.db.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (s *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := s.db.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (s *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (s *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func (s *JwtService) LoadAll() {
	var data []string
	err := s.db.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		zap_logger.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
