package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// GetDb returns db objects for custom operations
func GetDb() *gorm.DB {
	return db
}

// Account Table for basic auth
type Account struct {
	gorm.Model
	AccountID string `gorm:"not null;unique;index"`
	Password  uint32
}

// NewAccount inits Account struct
func NewAccount() *Account {
	return &Account{}
}

//TableName - returns name of the table
//Implement mysql.GenericTable interface
func (*Account) TableName() string {
	return "accounts"
}

// Contact Table to store contact
type Contact struct {
	ID          uint       `form:"id" query:"id" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt   time.Time  `form:"date_created" query:"date_created" json:"date_created,omitempty"`
	UpdatedAt   time.Time  `form:"date_updated" query:"date_updated" json:"date_updated,omitempty"`
	DeletedAt   *time.Time `sql:"index"`
	AccountID   string     `form:"account_id" query:"account_id" json:"account_id,omitempty" gorm:"unique_index:account_id_email; not null"`
	Number      string     `form:"number" query:"number" json:"number,omitempty"`
	FirstName   string     `form:"first_name" query:"first_name" json:"first_name,omitempty"`
	LastName    string     `form:"last_name" query:"last_name" json:"last_name,omitempty"`
	CompanyName string     `form:"company_name" query:"company_name" json:"company_name,omitempty"`
	Email       string     `form:"email" query:"email" json:"email,omitempty" gorm:"unique_index:account_id_email; not null"`
}

// NewContact - inits Contact struct
func NewContact() *Contact {
	return &Contact{}
}

//TableName - returns name of the table
//Implement mysql.GenericTable interface
func (*Contact) TableName() string {
	return "contacts"
}

// Init initialize db connection
func Init(dbEngine, connStr string) error {
	var err error
	db, err = gorm.Open(dbEngine, connStr)
	if err != nil {
		return err
	}
	if db := db.AutoMigrate(&Account{}); db.Error != nil {
		return db.Error
	}
	if db := db.AutoMigrate(&Contact{}); db.Error != nil {
		return db.Error
	}
	return err
}

//AddContact add a new contact to db
func AddContact(contact *Contact) (*Contact, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return contact, err
}

//UpdateContact updates values of contact
func UpdateContact(contact *Contact) (*Contact, error) {
	var tempContact Contact
	tempContact.ID = contact.ID
	if db := db.Find(&tempContact); db.Error != nil {
		return nil, db.Error
	}
	if contact.CompanyName != "" {
		tempContact.CompanyName = contact.CompanyName
	}
	if contact.FirstName != "" {
		tempContact.FirstName = contact.FirstName
	}
	if contact.LastName != "" {
		tempContact.LastName = contact.LastName
	}
	if contact.Email != "" {
		tempContact.Email = contact.Email
	}
	if contact.Number != "" {
		tempContact.Number = contact.Number
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempContact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempContact, err
}

//GetContactByID return contact obeject based on id
func GetContactByID(id uint) (*Contact, error) {
	var contact Contact
	contact.ID = id
	if db := db.Find(&contact); db.Error != nil {
		return nil, db.Error
	}
	return &contact, nil
}

//GetContactByName return contact obeject based on Name
func GetContactByName(name, accountID string) (*[]Contact, error) {
	var contacts []Contact
	if db := db.Limit(10).Where("first_name = ? AND account_id = ?", name, accountID).Find(&contacts); db.Error != nil {
		return nil, db.Error
	}
	fmt.Println(contacts)
	fmt.Println(db)
	return &contacts, nil
}

// DeleteContact delete a contact completely
func DeleteContact(id uint) error {
	var contact Contact
	contact.ID = id
	if db := db.Find(&contact); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&contact).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}

//GetAccountByID return contact obeject based on id
func GetAccountByID(accountID string) (*Account, error) {
	var account Account
	account.AccountID = accountID
	if db := db.Find(&account); db.Error != nil {
		return nil, db.Error
	}
	return &account, nil
}

//AddAccount add a new contact to db
func AddAccount(account *Account) (*Account, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return account, err
}
