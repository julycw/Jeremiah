package manage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/julycw/Jeremiah/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "manage/login.tpl"
	if redirect := this.GetString("redirect"); redirect != "" {
		this.Data["redirect"] = redirect
	}
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	redirect := this.GetString("redirect")
	md5Maker := md5.New()
	md5Maker.Write([]byte(password))
	md5Password := hex.EncodeToString(md5Maker.Sum(nil))

	var user models.User

	if err := models.Orm.QueryTable("user").Filter("user_name", username).Filter("pass_word", md5Password).One(&user); err != nil {
		this.Redirect(fmt.Sprintf("/manage/login?redirect=%s", redirect), 302)
	} else {
		sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("username", username)
		var role models.Role
		if err := models.Orm.QueryTable("role").Filter("role_id", user.RoleId).One(&role); err == nil {
			sess.Set("avaliables", role.Avaliables)
		}
		if redirect != "" {
			this.Redirect(redirect, 302)
		} else {
			this.Redirect("/manage", 302)
		}
	}
}
