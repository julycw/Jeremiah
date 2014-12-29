package weixin

import (
	// . "github.com/julycw/BeibeiWeixin/models"
	"github.com/julycw/Jeremiah/controllers/manage"
	"strconv"
)

type WeixinMenuController struct {
	manage.BaseController
}

type MenuItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (this *WeixinMenuController) Get() {
	this.TplNames = "manage/weixin/menu.tpl"

	this.SetPageHeader("菜单管理")
	this.SetPageHeaderSmall("管理微信菜单项目")
	this.LayoutSections["LayoutFooter"] = "manage/weixin/menu_script.tpl"
}

func (this *WeixinMenuController) Post() {

}

func (this *WeixinMenuController) GetMenuList() {
	var menuItems []MenuItem = make([]MenuItem, 0)
	var menuList map[string]MenuItem = make(map[string]MenuItem, 0)
	menuItems = append(menuItems, MenuItem{Name: "test", Type: "item"})
	menuItems = append(menuItems, MenuItem{Name: "test2", Type: "item"})
	menuItems = append(menuItems, MenuItem{Name: "test3", Type: "item"})
	for i, v := range menuItems {
		menuList[strconv.Itoa(i)] = v
	}
	this.Data["json"] = &menuList
	this.ServeJson()
}
