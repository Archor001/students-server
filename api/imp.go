package api

import "github.com/gin-gonic/gin"

const Ok int=200
const BadRequest int=400
const Disconnected int=401
const Unauthorized int=402
const Zero int=0
const defaultmm string="123456"

type Course struct {
	Cid      int `xorm:"pk autoincr"`
	Cname    string `xorm:"not null unique"`
	Ccredit  float64
	ElectClass string
	RetakeStudent string
}

type Student struct {
	Sid    int `xorm:"pk autoincr"`
	Sno    string `xorm:"not null unique"`
	Sname  string
	Ssex   string
	Sage   int
	Sgrade int `xorm:"not null"`
	Sclass string
}

type Score struct{
	Sno string `xorm:"pk"`
	Cid int `xorm:"pk"`
	Score float64
}

type User struct{
	Uid int `xorm:"pk autoincr"`
	Username string
	Password string
	Role string
}

type Authority struct{
	Role string `xorm:"pk"`
	Auth string
}

type IMP interface{
	CourseInfo(c *gin.Context)
	CourseEdit(c *gin.Context)
	CourseDelete(c *gin.Context)
	StudentInfo(c *gin.Context)
	StudentEdit(c *gin.Context)
	StudentDelete(c *gin.Context)
	StudentMultiInsert(c *gin.Context)
	ScoreEdit(c *gin.Context)
	ScoreMultiInsert(c *gin.Context)
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)

}