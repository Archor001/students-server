package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StudentDeleteRequest struct{
	Sno string `json:"sno,required"`
}

func (imp *impOne)StudentDelete(c *gin.Context){
	f:=StudentDeleteRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_student("json解析失败",nil,nil,c)
		return
	}
	//学生表中根据id找到元组信息
	stu:=new(Student)
	stu.Sno=f.Sno
	has,err:=imp.engine.Where("sno=?",stu.Sno).Get(stu)
	if err!=nil{
		errormsg_student("fail to delete student_info",stu,err,c)
		return
	}
	if has==false{
		errormsg_student("student doesn't exist",stu,err,c)
		return
	}

	affected,err:=imp.engine.Where("sno=?",stu.Sno).Delete(stu)  //学生表删除
	if err!=nil{
		errormsg_student("fail to delete student_info(1)",stu,err,c)
		return
	}
	log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"sid":stu.Sid,"sno":stu.Sno,"sname":stu.Sname,"ssex":stu.Ssex,"sage":stu.Sage,"sgrade":stu.Sgrade,"sclass":stu.Sclass}).Info("student_info has been deleted")

	user:=new(User)
	affected,err=imp.engine.Where("username=?",stu.Sno).Delete(user)  //用户表删除
	if err!=nil{
		errormsg_student("fail to delete student_info(2)",stu,err,c)
		return
	}
	log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"username":stu.Sno}).Info("student_info has been deleted in userlist")
	c.JSON(Ok,StudentResponse{
		Code: Ok,
		Msg:"success",
		Sid:stu.Sid,
		Sno:stu.Sno,
		Sname:stu.Sname,
		Ssex: stu.Ssex,
		Sage: stu.Sage,
		Sgrade: stu.Sgrade,
		Sclass:stu.Sclass,
	})
}