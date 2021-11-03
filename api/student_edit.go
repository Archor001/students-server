package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StudentEditRequest struct{
	Sno string `json:"sno"`
	Sname string `json:"sname"`
	Ssex string `json:"ssex"`
	Sage int `json:"sage"`
	Sgrade int `json:"sgrade"`
	Sclass string `json:"sclass"`
}

func (imp *impOne)StudentEditWork(c *gin.Context,f StudentEditRequest)Student{
	stu:=new(Student)
	stu.Sno=f.Sno
	stu.Sname=f.Sname
	stu.Ssex=f.Ssex
	stu.Sage=f.Sage
	stu.Sgrade=f.Sgrade
	stu.Sclass=f.Sclass
	if stu.Sno!=""{ //update
		checkstu:=new(Student)
		has,err:=imp.engine.Where("sno=?",stu.Sno).Get(checkstu)
		if has==false{
			errormsg_student("学生信息不存在，更新失败",stu,err,c)
			return Student{}
		}
		affected,err:=imp.engine.Where("sno=?",stu.Sno).Update(stu)
		if err!=nil{
			errormsg_student("学生信息更新失败",stu,err,c)
			return Student{}
		}
		log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"sid":stu.Sid,"sno":stu.Sno,"sname":stu.Sname,"ssex":stu.Ssex,"sage":stu.Sage,"sgrade":stu.Sgrade,"sclass":stu.Sclass}).Info("学生信息已经更新")
	}else{ //insert
		if stu.Sgrade==0{
			errormsg_student("缺少学生年级信息，编辑失败",stu,nil,c)
			return Student{}
		}
		rtn,err:=imp.UserInsert(stu,nil)
		if rtn==false{
			errormsg_student("学生信息添加失败",stu,err,c)
			return Student{}
		}
		affected,err:=imp.engine.Insert(stu)
		if err!=nil{
			errormsg_student("学生信息添加失败",stu,err,c)
			return Student{}
		}
		log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"sid":stu.Sid,"sno":stu.Sno,"sname":stu.Sname,"ssex":stu.Ssex,"sage":stu.Sage,"sgrade":stu.Sgrade,"sclass":stu.Sclass}).Info("学生信息已添加")
	}
	rnt:=Student{
		Sid: stu.Sid,
		Sno: stu.Sno,
		Sname: stu.Sname,
		Ssex: stu.Ssex,
		Sage: stu.Sage,
		Sgrade: stu.Sgrade,
		Sclass: stu.Sclass,
	}
	return rnt
}

func (imp *impOne)StudentEdit(c *gin.Context){
	f:=StudentEditRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_student("json解析失败",nil,nil,c)
		return
	}
	stu:=imp.StudentEditWork(c,f)
	if (stu==Student{}){
		errormsg_student("学生编辑失败",nil,nil,c)
		return
	}
	c.JSON(Ok,StudentResponse{
		Code:Ok,
		Msg:"success",
		Sid:stu.Sid,
		Sno:stu.Sno,
		Sname: stu.Sname,
		Ssex:stu.Ssex,
		Sage:stu.Sage,
		Sgrade: stu.Sgrade,
		Sclass: stu.Sclass,
	})
}