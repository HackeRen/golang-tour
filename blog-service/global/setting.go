package global

import (
	"github.com/golang-tour/blog-service/pkg/logger"
	"github.com/golang-tour/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
