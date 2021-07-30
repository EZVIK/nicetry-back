package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	ImageFilePath   string
	ImageMaxSize    int64
	ImageAllowExts  []string
	ImagePrefixUrl  string
	IfCheckLike     bool
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type CacheSettingS struct {
	DBType                 string
	UserName               string
	Password               string
	Host                   string
	DBName                 string
	TablePrefix            string
	Charset                string
	ParseTime              bool
	MaxIdleConns           int
	MaxOpenConns           int
	UserPrefix             string
	REDIS_NS_AUTH          string
	REDIS_NS_USER_ID       string
	REDIS_NS_NICE_LIKES    string
	REDIS_NS_COMMENT_LIKES string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
