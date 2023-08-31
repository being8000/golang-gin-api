package database

import (
	"fmt"
	"log"
	"time"
	"zehan/gin/model"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is connected MySQL DB
type Mysql struct {
	DB *gorm.DB
}

// Connect to MySQL server
func (m *Mysql) Connect() {
	config := viper.GetStringMap("mysql")
	dsn := fmt.Sprintf(
		"%s:%d@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["port"],
		config["name"],
	)
	zap.L().Info(string(dsn))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// PrepareStmt: true,
		Logger:                                   setLogger(),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		zap.L().Fatal("Database connect error", zap.Error(err))
	}
	m.DB = db
	// db.Use(dbresolver.Register(dbresolver.Config{
	// 	// use `db2` as sources, `db3`, `db4` as replicas
	// 	Sources:  []gorm.Dialector{mysql.Open("db2_dsn")},
	// 	Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
	// 	// sources/replicas load balancing policy
	// 	Policy: dbresolver.RandomPolicy{},
	// 	// print sources/replicas mode in logger
	// 	TraceResolverMode: true,
	// }))
	// .Register(dbresolver.Config{
	// 	// use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
	// 	Replicas: []gorm.Dialector{mysql.Open("db5_dsn")},
	// }, &User{}, &Address{}).Register(dbresolver.Config{
	// 	// use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
	// 	Sources:  []gorm.Dialector{mysql.Open("db6_dsn"), mysql.Open("db7_dsn")},
	// 	Replicas: []gorm.Dialector{mysql.Open("db8_dsn")},
	// }, "orders", &Product{}, "secondary"))
}

func setLogger() logger.Interface {
	l := &lumberjack.Logger{
		Filename:   "./logs/sql.log",
		MaxSize:    2, // megabytes
		MaxBackups: 5,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	}
	return logger.New(
		log.New(l, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
}

func (m *Mysql) AutoMigrate() {
	m.DB.AutoMigrate(
		// &model.Address{},
		&model.User{},
		// &model.Follow{},
		// &model.Article{},
		// &model.Comment{},
		// &model.Tag{},
	)
}
func (m *Mysql) AutoMigrateRegions() {
	m.DB.AutoMigrate(
		&model.RegionDistrict{},
		&model.RegionStreet{},
		&model.RegionCity{},
		&model.RegionProvince{},
		&model.RegionCountry{},
	)
}

func (m *Mysql) AutoMigratePermission() {
	m.DB.AutoMigrate(
		&model.Role{},
		&model.Menu{},
	)
}
