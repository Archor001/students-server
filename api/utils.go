package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

type CrsResponse struct {
	Cid           int
	Cname         string
	Ccredit       float64
	ElectClass    []string
	RetakeStudent []string
	Score         float64
}

type CourseResponse struct {
	Code int
	Msg  string
	CrsResponse
}

type StudentResponse struct {
	Code   int
	Msg    string
	Sid    int
	Sno    string
	Sname  string
	Ssex   string
	Sage   int
	Sgrade int
	Sclass string
}

type ScoreResponse struct {
	Code  int
	Msg   string
	Sno   string
	Cid   int
	Score float64
}

func getCludentSlice(ec string, rs string) ([]string, []string) {
	str1 := make([]string, 0)
	str2 := make([]string, 0)
	if ec != "" {
		str1 = strings.Split(ec, "|")
	}
	if rs != "" {
		str2 = strings.Split(rs, "|")
	}
	return str1, str2
}

func courseTransform(courseinfo []CourseInfo) []CrsResponse { //课程表逻辑层到 持久层
	ans := make([]CrsResponse, 0)
	for _, val := range courseinfo {
		str1, str2 := getCludentSlice(val.ElectClass, val.RetakeStudent)
		ans = append(ans, CrsResponse{
			Cid:           val.Cid,
			Cname:         val.Cname,
			Ccredit:       val.Ccredit,
			ElectClass:    str1,
			RetakeStudent: str2,
			Score:         val.Score,
		})
	}
	return ans
}

func errormsg_course(str string, course *Course, err error, c *gin.Context) {
	log.WithFields(log.Fields{"code": BadRequest, "error": err}).Warn(str)
	if course != nil {
		c.JSON(BadRequest, CourseResponse{
			Code: BadRequest,
			Msg:  str,
			CrsResponse: CrsResponse{
				Cid:           course.Cid,
				Cname:         course.Cname,
				Ccredit:       course.Ccredit,
				ElectClass:    stringtoslice(course.ElectClass),
				RetakeStudent: stringtoslice(course.RetakeStudent),
			},
		})
	} else {
		c.JSON(BadRequest, gin.H{
			"Code": BadRequest,
			"Msg":  str,
		})
	}
}

func errormsg_student(str string, stu *Student, err error, c *gin.Context) {
	log.WithFields(log.Fields{"code": BadRequest, "error": err}).Warn(str)
	if stu != nil {
		c.JSON(BadRequest, StudentResponse{
			Code:   BadRequest,
			Msg:    str,
			Sid:    stu.Sid,
			Sno:    stu.Sno,
			Sname:  stu.Sname,
			Ssex:   stu.Ssex,
			Sage:   stu.Sage,
			Sgrade: stu.Sgrade,
			Sclass: stu.Sclass,
		})
	} else {
		c.JSON(BadRequest, gin.H{
			"Code": BadRequest,
			"Msg":  str,
		})
	}
}

func errormsg_score(str string, score *Score, err error, c *gin.Context) {
	log.WithFields(log.Fields{"code": BadRequest, "error": err}).Warn(str)
	if score != nil {
		c.JSON(BadRequest, ScoreResponse{
			Code:  BadRequest,
			Msg:   str,
			Sno:   score.Sno,
			Cid:   score.Cid,
			Score: score.Score,
		})
	} else {
		c.JSON(BadRequest, gin.H{
			"Code": BadRequest,
			"Msg":  str,
		})
	}
}

func slicetostring(str []string) string {
	num := len(str)
	if num == 0 {
		return ""
	}
	var bt bytes.Buffer
	for i := 0; i < num-1; i++ {
		bt.WriteString(str[i])
		bt.WriteString("|")
	}
	bt.WriteString(str[num-1])
	s := bt.String()
	return s
}

func stringtoslice(str string) []string {
	var rtn []string
	var tmp string
	for _, s := range str {
		if s == '|' {
			rtn = append(rtn, tmp)
			tmp = ""
			continue
		}
		tmp += string(s)
	}
	rtn = append(rtn, tmp)
	return rtn
}

func myhash(str string) string { //用户密码hash函数
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func removeRepStudent(s []Student) []Student { //切片去重
	rnt := make([]Student, 0)
	m := make(map[Student]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			rnt = append(rnt, v)
			m[v] = true
		}
	}
	return rnt
}

///////testhaha
