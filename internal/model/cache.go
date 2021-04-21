package model

type Cache interface {
	GetCachePrefix() string
}

func GetCachePreName(c Cache) string {
	return c.GetCachePrefix()
}