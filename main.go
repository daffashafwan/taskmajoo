package main
import (
	_mysqlDriver "github.com/daffashafwan/taskmajoo/config"
	_middleware "github.com/daffashafwan/taskmajoo/app/middlewares"
	
	"github.com/daffashafwan/taskmajoo/app/routes"

	//user
	_userRepository "github.com/daffashafwan/taskmajoo/model/user"
	_userUsecase "github.com/daffashafwan/taskmajoo/business/users"
	_userController "github.com/daffashafwan/taskmajoo/controllers/user"
	
	"log"
	"time"
	"github.com/labstack/echo/v4"
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

	routesInit := routes.ControllerList{
		JwtConfig:             configJWT.Init(),
		UserController:        *userController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
