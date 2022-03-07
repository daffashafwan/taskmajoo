package main

import (
	_middleware "github.com/daffashafwan/taskmajoo/app/middlewares"
	_mysqlDriver "github.com/daffashafwan/taskmajoo/config"

	"github.com/daffashafwan/taskmajoo/app/routes"

	//user
	_userUsecase "github.com/daffashafwan/taskmajoo/business/users"
	_userController "github.com/daffashafwan/taskmajoo/controllers/user"
	_userRepository "github.com/daffashafwan/taskmajoo/model/user"

	//merchant
	_merchantUsecase "github.com/daffashafwan/taskmajoo/business/merchants"
	_merchantController "github.com/daffashafwan/taskmajoo/controllers/merchant"
	_merchantRepository "github.com/daffashafwan/taskmajoo/model/merchant"

	//outlet
	//_outletUsecase "github.com/daffashafwan/taskmajoo/business/outlets"
	_outletRepository "github.com/daffashafwan/taskmajoo/model/outlet"

	//transaction
	_transactionUsecase "github.com/daffashafwan/taskmajoo/business/transactions"
	_transactionController "github.com/daffashafwan/taskmajoo/controllers/transaction"
	_transactionRepository "github.com/daffashafwan/taskmajoo/model/transaction"

	"github.com/labstack/echo/v4"
	"log"
	"time"
	//"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()

	e := echo.New()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.CreateUserRepo(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	merchantRepository := _merchantRepository.CreateMerchantRepo(Conn)
	merchantUsecase := _merchantUsecase.NewMerchantUsecase(merchantRepository, timeoutContext, configJWT)
	merchantController := _merchantController.NewMerchantController(merchantUsecase)

	outletRepository := _outletRepository.CreateOutletRepo(Conn)
	//outletUsecase := _outletUsecase.NewOutletUsecase(outletRepository, timeoutContext, configJWT)

	transactionRepository := _transactionRepository.CreateTransactionRepo(Conn)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(transactionRepository,merchantRepository,outletRepository, timeoutContext, configJWT)
	transactionController := _transactionController.NewTransactionController(transactionUsecase, merchantUsecase)

	routesInit := routes.ControllerList{
		JwtConfig:          configJWT.Init(),
		UserController:     *userController,
		MerchantController: *merchantController,
		TransactionController: *transactionController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
