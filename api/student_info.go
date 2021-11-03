package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"strconv"
)

type StudentInfoRequest struct{
	Cid int `json:"cid"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	ClassFilter []string `json:"classFilter"`
	SexFilter []string `json:"sexFilter"`
	GradeFilter []int `json:"gradeFilter"`
	ScoreOrder string `json:"scoreOrder"`
}

type StudentInfo struct{
	Student `xorm:"extends"`
	Score float64
}

type StudentInfoResponse struct{
	Code int
	Msg string
	ClassOption []string
	GradeOption []int
	Data []StudentInfo
	Total int
}

func (imp *impOne)myIn(f StudentInfoRequest,typ int) *xorm.Session{
	var ses *xorm.Session
	if typ==0 {
		ses= imp.engine.Table("students")
	}else{
		ses=imp.engine.Table("students").Select("students.*,(case when scores.score is null then 0 else scores.score end) as score")
	}
	if len(f.ClassFilter)!=0{
		ses=ses.In("sclass",f.ClassFilter)
	}
	if len(f.SexFilter)!=0{
		ses=ses.In("ssex",f.SexFilter)
	}
	if len(f.GradeFilter)!=0{
		ses=ses.In("sgrade",f.GradeFilter)
	}
	return ses;
}

func (imp *impOne)StudentInfo(c *gin.Context){
	f:=StudentInfoRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_student("json解析失败",nil,nil,c)
		return
	}
	begin:=(f.Page-1)*f.PageSize
	studentinfo:=make([]StudentInfo,0)
	var err1,err2,err error
	var tot int64
	classop:=make([]string,0)
	gradeop:=make([]int,0)

	if f.Cid==Zero{//api2
		student:=make([]Student,0)
		err1 = imp.myIn(f,0).Find(&student)
		tot,err2 = imp.myIn(f,0).Count(&student)
		if err1!=nil{
			errormsg_student("学生查询失败#1",nil,err1,c)
			return
		}
		if err2!=nil{
			errormsg_student("学生查询失败#2",nil,err2,c)
			return
		}
		for _,val:=range student{
			studentinfo=append(studentinfo,StudentInfo{Student:val})
		}

		option:=make([]Student,0)
		err=imp.myIn(f,0).Distinct("sclass").Find(&option)
		if err!=nil{
			errormsg_student("学生查询失败#3",nil,err,c)
			return
		}
		for _,val:=range option{
			classop=append(classop,val.Sclass)
		}
		option=make([]Student,0)
		err=imp.myIn(f,0).Distinct("sgrade").Find(&option)
		if err!=nil{
			errormsg_student("学生查询失败#3",nil,err,c)
			return
		}
		for _,val:=range option{
			gradeop=append(gradeop,val.Sgrade)
		}

		log.WithFields(log.Fields{"code":Ok,"page":f.Page,"pagesize":f.PageSize,"classfilter":f.ClassFilter,"sexfilter":f.SexFilter,"gradefilter":f.GradeFilter,"order":f.ScoreOrder}).Info("学生信息查询成功")
	}else{//api3
		courseForClass:=&Course{}
		_,err:=imp.engine.Where("cid=?",f.Cid).Get(courseForClass)
		if err!=nil{
			errormsg_student("学生信息查询失败",nil,nil,c)
			return
		}
		classop,_=getCludentSlice(courseForClass.ElectClass,courseForClass.RetakeStudent)
		tot,err2=imp.myIn(f,0).In("sclass",classop).Join("LEFT OUTER","scores","students.sno=scores.sno and scores.cid=?",f.Cid).Count(&studentinfo)
		if err2!=nil{
			errormsg_student("此课程的所有学生信息查询失败#2",nil,err2,c)
			return
		}
		if f.ScoreOrder=="asc"{
			err1=imp.myIn(f,1).Join("LEFT OUTER","scores","students.sno=scores.sno and scores.cid=?",f.Cid).In("sclass",classop).Asc("score").Limit(f.PageSize,begin).Find(&studentinfo)
		}else if f.ScoreOrder=="desc"{
			err1=imp.myIn(f,1).Join("LEFT OUTER","scores","students.sno=scores.sno and scores.cid=?",f.Cid).In("sclass",classop).Desc("score").Limit(f.PageSize,begin).Find(&studentinfo)
		}else{
			err1=imp.myIn(f,1).Join("LEFT OUTER","scores","students.sno=scores.sno and scores.cid=?",f.Cid).In("sclass",classop).Limit(f.PageSize,begin).Find(&studentinfo)
		}
		if err1!=nil{
			errormsg_student("此课程的所有学生信息查询失败#1",nil,err1,c)
			return
		}

		option:=make([]Student,0)
		err=imp.myIn(f,0).In("sclass",classop).Distinct("sgrade").Find(&option)
		if err!=nil{
			errormsg_student("学生查询失败#3",nil,err,c)
			return
		}
		for _,val:=range option{
			gradeop=append(gradeop,val.Sgrade)
		}
		log.WithFields(log.Fields{"code":Ok,"page":f.Page,"pagesize":f.PageSize,"classfilter":f.ClassFilter,"sexfilter":f.SexFilter,"gradefilter":f.GradeFilter,"order":f.ScoreOrder}).Info("该课程的所有学生信息查询成功")
	}
	total,err:=strconv.Atoi(strconv.FormatInt(tot,10))
	if err!=nil{
		errormsg_course("学生信息查询失败#4",nil,err,c)
		return
	}
	c.JSON(Ok,StudentInfoResponse{
		Code: Ok,
		Msg: "success",
		ClassOption: classop,
		GradeOption: gradeop,
		Data: studentinfo,
		Total: total,
	})
}