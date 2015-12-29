package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/app","description":"Operations about App\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/app":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/app","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/update","description":"","operations":[{"httpMethod":"GET","nickname":"下载升级包","type":"","responseMessages":[{"code":200,"message":"下载成功","responseModel":""},{"code":403,"message":"下载错误","responseModel":""}]}]},{"path":"/check","description":"","operations":[{"httpMethod":"POST","nickname":"通过uuid获取最新版本","type":"","summary":"通过比较uuid，比较客户端和服务器的是否相同","parameters":[{"paramType":"query","name":"uuid","description":"\"The uuid for version\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success","responseModel":""},{"code":403,"message":"user not exist","responseModel":""}]}]},{"path":"/uploadcheck","description":"","operations":[{"httpMethod":"POST","nickname":"uploadCheck","type":"","responseMessages":[{"code":200,"message":"上传成功","responseModel":""}]}]}]}}`
    BasePath string= "/api"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	if beego.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocApi["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.Apis {
				a.Path = urlReplace(k + a.Path)
				v.Apis[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocApi[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
