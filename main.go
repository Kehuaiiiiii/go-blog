package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	_ "go-blog/routers"
	"go-blog/service/databsae"
	"go-blog/utils"
)

func init() {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	database, _ := db.NewDataBase(conf.String("db::dbType"))
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())

	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexAddOne", utils.IndexAddOne)
	beego.AddFuncMap("IndexDecrOne", utils.IndexDecrOne)
	beego.AddFuncMap("StringReplace", utils.StringReplace)
	beego.AddFuncMap("TimeStampToTime", utils.TimeStampToTime)

}

func main() {
	fmt.Println("Server begin......")
	//bee generate appcode -tables="cron" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/blog" -level=3
	beego.Run()
}
