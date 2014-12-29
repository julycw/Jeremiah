package weixin

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/julycw/Jeremiah/weixin/data_package_struct"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WeixinSettings struct {
	APPID                      string
	APPSecret                  string
	Token                      string
	CheckUpAccessTokenInterval time.Duration
}

type WeixinGloble struct {
	AccessToken          string
	AccessTokenExpiresIn time.Duration
}

var globalData WeixinGloble
var globalSetting WeixinSettings

func Str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func CheckSignature(w http.ResponseWriter, r *http.Request) (bool, string) {
	r.ParseForm()
	var token string = globalSetting.Token
	var signature string = strings.Join(r.Form["signature"], "")
	var timestamp string = strings.Join(r.Form["timestamp"], "")
	var nonce string = strings.Join(r.Form["nonce"], "")
	var echostr string = strings.Join(r.Form["echostr"], "")
	tmps := []string{token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := Str2sha1(tmpStr)
	if tmp == signature {
		return true, echostr
	} else {
		return false, ""
	}
}

func manageAccessToken() {
	for {
		globalData.AccessTokenExpiresIn = globalData.AccessTokenExpiresIn - globalSetting.CheckUpAccessTokenInterval
		if globalData.AccessTokenExpiresIn < globalSetting.CheckUpAccessTokenInterval {
			response, _ := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", globalSetting.APPID, globalSetting.APPSecret))
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(body))
			var errPkg ErrPackage = ErrPackage{}
			if err := json.Unmarshal(body, &errPkg); err == nil {
				if errPkg.ErrCode == 0 {
					var accessToken AccessTokenPackage
					json.Unmarshal(body, &accessToken)
					fmt.Println(accessToken)
				}
			}
			globalData.AccessTokenExpiresIn = 7200 * time.Second
		}
		time.Sleep(globalSetting.CheckUpAccessTokenInterval)
	}
}

func init() {
	globalSetting.APPID = beego.AppConfig.String("appid")
	globalSetting.APPSecret = beego.AppConfig.String("appsecret")
	globalSetting.Token = beego.AppConfig.String("token")
	globalSetting.CheckUpAccessTokenInterval = 60 * time.Second

	//go manageAccessToken()
}
