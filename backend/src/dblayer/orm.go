package dblayer

import (
	"errors"
	"go-rest-gin/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos(products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firtname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firtname, LastName: lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	if !checkPassword(pass) {
		return customer, errors.New("Invalid password")
	}
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignOutUserById(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").
		Joins("join customers on customers.id = customer_id").
		Joins("join products on products.id = product_id").
		Where("customer_id=?", id).
		Scan(&orders).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	// convert password string to byte slice so that we can sue it with the bcrypt package
	sBytes := []byte(*s)
	// Obtian hashed password via th GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}
