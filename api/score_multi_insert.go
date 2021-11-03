package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ScoreMultiInsertResponse struct{
	Code int
	Msg string
	Data []Score
}

func (imp *impOne)ScoreMultiInsert(c *gin.Context){
	form, err := c.FormFile("file")
	if err != nil {
		errormsg_score("获取文件失败",nil,err,c)
		return
	}
	filename:=form.Filename
	f,err:=excelize.OpenFile(filename)
	if err != nil {
		errormsg_student("文件打开失败",nil,err,c)
		return
	}

	ans:=make([]Score,0)
	rows := f.GetRows("Sheet1")
	for i, row := range rows {
		if i==0 {
			continue
		}
		jid,err:=strconv.Atoi(row[1])
		if err!=nil{
			errormsg_student("批量添加成绩失败",nil,err,c)
			return
		}
		jscore,err:=strconv.ParseFloat(row[2],64)
		if err!=nil{
			errormsg_student("批量添加成绩失败",nil,err,c)
			return
		}
		f:=ScoreEditRequest{
			Sno: row[0],
			Cid: jid,
			Score: jscore,
		}
		score:=imp.ScoreEditWork(c,f)
		if (score==Score{}){
			errormsg_student("批量添加学生失败",nil,nil,c)
			return
		}
		ans=append(ans,score)
	}
	c.JSON(Ok,ScoreMultiInsertResponse{
		Code:Ok,
		Msg:"success",
		Data:ans,
	})
}