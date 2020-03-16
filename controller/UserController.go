package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"mkAllGo/common"
	"mkAllGo/model"
	"net/http"
	"time"
)
func Login(c *gin.Context) {
	DB := common.MongoInit("mk_more_all")
	var form model.User
	var result model.User
	if err := c.ShouldBind(&form); err == nil {
		err = DB.Find(bson.M{"name": form.Name}).One(&result)
		if err != nil {
			fmt.Printf("try find record error[%s]\n", err.Error())
		}
		if result.Password != form.Password {
			c.JSON(http.StatusOK, gin.H{"status": "账号或密码不正确"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "登陆成功", "taobaosid": result.TaobaoSid, "jingfenid": result.JingdongId, "extime":result.ExTime})
		}
	}else{
		c.JSON(http.StatusOK, gin.H{"status": "123213"})
	}
}
func Register(c *gin.Context){
	DB := common.MongoInit("mk_more_all")
	var form model.User
	var result model.User
	if err := c.ShouldBind(&form); err == nil{
		err = DB.Find(bson.M{"name":form.Name}).One(&result)
		if err != nil {
			fmt.Printf("try find record error[%s]\n", err.Error())
		}
		if result.Name != ""{
			c.JSON(http.StatusOK, gin.H{"status": "该用户名已经注册,请登录"})
		}else{
			form.RegTime = time.Now()
			dd, _ := time.ParseDuration("24h")
			form.ExTime = form.RegTime.Add(dd)
			err := DB.Insert(form)
			if err != nil {
				fmt.Printf("try insert record error[%s]\n", err.Error())
				err_handler(err)
			}else{
				c.JSON(http.StatusOK, gin.H{"status": "注册成功", "taobaosid":form.TaobaoSid ,"jingfenid":form.JingdongId, "extime":form.ExTime })
			}

		}

	}

}
func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())

}
