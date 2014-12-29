package weixin

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/julycw/Jeremiah/models"
	"github.com/julycw/Jeremiah/weixin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	siteIp string
)

const (
	MSG_TYPE_TEXT  = "text"
	MSG_TYPE_IMG   = "image"
	MSG_TYPE_VOICE = "voice"
	MSG_TYPE_VIDEO = "video"
	MSG_TYPE_NEWS  = "news"
	MSG_TYPE_EVENT = "event"
)

type BaseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   //开发者微信号
	FromUserName string   //发送者，一个OpenId
	CreateTime   string   //创建时间
	MsgId        string   //消息id，64位整型
	MsgType      string   //image
}

type TextMessage struct { //文本
	BaseMessage
	XMLName xml.Name `xml:"xml"`
	Content string   //文本内容
}

type EventMessage struct { //文本
	BaseMessage
	XMLName xml.Name `xml:"xml"`
	Event   string   //文本内容
}

type WeixinController struct {
	beego.Controller
}

type NewsMessage struct {
	BaseMessage
	ArticleCount int
	Articles     Articles
}

type Articles struct {
	XMLName xml.Name `xml:"Articles"`
	Items   []Item
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

func (this *WeixinController) Get() {
	if ret, echostr := weixin.CheckSignature(this.Ctx.ResponseWriter, this.Ctx.Request); ret == true {
		fmt.Fprintf(this.Ctx.ResponseWriter, echostr)
	} else {
		log.Println(this.Ctx.ResponseWriter, "invalid!")
	}
}

func (this *WeixinController) Post() {
	if ret, _ := weixin.CheckSignature(this.Ctx.ResponseWriter, this.Ctx.Request); ret == true {
		handleRequest(this.Ctx.ResponseWriter, this.Ctx.Request)
	} else {
		log.Println(this.Ctx.ResponseWriter, "invalid!")
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadAll(r.Body)
	var bm BaseMessage
	xml.Unmarshal(content, &bm)
	switch bm.MsgType {
	case MSG_TYPE_TEXT:
		handleTextMsg(w, content)
	case MSG_TYPE_EVENT:
		handleEventMsg(w, content)
	default:
		return
	}
}

func handleTextMsg(w http.ResponseWriter, msg []byte) {
	var message TextMessage
	if err := xml.Unmarshal(msg, &message); err != nil {
		log.Println(err.Error())
		return
	}

	user := models.FollowUser{
		UserId:      message.FromUserName,
		UserName:    "",
		NickName:    "",
		Description: "",
	}

	//将msg存入数据库
	var db_msg models.Message
	db_msg.UserId = message.FromUserName
	db_msg.Content = message.Content
	db_msg.MsgType = message.MsgType
	db_msg.MsgId = message.MsgId

	if _, err := models.Orm.Insert(&db_msg); err != nil {
		log.Println(err.Error())
	}

	user.LastMessageOn = time.Now()

	if created, _, err := models.Orm.ReadOrCreate(&user, "UserId"); err == nil {
		if !created {
			user.LastMessageOn = time.Now()
			models.Orm.Update(&user)
		} else {
		}
	} else {
		log.Println(err.Error())
	}

	//

	switch {
	default:
		makeDefaultContent(&message)
	}

	message.FromUserName, message.ToUserName = message.ToUserName, message.FromUserName
	ret, _ := xml.Marshal(message)
	w.Write(ret)
}

func makeDefaultContent(message *TextMessage) {
	var option models.Option
	message.MsgType = MSG_TYPE_TEXT
	option.Key = "msg_feedback_text"
	if err := models.Orm.Read(&option, "Key"); err == nil {
		message.Content = fmt.Sprintf(option.Value, strings.TrimSpace(message.Content))
	} else {
		message.Content = fmt.Sprintf("您刚才说“%v”，小贝已经清楚地记下了！", strings.TrimSpace(message.Content))
	}
	message.Content += "您还可以...\n"
	message.Content += fmt.Sprintf("\n<a href=\"http://%s/weixin/viewOrder?userID=%s\">查询我的订单</a>", siteIp, message.FromUserName)
	message.Content += fmt.Sprintf("\n<a href=\"http://%s/weixin/user?userID=%s\">完善个人信息</a>", siteIp, message.FromUserName)
	message.Content += fmt.Sprintf("\n<a href=\"http://%s/weixin/bindOrderNumber?userID=%s\">绑定订单</a>", siteIp, message.FromUserName)

}

func handleEventMsg(w http.ResponseWriter, msg []byte) {
	var message EventMessage
	if err := xml.Unmarshal(msg, &message); err != nil {
		log.Println(err.Error())
		return
	}

	user := models.FollowUser{
		UserId:      message.FromUserName,
		UserName:    "",
		NickName:    "",
		Description: "",
	}

	var ret []byte
	//订阅
	if message.Event == "subscribe" {
		var feedback TextMessage
		var option models.Option
		option.Key = "follow_feedback_text"
		if err := models.Orm.Read(&option, "Key"); err == nil {
			feedback.Content = option.Value
		} else {
			feedback.Content = "欢迎关注贝贝科技生活馆，您有什么想和小贝说的呢？"
		}
		feedback.FromUserName = message.ToUserName
		feedback.ToUserName = message.FromUserName
		feedback.MsgType = MSG_TYPE_TEXT
		ret, _ = xml.Marshal(feedback)

		user.FollowOn = time.Now()

		if created, _, err := models.Orm.ReadOrCreate(&user, "UserId"); err == nil {
			if !created {
				user.FollowOn = time.Now()
				models.Orm.Update(&user)
			} else {
			}
		} else {
			log.Println(err.Error())
		}

	} else if message.Event == "unsubscribe" { //退订
		user.UnFollowOn = time.Now()
		if created, _, err := models.Orm.ReadOrCreate(&user, "UserId"); err == nil {
			if !created {
				user.UnFollowOn = time.Now()
				models.Orm.Update(&user)
			} else {
			}
		} else {
			log.Println(err.Error())
		}
	}
	w.Write(ret)
}

func init() {
	siteIp = beego.AppConfig.String("siteip")
}
