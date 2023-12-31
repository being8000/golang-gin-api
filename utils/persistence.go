package utils

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	_ "github.com/go-sql-driver/mysql"
)

var (
	GlobalCache *bigcache.BigCache
)

func init() {
	// Connect to DB
	//var err error
	//DB, err = gorm.Open("mysql", "your_db_url")
	//if err != nil {
	//	panic(fmt.Errorf("failed to connect to DB: %w", err))
	//}

	// Initialize cache
	var err error
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}
