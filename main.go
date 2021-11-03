package main

import (
	"github.com/Archor001/students-server/api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func log_init(){
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main(){
	log_init()
	imp:=api.NewIMPone()
	r:=gin.Default()
	r.POST("/students/editCourse",imp.CourseEdit)
	r.POST("/students/deleteCourse",imp.CourseDelete)
	r.POST("/students/queryCourse",imp.CourseInfo)
	r.POST("/students/editStudent",imp.StudentEdit)
	r.POST("/students/deleteStudent",imp.StudentDelete)
	r.POST("/students/queryStudent",imp.StudentInfo)
	r.POST("/students/multi_editStudent",imp.StudentMultiInsert)
	r.POST("/students/editScore",imp.ScoreEdit)
	r.POST("/students/multi_insertScore",imp.ScoreMultiInsert)
	r.Run(":8300")
}