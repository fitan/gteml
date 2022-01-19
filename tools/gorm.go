package main

import (
	"github.com/fitan/magic/dao/dal/model"
	"gorm.io/gen"
)

func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao/dal/query",
		Mode:    gen.WithoutContext,
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
	//db, _ := gorm.Open(mysql.Open("spider_dev:spider_dev123@tcp(10.170.34.22:3307)/gteml?parseTime=true"))
	//q := query.Use(db)
	//u := q.User
	//first, err := u.WithContext(context.Background()).Where(u.ID.Eq(1)).Preload(u.Roles).First()
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(first)
	//return

	//db = db.Debug()

	//g.UseDB(db)
	//
	//// apply basic crud api on structs or table models which is specified by table name with function
	//// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	//q := query.Use(db)
	//user := q.SysUser
	////_ = q.TblServicetree
	//first, err := user.WithContext(context.TODO()).Where(user.Id.Eq(1)).Preload(user.SysRoles).Preload(user.SysServicetree).First()
	//if err != nil {
	//	return
	//}
	//b, _ := json.Marshal(first)
	//fmt.Println(string(b))
	//return

	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.ApplyBasic(model.Service{}, model.Role{}, model.Permission{}, model.Audit{})
	g.ApplyInterface(func(method model.UserMethod) {}, model.User{})

	g.Execute()
}
