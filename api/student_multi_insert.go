package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"strconv"
)

type StudentMultiInsertResponse struct{
	Code int
	Msg string
	Data []Student
}

func (imp *impOne)StudentMultiInsert(c *gin.Context){
	form, err := c.FormFile("file")
	if err != nil {
		errormsg_student("获取文件失败",nil,err,c)
		return
	}
	filename:=form.Filename
	f,err:=excelize.OpenFile(filename)
	if err != nil {
		errormsg_student("文件打开失败",nil,err,c)
		return
	}

	ans:=make([]Student,0)
	rows := f.GetRows("Sheet1")
	for i, row := range rows {
		if i==0 {
			continue
		}
		jage,err:=strconv.Atoi(row[4])
		if err!=nil{
			errormsg_student("批量添加学生失败",nil,err,c)
			return
		}
		jgrade,err:=strconv.Atoi(row[5])
		if err!=nil{
			errormsg_student("批量添加学生失败",nil,err,c)
			return
		}
		f:=StudentEditRequest{
			Sno: row[1],
			Sname: row[2],
			Ssex: row[3],
			Sage: jage,
			Sgrade: jgrade,
			Sclass: row[6],
		}
		stu:=imp.StudentEditWork(c,f)
		if (stu==Student{}){
			errormsg_student("批量添加学生失败",nil,nil,c)
			return
		}
		ans=append(ans,stu)
	}
	c.JSON(Ok,StudentMultiInsertResponse{
		Code:Ok,
		Msg:"success",
		Data:ans,
	})
}