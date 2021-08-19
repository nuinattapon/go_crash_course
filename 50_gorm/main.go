package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var timezone *time.Location
var db *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time,
	fc func() (sql string, rowsAffected int64), err error) {

	sql, _ := fc()
	fmt.Printf("\n---\n%v\n--------\n", sql)

}

func main() {
	var err error
	// Initialize timezone variable
	timezone, err = time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err.Error())
	} else {
		log.Printf("Timezone is initialzied to '%v'\n", timezone)
	}

	dsn := "mysql:Welcome1@tcp(192.168.1.6)/gorm?charset=utf8mb4&parseTime=true&loc=Local"
	dial := mysql.Open(dsn)
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(Test{})
	// db.AutoMigrate(Gender{}, Test{}, Customer{})
	// CreateGender("XXX")
	// CreateGender("YYY")
	// GetGenders()
	// GetGender(1)
	// UpdateGender(2, "FEMALE")
	// DeleteGender(13)
	// GetGenders()
	// CreateCustomer("Wassana", 2)
	// CreateCustomer("Nattapon", 1)
	GetCustomers()
}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload("Gender").Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Println(customer.toString())
	}
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}
func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func UpdateGender(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)

	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}

	for _, gender := range genders {
		fmt.Println(gender.toString())
	}
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender.toString())
}

func GetGenderByName(name string) {
	gender := Gender{}
	tx := db.Where("name=?", name).First(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender.toString())
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender.toString())
}

type Test struct {
	gorm.Model
	Code uint
	Name string `gorm:"size:50;unique;not null"  json:"name"`
}
type Gender struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}

type Customer struct {
	gorm.Model
	Name     string
	Gender   Gender
	GenderID uint
}

func (gender Gender) toString() string {
	return fmt.Sprintf("ID %2d - Gender '%v'", gender.ID, gender.Name)
}

func (customer Customer) toString() string {
	return fmt.Sprintf("ID %2d - Name '%v' - Gender '%v'", customer.ID, customer.Name, customer.Gender.Name)
}
