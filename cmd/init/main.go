package main

import (
	"database/sql"
	"os"

	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/multierr"
)

type DataStore struct {
	Host     string
	User     string
	Password string
	DBName   string
}

var postgres = DataStore{
	os.Getenv("DB_HOST_NAME"),
	os.Getenv("DB_USER_NAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
}

var identity = DataStore{
	os.Getenv("DB_HOST_NAME"),
	os.Getenv("IDENTITY_DB_USER_NAME"),
	os.Getenv("IDENTITY_DB_PASSWORD"),
	os.Getenv("IDENTITY_DB_USER_NAME") + "db",
}

var dsList = []*DataStore{&identity}

func main() {
	rerr := run()
	if rerr != nil {
		for _, err := range multierr.Errors(rerr) {
			fmt.Println(err)
		}
	}
}

func run() error {
	driverName := "postgres"
	db, err := connectDB(driverName, postgres.Host, postgres.User, postgres.Password, postgres.DBName)
	if err != nil {
		return err
	}
	var rerr error
	for _, ds := range dsList {
		if err := createDB(db, ds); err != nil {
			rerr = multierr.Append(rerr, err)
		}
	}
	db.Close()

	for _, ds := range dsList {
		if err := initDB(driverName, ds); err != nil {
			rerr = multierr.Append(rerr, err)
		}
	}

	return rerr

}

func initDB(driverName string, ds *DataStore) error {
	db, err := connectDB(driverName, ds.Host, postgres.User, postgres.Password, ds.DBName)
	if err != nil {
		return err
	}
	var rerr error
	if err := createUser(db, ds); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := grantDB(db, ds); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := revokeDB(db, ds); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := createSchema(db, ds); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	if err := grantUser(db, ds); err != nil {
		rerr = multierr.Append(rerr, err)
	}
	defer db.Close()
	return rerr
}

func connectDB(driverName string, host, user, password, dbname string) (*sql.DB, error) {
	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open(driverName, connectStr)
	if err != nil {
		return nil, fmt.Errorf("DB Open: %w", err)
	}
	fmt.Printf("Success for Connect Database %s\n", dbname)
	return db, nil
}

func createDB(db *sql.DB, ds *DataStore) error {
	sqlStr := "CREATE DATABASE " + ds.DBName + ";"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("create db: %w", err)
	}
	fmt.Printf("Success for Create Database %s\n", ds.DBName)
	return nil
}

func createUser(db *sql.DB, ds *DataStore) error {
	sqlStr := "CREATE ROLE " + ds.User + " WITH LOGIN PASSWORD " + "'" + ds.Password + "'" + ";"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	fmt.Printf("Success for Create User %s\n", ds.User)
	return nil
}

func grantDB(db *sql.DB, ds *DataStore) error {
	sqlStr := "GRANT CONNECT ON DATABASE " + ds.DBName + " TO " + ds.User + ";"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("grant db: %w", err)
	}
	fmt.Printf("Success for grant Database %s\n", ds.DBName)
	return nil
}

func revokeDB(db *sql.DB, ds *DataStore) error {
	sqlStr := "REVOKE CONNECT ON DATABASE " + ds.DBName + " FROM PUBLIC;"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("revoke db: %w", err)
	}
	fmt.Printf("Success for revoke Database %s\n", ds.DBName)
	return nil
}

func createSchema(db *sql.DB, ds *DataStore) error {
	sqlStr := "CREATE SCHEMA " + ds.User + " AUTHORIZATION " + ds.User + ";"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	fmt.Printf("Success for Create Schema %s\n", ds.User)
	return nil
}
func grantUser(db *sql.DB, ds *DataStore) error {
	sqlStr := "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA " + ds.User + " TO " + ds.User + ";"
	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("grant user: %w", err)
	}
	fmt.Printf("Success for Grant User %s\n", ds.User)
	return nil
}
