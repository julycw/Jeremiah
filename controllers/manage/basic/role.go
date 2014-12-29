package basic

import (
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
)

type RoleController struct {
	manage.BaseController
}

func (this *RoleController) Get() {
	this.CheckAvaliable("角色管理")
	this.TplNames = "manage/basic/role.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/basic/role_script.tpl"
	var roleList []models.Role
	if _, err := models.Orm.QueryTable("Role").All(&roleList); err == nil {
		this.Data["roleList"] = &roleList
	}
}

func (this *RoleController) Post() {

}

func (this *RoleController) Add() {
	this.CheckAvaliable("角色管理")
	rolename := this.GetString("rolename")
	avaliables := this.GetString("avaliables")

	var lastRole []models.Role
	var roleId int
	if _, err := models.Orm.QueryTable("role").OrderBy("-RoleId").Limit(1).All(&lastRole); err == nil {
		roleId = lastRole[0].RoleId + 1
	} else {
		roleId = 1
	}
	new_role := models.Role{
		RoleId:     roleId,
		RoleName:   rolename,
		Avaliables: avaliables,
	}

	if _, err := models.Orm.Insert(&new_role); err == nil {

	}
	this.Redirect("/manage/basic/role", 302)
}

func (this *RoleController) Modify() {
	this.CheckAvaliable("角色管理")
	id, _ := this.GetInt("id")
	roleid, _ := this.GetInt("roleid")
	rolename := this.GetString("rolename")
	avaliables := this.GetString("avaliables")

	modify_role := models.Role{
		Id:         int(id),
		RoleId:     int(roleid),
		RoleName:   rolename,
		Avaliables: avaliables,
	}

	if _, err := models.Orm.Update(&modify_role); err == nil {

	}
	this.Redirect("/manage/basic/role", 302)
}

// @router /api/manage/basic/role/getRoleById [post]
func (this *RoleController) GetRoleById() {
	this.CheckAvaliable("角色管理")
	id, _ := this.GetInt("id")

	var role models.Role
	role.Id = int(id)
	if err := models.Orm.Read(&role); err == nil {
		this.Data["json"] = &role
	}

	this.ServeJson()
}

// @router /api/manage/basic/role/deleteRoleById [post]
func (this *RoleController) DeleteRoleById() {
	this.CheckAvaliable("角色管理")
	id, _ := this.GetInt("id")

	if _, err := models.Orm.Delete(&models.Role{Id: int(id)}); err == nil {
		this.Data["json"] = "success"
	} else {
		this.Data["json"] = "err:" + err.Error()
	}

	this.ServeJson()
}
