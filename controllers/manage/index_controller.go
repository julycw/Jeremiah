package manage

import (
// . "github.com/julycw/Jeremiah/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplNames = "manage/index.tpl"
}

func (this *IndexController) Post() {

}
