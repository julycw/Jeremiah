package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:NewsController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:NewsController"],
		beego.ControllerComments{
			"GetNewsById",
			`/api/manage/web/news/getNewsById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:NewsController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:NewsController"],
		beego.ControllerComments{
			"DeleteNewsById",
			`/api/manage/web/news/deleteNewsById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:ProductController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:ProductController"],
		beego.ControllerComments{
			"GetProductById",
			`/api/manage/web/product/getProductById`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:ProductController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/manage/web:ProductController"],
		beego.ControllerComments{
			"DeleteProductById",
			`/api/manage/web/product/deleteProductById`,
			[]string{"post"},
			nil})

}
