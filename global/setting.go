package global

import (
	"nicetry/pkg/logger"
	"nicetry/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting	*setting.CacheSettingS
	Logger          *logger.Logger
)
