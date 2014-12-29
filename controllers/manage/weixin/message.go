package weixin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
)

type MessageController struct {
	manage.BaseController
}

func (this *MessageController) Get() {
	this.CheckAvaliable("消息查看")
	this.TplNames = "manage/weixin/message.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/weixin/message_script.tpl"
}

func (this *MessageController) Post() {

}

// @router /api/manage/weixin/message/getMessageList [post]
func (this *MessageController) GetMessageList() {
	this.CheckAvaliable("消息查看")
	lastID, _ := this.GetInt("lastID")
	size, _ := this.GetInt("size")
	if size == 0 {
		size = 25
	}
	cond := orm.NewCondition()

	if lastID > 0 {
		cond = cond.And("Id__lt", lastID)
	}

	var messageList []models.Message
	if _, err := models.Orm.QueryTable("message").SetCond(cond).OrderBy("-Id").Limit(size).All(&messageList); err == nil {
		var messageExList []models.MessageEx = make([]models.MessageEx, len(messageList))
		for i, v := range messageList {
			var user models.FollowUser
			user.UserId = v.UserId
			models.Orm.Read(&user, "UserId")
			messageExList[i] = models.MessageEx{Message: v, User: user}
			messageExList[i].CreateOnFmt = v.CreateOn.Format("2006-01-02 15:04:05")
		}
		this.Data["json"] = &messageExList
	} else {
		beego.Error(err.Error())
	}

	this.ServeJson()
}
