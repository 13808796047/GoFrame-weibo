package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20210902_175536 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210902_175536{}
	m.Created = "20210902_175536"

	migration.Register("User_20210902_175536", m)
}

// Run the migrations
func (m *User_20210902_175536) Up() {
	m.CreateTable("users", "innodb", "utf8mb4")
	m.PriCol("id").SetAuto(true).SetDataType("int").SetUnsigned(true)
	m.NewCol("name").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("email").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("password").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("remember_token").SetDataType("varchar(255)").SetNullable(true)
	m.NewCol("email_verified_at").SetDataType("varchar(255)").SetNullable(true)
	m.NewCol("created_at").SetDataType("timestamp").SetDefault("NOW()")
	m.NewCol("updated_at").SetDataType("timestamp").SetNullable(false).SetDefault("NOW()")
	m.SQL(m.GetSQL())


}

// Reverse the migrations
func (m *User_20210902_175536) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
