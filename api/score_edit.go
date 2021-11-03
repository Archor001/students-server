package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xormplus/core"
)

type ScoreEditRequest struct{
	Sno string `json:"sno"`
	Cid int `json:"cid"`
	Score float64 `json:"score"`
}

func (imp *impOne)ScoreEditWork(c *gin.Context,f ScoreEditRequest)Score{
	score:=&Score{}
	score.Sno=f.Sno
	score.Cid=f.Cid
	score.Score=f.Score
	checkscore:=new(Score)
	has,err:=imp.engine.ID(core.PK{score.Sno,score.Cid}).Get(checkscore)
	if err!=nil{
		errormsg_score("成绩编辑失败",nil,nil,c)
		return Score{}
	}
	if has==false{  //成绩不存在
		affected,err:=imp.engine.Insert(score)
		if err!=nil{
			errormsg_score("成绩新增失败",score,nil,c)
			return Score{}
		}
		log.WithFields(log.Fields{"code":Ok,"affected rows":affected,"sno":score.Sno,"sid":score.Cid,"score":score.Score}).Info("成绩新增成功")
	}else {
		affected, err := imp.engine.ID(core.PK{score.Sno, score.Cid}).Update(score)
		if err != nil {
			errormsg_score("成绩修改失败", score, err, c)
			return Score{}
		}
		log.WithFields(log.Fields{"code": Ok, "affected rows": affected, "sno": score.Sno, "sid": score.Cid, "score": score.Score}).Info("成绩修改成功")
	}
	rnt:=Score{
		Cid: score.Cid,
		Sno: score.Sno,
		Score: score.Score,
	}
	return rnt
}

func (imp *impOne)ScoreEdit(c *gin.Context){
	f:=ScoreEditRequest{}
	if err:=c.ShouldBindJSON(&f);err!=nil{
		errormsg_score("json解析失败",nil,nil,c)
		return
	}
	score:=imp.ScoreEditWork(c,f)
	c.JSON(Ok,ScoreResponse{
		Code:Ok,
		Msg:"success",
		Sno:score.Sno,
		Cid:score.Cid,
		Score: score.Score,
	})
}