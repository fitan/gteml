package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/core"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

type Yaml2Go struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

var genconf = &cobra.Command{
	Use: "genconf",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile(*genConfSrc)
		if err != nil {
			panic(err.Error())
		}

		y2g := Yaml2Go{}
		_, err = resty.New().R().SetResult(&y2g).SetFormData(map[string]string{"schema": string(b)}).Post("https://www.printlove.cn/api/yaml2go")
		if err != nil {
			panic(err.Error())
		}

		data := strings.ReplaceAll(y2g.Data, "AutoGenerated", "MyConf")
		data = "package types\n\n" + data

		_, err = os.Stat(*genConfDest)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err.Error())
			}
		} else {
			err = os.Truncate(*genConfDest, 0)
			if err != nil {
				panic(err.Error())
			}
		}
		f, err := os.OpenFile(*genConfDest, os.O_CREATE|os.O_SYNC|os.O_RDWR, 0666)
		if err != nil {
			panic(err.Error())
		}
		_, err = f.Write([]byte(data))
		if err != nil {
			panic(err.Error())
		}
	},
}
var genConfSrc *string
var genConfDest *string

var migrate = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		//conf := core.NewConfReg()
		//client, err := ent.Open("mysql", conf.Confer.GetMyConf().Mysql.Url)
		//if err != nil {
		//	log.Fatalf("failed connecting to mysql: %v", err)
		//}
		//
		//if err := client.Schema.Create(context2.Background()); err != nil {
		//	log.Fatalf("failed creating schema resources: %v", err)
		//	return
		//}
	},
}

var gormMigrate = &cobra.Command{
	Use: "gorm-migrate",
	Run: func(cmd *cobra.Command, args []string) {
		conf := core.NewConfReg()
		client, err := gorm.Open(mysql.Open(conf.Confer.GetMyConf().Mysql.Url), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Fatalf("failed connecting to mysql: %v", err)
		}

		err = client.AutoMigrate(&model.User{}, &model.Role{}, &model.Service{}, &model.Permission{}, &model.Audit{}, &model.CasbinRule{})
		if err != nil {
			log.Fatalf("failed AutoMigrate %v", err)
		}
	},
}

var gormGenFake = &cobra.Command{
	Use: "gen-fake",
	Run: func(cmd *cobra.Command, args []string) {
		conf := core.NewConfReg()
		client, err := gorm.Open(mysql.Open(conf.Confer.GetMyConf().Mysql.Url), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Fatalf("failed connecting to mysql: %v", err)
		}
		for i := 0; i < 1000; i++ {
			err = client.Create(&model.User{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				PassWord: gofakeit.Password(true, true, true, true, true, 8),
				Token:    gofakeit.UUID(),
				Enable:   false,
			}).Error
			if err != nil {
				log.Fatalf("create err: %v", err.Error())
			}
		}
	},
}

func main() {
	genConfSrc = genconf.Flags().StringP("src", "s", "", "")
	genConfDest = genconf.Flags().StringP("dest", "d", "", "")
	var rootCmd = &cobra.Command{Use: "gen"}
	rootCmd.AddCommand(genconf)
	rootCmd.AddCommand(migrate)
	rootCmd.AddCommand(gormMigrate)
	rootCmd.AddCommand(gormGenFake)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("execute err: %v", err.Error())
	}
}
