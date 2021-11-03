package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CourseEditRequest struct{
	Cid int       `json:"cid"`
	Cname string     `json:"cname"`
	Ccredit float64      `json:"ccredit"`
	ElectClass []string  `json:"electClass"`
	RetakeStudent []string `json:"retakeStudent"`
}

func (imp *impOne)checkclassexist(class []string)bool{
	for _,c:=range class{
		tmp:=new(Student)
		total,err:=imp.engine.Where("sclass=?",c).Count(tmp)
		if err!=nil||total==0{
			return false
		}
	}
	return true
}

func (imp *impOne)checkstudentexist(student []string)bool{
	for _,s:=range student{
		tmp:=new(Student)
		total,err:=imp.engine.Where("sno=?",s).Count(tmp)
		if err!=nil||total==0{
			return false
		}
	}
	return true
}

func (imp *impOne)CourseEdit(c *gin.Context){
	f:=CourseEditRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_course("json解析失敗",nil,nil,c)
		return
	}
	if imp.checkclassexist(f.ElectClass)==false{
		errormsg_course("选课班级不存在，编辑课程失败",nil,nil,c)
		return
	}
	if imp.checkstudentexist(f.RetakeStudent)==false {
		errormsg_course("重修学生不存在，编辑课程失败",nil,nil,c)
		return
	}
	course:=&Course{}
	course.Cid=f.Cid
	course.Cname=f.Cname
	course.Ccredit=f.Ccredit
	course.ElectClass=slicetostring(f.ElectClass)
	course.RetakeStudent=slicetostring(f.RetakeStudent)
	if course.Cid==Zero{  //insert
		if course.Cname==""{
			errormsg_course("缺少课程名称",course,nil,c)
			return
		}
		affected,err:=imp.engine.Insert(course)
		if err!=nil {
			errormsg_course("新增课程失败",course,err,c)
			return
		}
		log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"cid":course.Cid,"cname":course.Cname,"ccredit":course.Ccredit,"elect_class":course.ElectClass,"retake_student":course.RetakeStudent}).Info("课程新增成功")
	}else{ //update
		checkcourse:=new(Course)
		has,err:=imp.engine.ID(course.Cid).Get(checkcourse)
		if has==false{
			errormsg_course("课程信息不存在，课程更新失败",course,err,c)
			return
		}
		affected,err:=imp.engine.ID(course.Cid).Update(course)
		if err!=nil{
			errormsg_course("课程更新失败",course,err,c)
			return
		}
		log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"cid":course.Cid,"cname":course.Cname,"ccredit":course.Ccredit,"elect_class":course.ElectClass,"retake_student":course.RetakeStudent}).Info("课程更新成功")
	}
	ec,rs:=getCludentSlice(course.ElectClass,course.RetakeStudent)
	c.JSON(Ok, CourseResponse{
		Code: Ok,
		Msg:  "Success",
		CrsResponse: CrsResponse{
			Cid:  course.Cid,
			Cname: course.Cname,
			Ccredit: course.Ccredit,
			ElectClass: ec,
			RetakeStudent: rs,
		},
	})
}