package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *AreaController) GetArea(){
	beego.Info("connect successfully")
	resp:=make(map[string]interface{})

	//packaged into JSON and returned to the front end
	defer this.RetData(resp)
	//get data from the session

	//get area data from the mysql database
	o:=orm.NewOrm()
	area:=[]models.Area{}
	num,err:=o.QueryTable("area").All(&area)
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num==0{
		resp["errno"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	resp["data"]=area
	beego.Info("query data successfully,resp=",resp,"num=",num)
}
