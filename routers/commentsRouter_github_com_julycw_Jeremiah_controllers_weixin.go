package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"],
		beego.ControllerComments{
			"UserInfo",
			`/weixin/user`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"],
		beego.ControllerComments{
			"UserInfoSave",
			`/weixin/user/save`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"],
		beego.ControllerComments{
			"ViewOrder",
			`/weixin/viewOrder`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"],
		beego.ControllerComments{
			"BindOrderNumber",
			`/weixin/bindOrderNumber`,
			[]string{"get","post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/weixin:WeixinPageController"],
		beego.ControllerComments{
			"ShowError",
			`/weixin/error`,
			[]string{"get"},
			nil})

}
