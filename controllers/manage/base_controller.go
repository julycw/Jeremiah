package manage

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/julycw/Jeremiah/models"
	"log"
	"strconv"
	"strings"
)

var globalSessions *session.Manager

type BaseController struct {
	beego.Controller
	pageHeader      string
	pageHeaderSmall string
	User            models.User
}

func (this *BaseController) Prepare() {
	this.CheckAvaliable("登陆")
	this.Data["SiteIp"] = beego.AppConfig.String("siteip")
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	if username := sess.Get("username"); username != nil {
		if username.(string) != "" {
			this.User.UserName = username.(string)
			this.Data["user"] = &this.User
			this.Layout = "manage/layout.tpl"
			this.LayoutSections = make(map[string]string)
			return
		}
	}
	if this.IsAjax() {
		this.Data["json"] = "No access!"
		this.ServeJson()
	} else {
		this.Redirect(fmt.Sprintf("/manage/login?redirect=%s", this.Ctx.Request.URL.String()), 302)
	}
	this.StopRun()
}

func (this *BaseController) Get() {
}

func (this *BaseController) Post() {

}

func (this *BaseController) CheckAvaliable(avaliables string) {
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	//获取用户权限成功
	if userAvaliables := sess.Get("avaliables"); userAvaliables != nil {
		var strAvaliable string = userAvaliables.(string)
		if strAvaliable != "" {
			avaliablesList := strings.Split(strAvaliable, ",")
			//分析所需权限
			needAvaliablesList := strings.Split(avaliables, ",")
			//所需权限个数
			needAvaliableCount := len(needAvaliablesList)
			//已匹配权限个数
			paired := 0
			for _, needAvaliable := range needAvaliablesList {
				for _, v := range avaliablesList {
					//权限匹配
					if needAvaliable == v {
						paired = paired + 1
						break
					}
				}
			}
			if paired == needAvaliableCount {
				return
			}
		}
	}

	if this.IsAjax() {
		this.Data["json"] = "No access!"
	} else {
		this.Redirect(fmt.Sprintf("/manage/login?redirect=%s", this.Ctx.Request.URL.String()), 302)
	}

	this.StopRun()
}

func (this *BaseController) ResponseError(err error) {
	if this.IsAjax() {
		this.Data["json"] = &err
		this.ServeJson()
		this.StopRun()
	}
}

func (this *BaseController) GetPageHeader() string {
	return this.pageHeader
}

func (this *BaseController) SetPageHeader(pageHeader string) {
	this.pageHeader = pageHeader
	this.Data["page_header"] = this.pageHeader
}

func (this *BaseController) GetPageHeaderSmall() string {
	return this.pageHeader
}

func (this *BaseController) SetPageHeaderSmall(pageHeaderSmall string) {
	this.pageHeaderSmall = pageHeaderSmall
	this.Data["page_header_small"] = this.pageHeaderSmall
}

//如果传入字符串，则原样放回，否则转换成json对象后返回
func (this *BaseController) ResponseData(data interface{}) {
	var toResponse []byte
	var err error
	switch data.(type) {
	case int:
		toResponse = []byte(strconv.Itoa(data.(int)))
	case int32:
		toResponse = []byte(strconv.Itoa(int(data.(int32))))
	case int64:
		toResponse = []byte(strconv.Itoa(int(data.(int64))))
	case string:
		toResponse = []byte(data.(string))
	case []byte:
		toResponse = data.([]byte)
	default:
		toResponse, err = json.Marshal(data)
		if err != nil {
			log.Println(err.Error())
		}
	}

	this.Ctx.ResponseWriter.Write(toResponse)
}

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}
