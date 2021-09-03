package repository

import (
	"log"
	"time"

	"github.com/ahmadirfaan/loan-payment/models"
	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

// Customer : represent the customer's repository contract
type CustomerRepository interface {
	Save(models.Customer) (models.Customer, error)
	Migrate() error
	GetAllLimitToday() (isLimit bool)
}

// NewCustomerRepository -> returns new customer repository
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return customerRepository{
		DB: db,
	}
}

func (c customerRepository) Save(customer models.Customer) (models.Customer, error) {
	log.Print("[UserRepository]...Save")
	err := c.DB.Create(&customer).Error
	return customer, err

}

func (c customerRepository) Migrate() error {
	log.Print("[CustomerRepository]...Migrate")
	return c.DB.AutoMigrate(&models.Customer{})
}

func (c customerRepository) GetAllLimitToday() (isLimit bool) {
	log.Print("[UserRepository]...Get All Limit Today")
	var customer []models.Customer
	var count int64
	now := time.Now()
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Local().Location())
	nextDate := currentDate.AddDate(0, 0, 1)
	c.DB.Where("created_at between ? AND ?", currentDate, nextDate).Find(&customer).Count(&count)
	log.Print("Jumlah limit sekarang adalah: ", count)
	if count <= 50 {
		isLimit = false
	} else {
		isLimit = true
	}

	return isLimit
}



