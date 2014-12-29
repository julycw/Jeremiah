package basic

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
)

type UserController struct {
	manage.BaseController
}

func (this *UserController) Get() {
	this.CheckAvaliable("用户管理")
	this.TplNames = "manage/basic/user.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/basic/user_script.tpl"
	var userList []models.User
	if _, err := models.Orm.QueryTable("User").All(&userList); err == nil {
		this.Data["userList"] = &userList
	}
	var roleList []models.Role
	if _, err := models.Orm.QueryTable("Role").All(&roleList); err == nil {
		this.Data["roleList"] = &roleList
	}
}

func (this *UserController) Post() {

}

func (this *UserController) Add() {
	this.CheckAvaliable("用户管理")
	username := this.GetString("username")
	password := this.GetString("password")
	nickname := this.GetString("nickname")
	roleid, _ := this.GetInt("roleid")

	md5Maker := md5.New()
	md5Maker.Write([]byte(password))
	md5Password := hex.EncodeToString(md5Maker.Sum(nil))

	new_user := models.User{
		UserName: username,
		PassWord: md5Password,
		NickName: nickname,
		RoleId:   int(roleid),
	}

	if _, err := models.Orm.Insert(&new_user); err == nil {

	}
	this.Redirect("/manage/basic/user", 302)
}

func (this *UserController) Modify() {
	this.CheckAvaliable("用户管理")
	var updateField []string = []string{"NickName", "UserName", "RoleId"}
	id, _ := this.GetInt("id")
	username := this.GetString("username")
	password := this.GetString("password")
	nickname := this.GetString("nickname")
	roleid, _ := this.GetInt("roleid")

	if password != "" {
		updateField = append(updateField, "PassWord")
		md5Maker := md5.New()
		md5Maker.Write([]byte(password))
		password = hex.EncodeToString(md5Maker.Sum(nil))
	}

	modify_user := models.User{
		Id:       int(id),
		UserName: username,
		PassWord: password,
		NickName: nickname,
		RoleId:   int(roleid),
	}

	if _, err := models.Orm.Update(&modify_user, updateField...); err == nil {

	}
	this.Redirect("/manage/basic/user", 302)
}

// @router /api/manage/basic/user/getUserById [post]
func (this *UserController) GetUserById() {
	this.CheckAvaliable("用户管理")
	id, _ := this.GetInt("id")

	var user models.User
	user.Id = int(id)
	if err := models.Orm.Read(&user); err == nil {
		user.PassWord = "就不给看:P"
		this.Data["json"] = &user
	}

	this.ServeJson()
}

// @router /api/manage/basic/user/deleteUserById [post]
func (this *UserController) DeleteUserById() {
	this.CheckAvaliable("用户管理")
	id, _ := this.GetInt("id")

	if _, err := models.Orm.Delete(&models.User{Id: int(id)}); err == nil {
		this.Data["json"] = "success"
	} else {
		this.Data["json"] = "err:" + err.Error()
	}

	this.ServeJson()
}
