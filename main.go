package main

import (
	"fmt"
	"gohexa/handler"
	"gohexa/logs"
	"gohexa/repositorie"
	"gohexa/service"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	// y := repositorie.NewAccountRepositoryDB()

	initConfig()

	db := initDB()

	customerRepositoryDB := repositorie.NewCustomerRepositoryDB(db)

	// customerRepositoryMock := repositorie.NewCustomerRepositoryMock()

	// _ = customerRepositoryDB
	// _ = customerRepositoryMock

	// _ = customerRepository
	customerService := service.NewCustomerService(customerRepositoryDB)

	// _ = customerService

	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepositoryDB := repositorie.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customer/{id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/custormers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)

	router.HandleFunc("/custormers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	// fmt.Printf("Test " + viper.GetString("app.port"))
	logs.Info("Banking service started at port " + viper.GetString("app.port"))

	http.ListenAndServe(":8000", router)

	// customers, err := customerService.GetCustomer(1)
	// // customers, err := customerService.GetCustomers()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)
	// _ = customerRepository

	// customer, err := customerRepository.GetById(1)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customer)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.table"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)

	return db
}
