package models

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/orm"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/config/oss"
)

type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Mongo     Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Email     Email   `mapstructure:"email" json:"email" yaml:"email"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  orm.Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  orm.Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  orm.Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle orm.Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite orm.Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []orm.SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local        oss.Local        `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu        oss.Qiniu        `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS    oss.AliyunOSS    `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs    oss.HuaWeiObs    `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS   oss.TencentCOS   `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3        oss.AwsS3        `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	CloudflareR2 oss.CloudflareR2 `mapstructure:"cloudflare-r2" json:"cloudflare-r2" yaml:"cloudflare-r2"`
	Minio        oss.Minio        `mapstructure:"minio" json:"minio" yaml:"minio"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
