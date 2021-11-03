package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CourseDeleteRequest struct{
	Cid int `json:"cid"`
}

func (imp *impOne)CourseDelete(c *gin.Context){
	f:=CourseDeleteRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_course("json绑定失败",nil,nil,c)
		return
	}

	//在课程表中根据Cid找到元组
	course:=new(Course)
	course.Cid=f.Cid
	has,err:=imp.engine.ID(f.Cid).Get(course)
	if err!=nil{
		errormsg_course("未能找到课程信息，课程删除失败",course,err,c)
		return
	}
	if has==false{
		errormsg_course("课程信息不存在，课程删除失败",course,err,c)
		return
	}

	//课程表中删除元组
	affected,err:=imp.engine.ID(course.Cid).Delete(course)
	if err!=nil{
		errormsg_course("课程删除失败",course,err,c)
		return
	}
	log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"cid":course.Cid,"cname":course.Cname,"ccredit":course.Ccredit,"elect_class":course.ElectClass,"retake_student":course.RetakeStudent}).Info("课程删除成功")
	c.JSON(Ok,CourseResponse{
		Code: Ok,
		Msg:"success",
		CrsResponse: CrsResponse{
			Cid:course.Cid,
			Cname:course.Cname,
			Ccredit: course.Ccredit,
			ElectClass: stringtoslice(course.ElectClass),
			RetakeStudent: stringtoslice(course.RetakeStudent),
		},
	})
}