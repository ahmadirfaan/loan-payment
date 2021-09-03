package service

import (
	"errors"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/ahmadirfaan/loan-payment/models"
	"github.com/ahmadirfaan/loan-payment/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CustomerService interface {
	Save(models.Customer) (models.Customer, error)
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

// NewCustomerSerivce -> returns new Customer service
func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return customerService{
		customerRepository: r,
	}
}

func (c customerService) Save(customer models.Customer) (models.Customer, error) {
	v := validator.New()
	err := v.Struct(&customer)
	if err != nil {
		return customer, err
	}
	if !verifyProvinceKTP(&customer) {
		return customer, errors.New("your province is not at the list")
	} 
	if !verifyTenor(&customer) {
		return customer, errors.New("your not choose the appropriate tenor")
	} 
	if c.getAllLimitToday() {
		return customer, errors.New("sorry the limit is exceed")
	} 
	if !verifyAge(&customer) {
		return customer, errors.New("sorry your age is not allowed to doing this transaction")
	} 
	id := uuid.New()
	customer.ID = id.String()
	customer.AddressProvince = generatedAddressProvinceByKTP(&customer)
	customer.Status = generateStatus(&customer)
	return c.customerRepository.Save(customer)

}

func (c customerService) getAllLimitToday() (isLimit bool) {
	return c.customerRepository.GetAllLimitToday()
}

//function verifyProvinceKTP to verification number in KTP
func verifyProvinceKTP(customer *models.Customer) bool {
	sliceProv, _ := strconv.Atoi(customer.Ktp[0:2])

	var provPermitted = map[int]string{
		12: "Sumatera Utara",
		31: "DKI Jakarta",
		32: "Jawa Barat",
		35: "Jawa Timur",
	}

	_, isExist := provPermitted[sliceProv]

	return isExist
}

//function verifyProvinceKTP to verification number in KTP
func generatedAddressProvinceByKTP(customer *models.Customer) string {
	sliceProv, _ := strconv.Atoi(customer.Ktp[0:2])

	switch sliceProv {
	case 12:
		return "Sumatera Utara"
	case 31:
		return "DKI Jakarta"
	case 32:
		return "Jawa Barat"
	case 35:
		return "Jawa Timur"
	default:
		return ""
	}
}

//verify Tenor the Input
func verifyTenor(customer *models.Customer) bool {
	listTenor := [5]int{3, 6, 9, 12, 24}
	for _, tenor := range listTenor {
		if tenor == customer.Tenor {
			return true
		}
	}
	return false
}

//Verify Age based on the birth date and date now
func verifyAge(customer *models.Customer) bool {
	log.Print("Verify Age")
	nowTime := time.Now()
	birthCustomer := &customer.BirthDate
	log.Print("Customer Birth is:", birthCustomer)
	nowDate := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	birtDate := time.Date(birthCustomer.Year(), birthCustomer.Month(), birthCustomer.Day(), 0, 0, 0, 0, time.Local)
	age := math.Floor(float64(nowDate.Sub(birtDate).Hours() / 24 / 365))
	log.Print("Age is:", age)
	if age > 17 && age < 80 {
		return true
	} else {
		return false
	}
}

//Generated Status is Accepted or Rejected
func generateStatus(customer *models.Customer) string {
	if customer.AddressProvince == "Jawa Timur" && customer.Amount < 8000000 {
		return "rejected"
	}
	if customer.AddressProvince == "Jawa Barat" && customer.Amount < 7000000 {
		return "rejected"
	}
	if customer.AddressProvince == "Sumatera Utara" && customer.Amount > 6000000 {
		return "rejected"
	}
	return "accepted"
}
