package global

import (
	"github.com/zhouya0/blog/pkg/logger"
	"github.com/zhouya0/blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger *logger.Logger
)
