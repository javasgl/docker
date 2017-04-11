package controllers

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	db, err := sql.Open("mysql", "root:123456@tcp(mysql:3306)/mysql")
	if err != nil {
		fmt.Println("sql open error:" + err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("sql ping error:" + err.Error())
	} else {
		fmt.Println("conn to mysql success!")
		rows, err := db.Query("select User,Host from user")
		if err != nil {
			fmt.Println("sql query error:" + err.Error())
		} else {
			fmt.Println(rows)
			cols, _ := rows.Columns()
			for i := range cols {
				fmt.Println("cols:" + cols[i])
			}
			for rows.Next() {
				var User, Host string
				rows.Scan(&User, &Host)
				fmt.Println("user:" + User + ",Host:" + Host)
			}
		}
	}

	c.Data["Website"] = "songgl.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
