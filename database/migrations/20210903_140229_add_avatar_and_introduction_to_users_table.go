package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddAvatarAndIntroductionToUsersTable_20210903_140229 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddAvatarAndIntroductionToUsersTable_20210903_140229{}
	m.Created = "20210903_140229"

	migration.Register("AddAvatarAndIntroductionToUsersTable_20210903_140229", m)
}

// Run the migrations
func (m *AddAvatarAndIntroductionToUsersTable_20210903_140229) Up() {
	m.AlterTable("users")
	m.NewCol("avatar").SetDataType("varchar(255)").SetNullable(true)
	m.NewCol("introduction").SetDataType("text").SetNullable(true)
	m.SQL(m.GetSQL())
}

// Reverse the migrations
func (m *AddAvatarAndIntroductionToUsersTable_20210903_140229) Down() {

}
