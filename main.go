package main

import (
	"github.com/astaxie/beego"
	"github.com/julycw/Jeremiah/models"
	_ "github.com/julycw/Jeremiah/routers"
	"reflect"
	"strings"
)

func productType(in int) (out string) {
	var producttype models.ProductType
	p := reflect.ValueOf(&producttype)
	p.Elem().SetInt(int64(in))
	out = producttype.String()
	return
}
func newsCategory(in int) (out string) {
	var category models.NewsCategory
	p := reflect.ValueOf(&category)
	p.Elem().SetInt(int64(in))
	out = category.String()
	return
}

func newsState(in int) (out string) {
	var state models.NewsState
	p := reflect.ValueOf(&state)
	p.Elem().SetInt(int64(in))
	out = state.String()
	return
}

func roleName(in int) (out string) {
	role := models.Role{RoleId: in}
	if err := models.Orm.QueryTable("Role").Filter("role_id", in).One(&role); err == nil {
		return role.RoleName
	} else {
		println(err.Error())
	}

	return "未知"
}

func nickName(in string) (out string) {
	var user models.User
	if err := models.Orm.QueryTable("User").Filter("user_name", in).One(&user); err == nil {
		return user.NickName
	}

	return in
}
func getFileName(in string) (out string) {
	index := strings.LastIndex(in, "/")
	index2 := strings.Index(in, ".jpg")
	return in[index+1 : index2]
}

func main() {
	beego.AddFuncMap("newsCategory", newsCategory)
	beego.AddFuncMap("newsState", newsState)
	beego.AddFuncMap("productType", productType)
	beego.AddFuncMap("roleName", roleName)
	beego.AddFuncMap("nickName", nickName)
	beego.AddFuncMap("getFileName", getFileName)
	beego.EnableAdmin = true
	beego.AdminHttpPort = 8088
	beego.Run()
}
