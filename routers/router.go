package routers

import (
	"github.com/astaxie/beego"
	"github.com/julycw/Jeremiah/controllers"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/controllers/manage/basic"
	"github.com/julycw/Jeremiah/controllers/manage/note"
	"github.com/julycw/Jeremiah/controllers/manage/web"
	"github.com/julycw/Jeremiah/controllers/manage/weixin"
	"github.com/julycw/Jeremiah/controllers/module"
	weixin_router "github.com/julycw/Jeremiah/controllers/weixin"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/root.txt", &controllers.HomeController{}, "get:RootTxt")
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/gallery", &controllers.GalleryController{})
	beego.Router("/gallery/:id:int", &controllers.GalleryController{}, "get:Detail")
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/WhatWeDo", &controllers.NewsController{}, "get:WhatWeDo")
	beego.Router("/news/:id:int", &controllers.NewsController{}, "get:Detail")
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/services", &controllers.ServicesController{})
	beego.Router("/weixin", &weixin_router.WeixinController{})
	beego.Router("/resource/image/:gid:string/:size:string", &controllers.ResourceController{}, "get:GetThumbnailImage")

	beego.Router("/manage", &manage.IndexController{})
	beego.Router("/manage/home", &manage.IndexController{})
	beego.Router("/manage/note", &note.NoteController{})
	beego.Router("/manage/note/add", &note.NoteController{}, "post:Add")
	beego.Router("/manage/login", &manage.LoginController{})
	beego.Router("/manage/logout", &manage.LogoutController{})
	beego.Router("/manage/weixin/menu", &weixin.WeixinMenuController{})
	beego.Router("/manage/weixin/menu/get_menu_list", &weixin.WeixinMenuController{}, "get:GetMenuList;post:GetMenuList")
	beego.Router("/manage/weixin/order", &weixin.FixOrderController{})
	beego.Router("/manage/weixin/message", &weixin.MessageController{})
	beego.Router("/manage/weixin/follower", &weixin.FollowerController{})
	beego.Router("/manage/weixin/follower/modify", &weixin.FollowerController{}, "post:Modify")
	beego.Router("/manage/basic/option", &basic.OptionController{})
	beego.Router("/manage/basic/option/:section", &basic.OptionController{})
	beego.Router("/manage/basic/option/modify", &basic.OptionController{}, "post:Modify")
	beego.Router("/manage/basic/user", &basic.UserController{})
	beego.Router("/manage/basic/user/add", &basic.UserController{}, "post:Add")
	beego.Router("/manage/basic/user/modify", &basic.UserController{}, "post:Modify")
	beego.Router("/manage/basic/role", &basic.RoleController{})
	beego.Router("/manage/basic/role/add", &basic.RoleController{}, "post:Add")
	beego.Router("/manage/basic/role/modify", &basic.RoleController{}, "post:Modify")
	beego.Router("/manage/basic/log", &basic.LogController{})
	beego.Router("/manage/web/news", &web.NewsController{})
	beego.Router("/manage/web/news/add", &web.NewsController{}, "post:Add")
	beego.Router("/manage/web/news/modify", &web.NewsController{}, "post:Modify")
	beego.Router("/manage/web/news/delete", &web.NewsController{})
	beego.Router("/manage/web/product", &web.ProductController{})
	beego.Router("/manage/web/product/add", &web.ProductController{}, "post:Add")
	beego.Router("/manage/web/product/modify", &web.ProductController{}, "post:Modify")
	beego.Router("/manage/web/product/delete", &web.ProductController{})
	beego.Router("/manage/web/erp", &web.ERPController{})
	beego.Router("/manage/web/erp/scan", &web.ERPController{}, "get:Scan")

	beego.Include(&basic.OptionController{})
	beego.Include(&basic.UserController{})
	beego.Include(&basic.RoleController{})
	beego.Include(&note.NoteController{})
	beego.Include(&weixin.FixOrderController{})
	beego.Include(&weixin.MessageController{})
	beego.Include(&weixin.FollowerController{})
	beego.Include(&module.CodeScanController{})
	beego.Include(&weixin_router.WeixinController{})
	beego.Include(&weixin_router.WeixinPageController{})
	beego.Include(&web.NewsController{})
	beego.Include(&web.ProductController{})
	beego.Include(&web.ERPController{})
}
