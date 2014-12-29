package weixin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
)

type FollowerController struct {
	manage.BaseController
}

func (this *FollowerController) Get() {
	this.CheckAvaliable("粉丝管理")
	this.TplNames = "manage/weixin/follower.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/weixin/follower_script.tpl"
}

func (this *FollowerController) Post() {
	this.CheckAvaliable("粉丝管理")

}

// @router /api/manage/weixin/follower/getFollowerList [post]
func (this *FollowerController) GetFollowerList() {
	this.CheckAvaliable("粉丝管理")
	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	size, _ := this.GetInt("size")
	if size == 0 {
		size = 100
	}

	cond := orm.NewCondition()

	var followerList []models.FollowUser
	if _, err := models.Orm.QueryTable("follow_user").SetCond(cond).OrderBy("-Id").Limit(size).All(&followerList); err == nil {
		this.Data["json"] = &followerList
	} else {
		beego.Error(err.Error())
	}

	this.ServeJson()
}

// @router /api/manage/weixin/follower/getFollowerById [post]
func (this *FollowerController) GetFollowerById() {
	this.CheckAvaliable("粉丝管理")
	id, _ := this.GetInt("id")

	var user models.FollowUser
	user.Id = int(id)
	if err := models.Orm.Read(&user); err == nil {
		this.Data["json"] = &user
	} else {
		beego.Error(err.Error())
	}

	this.ServeJson()
}

func (this *FollowerController) Modify() {
	this.CheckAvaliable("粉丝管理")
	id, _ := this.GetInt("id")
	userid := this.GetString("userid")
	userName := this.GetString("username")
	nickName := this.GetString("nickname")
	description := this.GetString("description")
	phone := this.GetString("phone")
	sex := this.GetString("sex")
	name := this.GetString("name")

	user := models.FollowUser{
		Id:          int(id),
		UserId:      userid,
		UserName:    userName,
		NickName:    nickName,
		Description: description,
		Sex:         sex,
		Phone:       phone,
		Name:        name,
	}

	if _, err := models.Orm.Update(&user); err == nil {

	}
	this.Redirect("/manage/weixin/follower", 302)
}
