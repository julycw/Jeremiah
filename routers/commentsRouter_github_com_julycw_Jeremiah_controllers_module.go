package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/module:CodeScanController"] = append(beego.GlobalControllerRouter["github.com/julycw/Jeremiah/controllers/module:CodeScanController"],
		beego.ControllerComments{
			"ScanFixOrder",
			`/api/manage/weixin/fix/scanFixOrder`,
			[]string{"get"},
			nil})

}
