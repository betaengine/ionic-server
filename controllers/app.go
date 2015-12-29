package controllers
import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"encoding/json"
	"ionic_server/models"
	"fmt"
)

var log = logs.NewLogger(1000)
// Operations about App
type AppController struct {
	beego.Controller
}


// @Title 下载升级包
// @Success 200 下载成功
// @Failure 403 下载错误
// @router /update [get]
func (a *AppController)Update() {
	a.Ctx.Output.Download("update/www.zip")
}

// @Title 通过uuid获取最新版本
// @Description 通过比较uuid，比较客户端和服务器的是否相同
// @Param	uuid		query 	string	true		"The uuid for version"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /check [post]
func (a *AppController) Check() {
	var deviceInfo models.DeviceInfo
	json.Unmarshal(a.Ctx.Input.RequestBody, &deviceInfo)
	fmt.Println(deviceInfo.Device_deploy_uuid)
	fileName := "update/version.json"
	a.Ctx.Output.Download(fileName)
}

// @Title uploadCheck
// @Description
// @Success 200 上传成功
// @router /uploadcheck [post]
func (a *AppController) UploadCheck() {
	app_name := a.Input().Get("name")
	app_id := a.Input().Get("appid")
	json := `{"app_id": `+ app_id + `, "app_name":`+ app_name +`, "signed_request": "http://192.1.10
	.87:8080/api/app/uploadfile", "errors": []}`
	fmt.Println(json)
	a.Ctx.WriteString(json)
	a.ServeJson()
}
