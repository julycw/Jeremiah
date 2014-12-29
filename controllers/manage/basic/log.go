package basic

import (
	// . "github.com/julycw/BeibeiWeixin/models"
	"github.com/julycw/Jeremiah/controllers/manage"
)

type LogController struct {
	manage.BaseController
}

func (this *LogController) Get() {
	this.CheckAvaliable("查看日志")
	this.TplNames = "manage/basic/log.tpl"
}

func (this *LogController) Post() {

}
