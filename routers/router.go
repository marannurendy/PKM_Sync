package routers

import (
	"beego/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/referral", &controllers.MainController{}, "post:PostTransaction")
	beego.Router("/referral", &controllers.MainController{}, "get:GetTransaction")
	beego.Router("/referral-inquiry", &controllers.MainController{}, "post:PostInquiry")
	beego.Router("/referral-update", &controllers.MainController{}, "post:UpdateTransaction")
	beego.Router("/webAuth", &controllers.MainController{}, "post:WebAuth")
	beego.Router("/listUmi", &controllers.MainController{}, "post:GetListNasabah")
	beego.Router("/listUmi/:transaction_id", &controllers.MainController{}, "get:GetListNasabahById")
	beego.Router("/sync-angsuran", &controllers.MainController{}, "post:SyncPKM_Angsuran")
	beego.Router("/sync-titipan", &controllers.MainController{}, "post:SyncPKM_Titipan")
}
