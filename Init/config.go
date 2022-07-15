package Init

import (
	"dousheng/config"

	"github.com/minio/minio-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//注册全局变量
var (
	Settings    config.ServerConfig //定义接收服务器配置
	DB          *gorm.DB            //注册gorm.DB
	Lg          *zap.Logger         //注册日志zp.logger
	MinioClient *minio.Client       //注册minio客户端
)
