package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type CourseInfoRequest struct{
	Sno string `json:"sno"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	PointOrder string `json:"pointOrder"`
}

type CourseInfo struct{
	Course `xorm:"extends"`
	Score float64
}

type CourseInfoResponse struct{
	Code int
	Msg string
	Data []CrsResponse
	Total int
}

func (imp *impOne)CourseInfo(c *gin.Context){
	f:=CourseInfoRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_course("json解析失败",nil,nil,c)
		return
	}
	begin:=(f.Page-1)*f.PageSize
	courseinfo :=make([]CourseInfo,0)
	var err1,err2 error
	var tot int64
	if f.Sno==""{//api1
		course:=make([]Course,0)
		tot,err2=imp.engine.Table("courses").Count(&course)
		if f.PointOrder=="asc"{
			err1=imp.engine.Table("courses").Asc("ccredit").Limit(f.PageSize,begin).Find(&course)
		}else if f.PointOrder=="desc"{
			err1=imp.engine.Table("courses").Desc("ccredit").Limit(f.PageSize,begin).Find(&course)
		}else {
			err1=imp.engine.Table("courses").Limit(f.PageSize,begin).Find(&course)
		}
		if err1!=nil{
			errormsg_course("课程查询失败#1",nil,err1,c)
			return
		}
		if err2!=nil{
			errormsg_course("课程查询失败#2",nil,err2,c)
			return
		}
		for _,val:=range course{
			courseinfo = append(courseinfo, CourseInfo{Course: val});
		}
		log.WithFields(log.Fields{"code":Ok,"page":f.Page,"pagesize":f.PageSize,"order":f.PointOrder}).Info("课程信息查询成功")
	}else{//api4
		tot,err2=imp.engine.Table("courses").Join("INNER","scores","sno=? and scores.cid=courses.cid",f.Sno).Count(&courseinfo)
		if err2!=nil{
			errormsg_course("此学生的所有课程信息查询失败#2",nil,err2,c)
			return
		}
		if f.PointOrder=="asc"{
			err1=imp.engine.Table("courses").Join("INNER","scores","sno=? and scores.cid=courses.cid",f.Sno).Asc("ccredit").Limit(f.PageSize,begin).Find(&courseinfo)
		}else if f.PointOrder=="desc"{
			err1=imp.engine.Table("courses").Join("INNER","scores","sno=? and scores.cid=courses.cid",f.Sno).Desc("ccredit").Limit(f.PageSize,begin).Find(&courseinfo)
		}else{
			err1=imp.engine.Table("courses").Join("INNER","scores","sno=? and scores.cid=courses.cid",f.Sno).Limit(f.PageSize,begin).Find(&courseinfo)
		}
		if err1!=nil{
			errormsg_course("此学生的所有课程信息查询失败#1",nil,err1,c)
			return
		}
		log.WithFields(log.Fields{"code":Ok,"page":f.Page,"pagesize":f.PageSize,"order":f.PointOrder}).Info("此学生的所有课程信息查询成功")
	}
	total,err:=strconv.Atoi(strconv.FormatInt(tot,10))
	if err!=nil{
		errormsg_course("课程信息查询失败#3",nil,err,c)
		return
	}
	ans:=courseTransform(courseinfo)
	c.JSON(Ok,CourseInfoResponse{
		Code:Ok,
		Msg:"success",
		Data:ans,
		Total:total,
	})
}