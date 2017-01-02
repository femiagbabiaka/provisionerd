package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"database/sql"
	"os"
	"fmt"
	"errors"
)

type virtualMailer struct {
	AutomationMailerID int `json:"automation_mailer_id" sql:"automationmailer_id"`
	Name sql.NullString `json:"name" sql:"name"`
	SMTPHost sql.NullString `json:"smtp_host" sql:"smtphost"`
	BounceFormat sql.NullString `json:"-" sql:"bounceformat"`
	IPAddress sql.NullString `json:"ip_address" sql:"ip"`
	Category int `json:"category" sql:"category"`
}

func (vm virtualMailer) CreateMailer() (virtualMailer, error) {
	if ! vm.ValidateMailer() {
		return vm, ErrInvalidMailer
	}
	
	db, err := databaseConnect()

	if err != nil {
		return vm, ErrInvalidMySQLServer
	}
	
	_, err = db.NamedExec(`INSERT INTO mailserver (automationmailer_id, name, smtphost, smtpport, bounceformat, vmtaname, ip, category) VALUES (:automationmailer_id, :name, :smtphost, :smtpport, :bounceformat, :vmtaname, :ip, :category)`, vm)
	
	if err != nil {
		return vm, err
	}

	return vm, nil
}

// DeleteMailer Is responsible for deleting a mailserver.
func DeleteMailer(id int) (bool, error) {
	if id <= 0 {
		return false, ErrInvalidMailer
	}
	
	db, err := databaseConnect()

	if err != nil {
		return false, ErrInvalidMySQLServer
	}
	
	_, err = db.NamedExec(`DELETE FROM mailserver WHERE id = :id`,
		map[string]interface{}{
			"id": id,
		})
	
	if err != nil {
		return false, err
	}
	
	return true, nil
	
}

func (vm virtualMailer) ValidateMailer() bool {
	if vm.Name.String == "" {
		return false
	}
	
	if vm.SMTPHost.String == "" {
		return false
	}
	
	if vm.BounceFormat.String == "" {
		return false
	} 
	
	if vm.IPAddress.String == "" {
		return false
	}
	
	if vm.Category == 0 {
		return false
	}
	
	return true
}

func databaseConnect() (*sqlx.DB, error) {
	mysqlHost := os.Getenv("PROVISIONERD_VIRTUALMAILER_MYSQL_HOST")
	mysqlPort := os.Getenv("PROVISIONERD_VIRTUALMAILER_MYSQL_PORT")
	mysqlUsername := os.Getenv("PROVISIONERD_VIRTUALMAILER_MYSQL_USERNAME")
	mysqlPassword := os.Getenv("PROVISIONERD_VIRTUALMAILER_MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("PROVISIONERD_VIRTUALMAILER_MYSQL_DATABASE")
	
	db, err := sqlx.Connect(
		"mysql", 
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
			mysqlUsername,
			mysqlPassword,
			mysqlHost,
			mysqlPort,
			mysqlDatabase))
	if err != nil {
		return db, ErrInvalidMySQLServer
	}
	
	return db, nil

}

// ErrInvalidMailer Error representing an invalid mailer.
var ErrInvalidMailer = errors.New("Invalid VMTA provided.")
// ErrInvalidMySQLServer Error representing an inability to connect to MySQL.
var ErrInvalidMySQLServer = errors.New("Could not connect to database.")
