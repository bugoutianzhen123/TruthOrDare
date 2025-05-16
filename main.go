package main

import (
	"github.com/bugoutianzhen123/TruthOrDare/handler"
	"github.com/bugoutianzhen123/TruthOrDare/ioc"
	"github.com/bugoutianzhen123/TruthOrDare/repository"
	"github.com/bugoutianzhen123/TruthOrDare/repository/dao"
	"github.com/bugoutianzhen123/TruthOrDare/router"
	"github.com/bugoutianzhen123/TruthOrDare/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViper()

	l := ioc.InitLogger()
	db := ioc.InitDB(l)
	d := dao.NewDao(db, l)
	repo := repository.NewRepository(d)
	ser := service.NewService(repo)
	h := handler.NewHandler(ser, l)
	e := router.InitEngine(h)
	e.Run(":8080")
}

func initViper() {
	cfile := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
