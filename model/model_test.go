package model

import (
	"contacts/utils"
	"fmt"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInit(t *testing.T) {
	type args struct {
		dbEngine string
		connStr  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "first test", args: args{dbEngine: "mysql", connStr: "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.dbEngine, tt.args.connStr); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestDeleteContact(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	contact := &Contact{AccountID: "uber", Email: "raushan@uber.com"}
	contact, err := AddContact(contact)
	if err != nil {
		t.Fail()
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testing delete func", args{id: contact.ID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteContact(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestAddContact(t *testing.T) {
	type args struct {
		contact *Contact
	}
	tests := []struct {
		name    string
		args    args
		want    *Contact
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Checking Add Contacts", args{contact: &Contact{AccountID: "uber", Email: "raushan@uber.com"}}, &Contact{AccountID: "uber", Email: "raushan@uber.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbEngine := "mysql"
			connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
			if err := Init(dbEngine, connStr); err != nil {
				fmt.Println(err)
				t.Fail()
			}
			got, err := AddContact(tt.args.contact)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountID, tt.want.AccountID) || !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("AddContact() = %v, want %v", got, tt.want)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestUpdateContact(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	contact := &Contact{AccountID: "uber", Email: "raushan@uber.com"}
	contact, err := AddContact(contact)
	if err != nil {
		t.Fail()
	}
	contact.Number = "8867997432"
	type args struct {
		contact *Contact
	}
	tests := []struct {
		name    string
		args    args
		want    *Contact
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testing Update Scenario", args{contact: contact}, contact, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateContact(tt.args.contact)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountID, tt.want.AccountID) || !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("AddContact() = %v, want %v", got, tt.want)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestGetContactByID(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	contact := &Contact{AccountID: "uber", Email: "raushan@uber.com"}
	contact, err := AddContact(contact)
	if err != nil {
		t.Fail()
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Contact
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Testing GetContactByID", args{id: contact.ID}, contact, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContactByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContactByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountID, tt.want.AccountID) || !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("AddContact() = %v, want %v", got, tt.want)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestGetAccountByID(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	account := &Account{AccountID: "uber", Password: utils.Hash("raushan@uber.com")}
	account, err := AddAccount(account)
	if err != nil {
		t.Fail()
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		// TODO: Add test cases.
		{"TestGetAccountByID ", args{id: "uber"}, account, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccountByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountID, tt.want.AccountID) || !reflect.DeepEqual(got.Password, tt.want.Password) {
				t.Errorf("AddAccount() = %v, want %v", got, tt.want)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}

func TestAddAccount(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	type args struct {
		account *Account
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Checking Add Account", args{account: &Account{AccountID: "uber", Password: utils.Hash("djcfdbvheczjcwdhcd")}}, &Account{AccountID: "uber", Password: utils.Hash("djcfdbvheczjcwdhcd")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddAccount(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountID, tt.want.AccountID) || !reflect.DeepEqual(got.Password, tt.want.Password) {
				t.Errorf("AddAccount() = %v, want %v", got, tt.want)
			}
		})
	}
	//db.DropTableIfExists(&Contact{})
	//db.DropTableIfExists(&Account{})
}

func TestNewContact(t *testing.T) {
	tests := []struct {
		name string
		want *Contact
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContact(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContact_TableName(t *testing.T) {
	type fields struct {
		Model       gorm.Model
		AccountID   string
		Number      string
		FirstName   string
		LastName    string
		CompanyName string
		URI         string
		Email       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				AccountID:   tt.fields.AccountID,
				Number:      tt.fields.Number,
				FirstName:   tt.fields.FirstName,
				LastName:    tt.fields.LastName,
				CompanyName: tt.fields.CompanyName,
				Email:       tt.fields.Email,
			}
			if got := c.TableName(); got != tt.want {
				t.Errorf("Contact.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDb(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDb(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name string
		want *Account
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_TableName(t *testing.T) {
	type fields struct {
		Model     gorm.Model
		AccountID string
		Password  uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Model:     tt.fields.Model,
				AccountID: tt.fields.AccountID,
				Password:  tt.fields.Password,
			}
			if got := a.TableName(); got != tt.want {
				t.Errorf("Account.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContactByName(t *testing.T) {
	dbEngine := "mysql"
	connStr := "root:heimdall@tcp(172.17.0.2:3306)/contact?charset=utf8&parseTime=True"
	if err := Init(dbEngine, connStr); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	contact1 := &Contact{AccountID: "uber", FirstName: "Raushan", Email: "raushan@uber.com"}
	contact1, err := AddContact(contact1)
	if err != nil {
		t.Fail()
	}
	contact2 := &Contact{AccountID: "uber", FirstName: "Raushan", Email: "raushan1@uber.com"}
	contact2, err = AddContact(contact2)
	if err != nil {
		t.Fail()
	}
	var contacts []Contact
	contacts = append(contacts, *contact1, *contact2)

	if err != nil {
		t.Fail()
	}
	type args struct {
		name       string
		accountSid string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]Contact
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Testing GetContactByID", args{name: contact1.FirstName, accountSid: contact1.AccountID}, &contacts, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContactByName(tt.args.name, tt.args.accountSid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContactByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(*got), len(*tt.want)) {
				t.Errorf("GetContactByName() = %v, want %v", got, tt.want)
			}
		})
	}
	db.DropTableIfExists(&Contact{})
	db.DropTableIfExists(&Account{})
}
