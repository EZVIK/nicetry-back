package global

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	CacheEngine *redis.Client
)


func Tx(funcs ...func(db *gorm.DB) error) (err error) {
	tx := DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", err)
		}
	}()
	for _, f := range funcs {
		err = f(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

func main() {

}