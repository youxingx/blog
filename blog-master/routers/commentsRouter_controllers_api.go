package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"],
		beego.ControllerComments{
			Method: "GetCaptcha",
			Router: `/captcha`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"],
		beego.ControllerComments{
			Method: "PostCaptchaCheck",
			Router: `/captcha/check`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["blog-master/controllers/api:FrontController"],
		beego.ControllerComments{
			Method: "GetTest",
			Router: `/test/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
