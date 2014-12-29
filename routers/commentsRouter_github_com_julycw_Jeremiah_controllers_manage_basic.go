package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:OptionController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:OptionController"],
		beego.ControllerComments{
			"GetOptionById",
			`/api/manage/basic/option/getOptionById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:OptionController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:OptionController"],
		beego.ControllerComments{
			"ModifyOptionById",
			`/api/manage/basic/option/modifyOptionById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:RoleController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:RoleController"],
		beego.ControllerComments{
			"GetRoleById",
			`/api/manage/basic/role/getRoleById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:RoleController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:RoleController"],
		beego.ControllerComments{
			"DeleteRoleById",
			`/api/manage/basic/role/deleteRoleById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:UserController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:UserController"],
		beego.ControllerComments{
			"GetUserById",
			`/api/manage/basic/user/getUserById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:UserController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/basic:UserController"],
		beego.ControllerComments{
			"DeleteUserById",
			`/api/manage/basic/user/deleteUserById`,
			[]string{"post"},
			nil})

}
