package persist

import (
	"errors"
	"fmt"
	"log"

	"github.com/el-Mike/gochat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormWrapper struct {
	db *gorm.DB
}

// GormBroker - database broker based on Gorm.
var GormBroker *gormWrapper

// First - wrapper for Gorm's First method.
func (gm *gormWrapper) First(dest interface{}, conds ...interface{}) *DBResponse {
	res := gm.db.First(dest, conds...)

	return dbResponseFromGormResult(res)
}

// Where - wrapper for Gorm's Where method.
func (gm *gormWrapper) Where(query interface{}, args ...interface{}) *DBResponse {
	res := gm.db.Where(query, args...)

	return dbResponseFromGormResult(res)
}

// FirstWhere - returns first record that matches given criteria.
func (gm *gormWrapper) FirstWhere(dest interface{}, query interface{}, args ...interface{}) *DBResponse {
	res := gm.db.Where(query, args...).First(dest)

	return dbResponseFromGormResult(res)
}

// Find - wrapper for Gorm's Find method.
func (gm *gormWrapper) Find(dest interface{}, conds ...interface{}) *DBResponse {
	res := gm.db.Find(dest, conds...)

	return dbResponseFromGormResult(res)
}

// Save - wrapper for Gorm's Save method.
func (gm *gormWrapper) Save(value interface{}) *DBResponse {
	res := gm.db.Save(value)

	return dbResponseFromGormResult(res)
}

func dbResponseFromGormResult(result *gorm.DB) *DBResponse {
	res := NewDBResponse()

	if result.Error != nil {
		res.SetErr(result.Error)
	}

	return res
}

// InitDatabase - initializes database driver.
func InitDatabase(username, password, dbname, host, port string) (*gormWrapper, error) {
	if GormBroker != nil {
		return GormBroker, nil
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	gormWrapper := &gormWrapper{
		db: conn,
	}

	GormBroker = gormWrapper

	log.Println("Database connection estabilished")

	return GormBroker, nil
}

// AutoMigrate - fires the migrations for DB schemas.
func AutoMigrate() error {
	if GormBroker == nil {
		return errors.New("Database has not been initialized")
	}

	err := GormBroker.db.AutoMigrate(
		&models.UserModel{},
	)

	if err != nil {
		return err
	}

	return nil
}
