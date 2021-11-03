package api

import (
	log "github.com/sirupsen/logrus"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

type impOne struct {
	engine *xorm.Engine
}

func NewIMPone() IMP{
	imp:=new(impOne)
	engine,err :=xorm.NewEngine("mysql","root:132311@/xormtest?charset=utf8mb4")
	if err != nil {
		log.WithFields(log.Fields{
			"code":BadRequest,
			"error":err,
		}).Fatal("fail to connect database")
		return nil
	}
	imp.engine = engine
	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, "s")
	imp.engine.SetTableMapper(tbMapper)
	t1,_:=imp.engine.IsTableExist(&Course{})

	if t1==false {
		err=imp.engine.CreateTables(&Course{})
		if err!=nil{
			log.WithFields(log.Fields{"code":BadRequest, "error":err,}).Fatal("fail to create table")
			return nil
		}
		imp.engine.CreateUniques(&Course{})
	}
	t2,_:=imp.engine.IsTableExist(&Student{})
	if t2==false {
		imp.engine.CreateTables(&Student{})
		if err!=nil{
			log.WithFields(log.Fields{"code":BadRequest, "error":err,}).Fatal("fail to create table")
			return nil
		}
		imp.engine.CreateUniques(&Student{})
	}
	t3,_:=imp.engine.IsTableExist(&Score{})
	if t3==false {
		imp.engine.CreateTables(&Score{})
		if err!=nil{
			log.WithFields(log.Fields{"code":BadRequest, "error":err,}).Fatal("fail to create table")
			return nil
		}
	}
	t4,_:=imp.engine.IsTableExist(&User{})
	if t4==false{
		imp.engine.CreateTables(&User{})
		if err!=nil{
			log.WithFields(log.Fields{"code":BadRequest,"error":err}).Fatal("fail to create table")
			return nil
		}
	}
	t5,_:=imp.engine.IsTableExist(&Authority{})
	if t5==false{
		imp.engine.CreateTables(&Authority{})
		if err!=nil{
			log.WithFields(log.Fields{"code":BadRequest,"error":err}).Fatal("fail to create table")
			return nil
		}
	}
	return imp
}

