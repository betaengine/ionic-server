package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ionic_server/controllers:AppController"] = append(beego.GlobalControllerRouter["ionic_server/controllers:AppController"],
		beego.ControllerComments{
			"Update",
			`/update`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ionic_server/controllers:AppController"] = append(beego.GlobalControllerRouter["ionic_server/controllers:AppController"],
		beego.ControllerComments{
			"Check",
			`/check`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ionic_server/controllers:AppController"] = append(beego.GlobalControllerRouter["ionic_server/controllers:AppController"],
		beego.ControllerComments{
			"UploadCheck",
			`/uploadcheck`,
			[]string{"post"},
			nil})

}
