package controllers

import(
	"github.com/astaxie/beego"
	"loveHome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	user:=models.User{}
	/*user.Name="lightning"

	resp["errno"]=0
	resp["errmsg"]="OK"
	resp["data"]=user*/

	resp["errno"]=models.RECODE_DBERR
	resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)

	v:=this.GetSession("name")
	if v!=nil{
		user.Name=v.(string)
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		resp["data"]=user
	}
}

func (this *SessionController) DeleteSessionData(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
}
