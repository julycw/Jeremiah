package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:MessageController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:MessageController"],
		beego.ControllerComments{
			"GetMessageList",
			`/api/manage/weixin/message/getMessageList`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"],
		beego.ControllerComments{
			"GetOrderNumber",
			`/api/manage/weixin/fix/getOrderNumber`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"],
		beego.ControllerComments{
			"GetFixOrderDetailsByOrderId",
			`/api/manage/weixin/fix/getFixOrderDetailsByOrderId`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"],
		beego.ControllerComments{
			"AddFixOrder",
			`/api/manage/weixin/fix/addFixOrder`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"],
		beego.ControllerComments{
			"AddFixOrderDetail",
			`/api/manage/weixin/fix/addFixOrderDetail`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FixOrderController"],
		beego.ControllerComments{
			"GetOrderBindedUser",
			`/api/manage/weixin/fix/getOrderBindedUser`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FollowerController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FollowerController"],
		beego.ControllerComments{
			"GetFollowerList",
			`/api/manage/weixin/follower/getFollowerList`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FollowerController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/weixin:FollowerController"],
		beego.ControllerComments{
			"GetFollowerById",
			`/api/manage/weixin/follower/getFollowerById`,
			[]string{"post"},
			nil})

}
