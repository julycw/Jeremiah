package basic

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
	"log"
)

type OptionController struct {
	manage.BaseController
}

func (this *OptionController) Get() {
	this.CheckAvaliable("选项设置")
	this.TplNames = "manage/basic/option.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/basic/option_script.tpl"

	cond := orm.NewCondition()
	if section := this.GetString(":section"); section != "" {
		cond = cond.And("section", section)
	}

	var optionSections []models.OptionSection = make([]models.OptionSection, 0)
	var optionList []models.Option
	if _, err := models.Orm.QueryTable("option").SetCond(cond).All(&optionList); err == nil {
		for _, option := range optionList {
			if !putOptionToSection(option, &optionSections) {
				var new_section models.OptionSection
				new_section.Options = make([]models.Option, 1)
				new_section.SectionName = option.SectionName
				new_section.Options[0] = option
				if len(optionSections) > 0 {
					optionSections[len(optionSections)-1].HasNext = true
				}
				optionSections = append(optionSections, new_section)
			}
		}
	} else {
		beego.Error(err.Error())
	}

	this.Data["OptionSections"] = &optionSections
}

func putOptionToSection(option models.Option, sections *[]models.OptionSection) bool {
	for s_index, section := range *sections {
		if section.SectionName == option.SectionName {
			(*sections)[s_index].Options = append((*sections)[s_index].Options, option)
			return true
		}
	}
	return false
}

func (this *OptionController) Post() {

}

// @router /api/manage/basic/option/getOptionById [post]
func (this *OptionController) GetOptionById() {
	this.CheckAvaliable("选项设置")
	id, _ := this.GetInt("id")

	var user models.Option
	user.Id = int(id)
	if err := models.Orm.Read(&user); err == nil {
		this.Data["json"] = &user
	} else {
		beego.Error(err.Error())
	}

	this.ServeJson()
}

// @router /api/manage/basic/option/modifyOptionById [post]
func (this *OptionController) ModifyOptionById() {
	this.CheckAvaliable("选项设置")
	id, _ := this.GetInt("id")
	value := this.GetString("value")

	user := models.Option{
		Id:    int(id),
		Value: value,
	}

	if _, err := models.Orm.Update(&user, "Value"); err == nil {
		this.Data["json"] = "success"
	} else {
		this.Data["json"] = "failed"
		log.Println(err.Error())
	}

	this.ServeJson()
}

func (this *OptionController) Modify() {
	this.CheckAvaliable("选项设置")
	id, _ := this.GetInt("id")
	value := this.GetString("value")

	user := models.Option{
		Id:    int(id),
		Value: value,
	}

	if _, err := models.Orm.Update(&user, "Value"); err == nil {

	}
	this.Redirect("/manage/weixin/option", 302)
}
