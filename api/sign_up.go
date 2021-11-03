package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (imp *impOne)SignUp(c *gin.Context){

}

func (imp *impOne)UserInsert(stu *Student,user *User) (bool,error){
	uuser:=new(User)
	if stu!=nil{ //学生
		maxno:=new(Student)
		has,err:=imp.engine.Where("sgrade = ?",stu.Sgrade).Desc("sno").Limit(1).Get(maxno)
		if err!=nil {
			return false,err
		}
		if has==false {
			stu.Sno = "U" + strconv.Itoa(stu.Sgrade) + "00001"
		}else{
			no,_:=strconv.Atoi(maxno.Sno[5:])
			no+=1
			s:=fmt.Sprintf("%05d",no)
			stu.Sno="U" + strconv.Itoa(stu.Sgrade) +s
		}
		uuser.Username=stu.Sno
		uuser.Role="学生"
		//hash两次
		uuser.Password=myhash(defaultmm)
		uuser.Password=myhash(uuser.Password)
	}else if user!=nil{
		//hash一次
		uuser=user
		uuser.Password=myhash(uuser.Password)
	}
	_,err:=imp.engine.Insert(uuser)
	if err!=nil{
		return false,err
	}else {
		return true,nil
	}
}