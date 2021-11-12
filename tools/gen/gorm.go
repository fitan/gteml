package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/tools/gen/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
	db, _ := gorm.Open(mysql.Open("spider_dev:spider_dev123@tcp(10.170.34.22:3307)/spider_dev?charset=utf8mb4&parseTime=True&loc=Local"))
	db = db.Debug()
	g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	q := query.Use(db)
	user := q.SysUser
	_ = q.TblServicetree
	first, err := user.WithContext(context.TODO()).Preload(field.Associations).Where(user.Id.Eq(1)).First()
	if err != nil {
		return
	}
	b, _ := json.Marshal(first)
	fmt.Println(string(b))
	return

	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.ApplyBasic(model.SysUser{}, model.TblServicetree{}, model.SysRole{})

	g.Execute()
}
