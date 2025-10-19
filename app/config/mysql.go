package config

import (
	"database/sql"
	"fmt"

	"github.com/Uttamnath64/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	ReadWriteDB *gorm.DB
	ReadOnlyDB  *gorm.DB
	LogDB       *gorm.DB
}

func EnsureDB(env *AppEnv, log *logger.Logger) error {

	db, err := sql.Open("mysql", env.MySQL.DB.Main)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL without DB: %w", err)
	}
	defer db.Close()

	queries := []string{
		fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", env.MySQL.DB.Main),
		fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", env.MySQL.DB.Log),
	}

	for _, query := range queries {
		log.Info("app-db-check", "running", query)
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error executing '%s': %w", query, err)
		}
	}

	log.Info("app-db-check", "message", "✅ Required databases ensured.")
	return nil
}

func InitMysql(env *AppEnv) (*Mysql, error) {
	db := &Mysql{}
	var err error

	db.ReadWriteDB, err = connect(env.MySQL.DNS.ReadWrite)
	if err != nil {
		return nil, err
	}

	db.ReadOnlyDB, err = connect(env.MySQL.DNS.ReadOnly)
	if err != nil {
		return nil, err
	}

	db.LogDB, err = connect(env.MySQL.DNS.Log)
	if err != nil {
		return nil, err
	}

	db.ReadWriteDB = db.ReadWriteDB.Debug()
	db.ReadOnlyDB = db.ReadOnlyDB.Debug()
	db.LogDB = db.LogDB.Debug()

	return db, nil
}

// connect to DB
func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
