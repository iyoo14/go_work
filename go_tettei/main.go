package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go_work/go_tettei/controller"
	"go_work/go_tettei/infra"
	"go_work/go_tettei/repository"
	"go_work/go_tettei/transaction"
	"go_work/go_tettei/usecase"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var dsn string
var exePath string

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type config struct {
	Dsn    string `json:"dsn"`
	Suffix string `json:suffix`
}

func main() {
	exe, _ := os.Executable()
	exePath = filepath.Dir(exe)
	setEnv()
	db := infra.Connect(dsn)
	fmt.Println(db)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	fmt.Println("Hello World")
	tr := transaction.NewTransaction(db)
	ur := repository.NewUserRepository(db)
	mr := repository.NewMessageRepository(db)
	uu := usecase.NewUserUsecase(ur, mr, tr)
	uc := controller.NewUserController(uu)
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users/:id", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.Start(":8080")

}

func setEnv() {
	fname := filepath.Join(exePath, "..", "config", "env.json")
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	dsn = cfg.Dsn
}
