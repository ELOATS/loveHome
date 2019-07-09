package controllers

import(
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *UserController) Reg(){
	resp:=make(map[string]interface{})
	defer this.RetData(resp)

	//get the JSON data passed from the front end
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	/*beego.Info(`resp["mobile"]=`,resp["mobile"])
	beego.Info(`resp["password]="`,resp["password"])
	beego.Info(`resp["sms_code"]=`,resp["sms_code"])*/

	//insert the database
	o:=orm.NewOrm()
	user:=models.User{}
	user.Password_hash=resp["password"].(string)
	user.Name=resp["mobile"].(string)
	user.Mobile=resp["mobile"].(string)

	id,err:=o.Insert(&user)
	if err!=nil{
		resp["errno"]=models.RECODE_USERERR
		resp["errmsg"]="registration failed"
		return
	}
	beego.Info("registration successfully,id=",id)
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]="registration successfully"

	this.SetSession("name",user.Name)
}
