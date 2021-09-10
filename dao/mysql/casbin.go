package mysql

import (
	"github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
)

var Enforcer *casbin.Enforcer

func CasbinInit() error{
	adapter, err :=sqlxadapter.NewAdapter(db,"")
	if err != nil {
		zap.L().Error("casbin sqlxadapter.NewAdapter failed",zap.Error(err))
		return err
	}
	Enforcer, err =casbin.NewEnforcer("./conf/casbinModel.conf", adapter)
	if err != nil {
		zap.L().Error("casbin casbin.NewEnforcer failed",zap.Error(err))
		return err
	}
	err = Enforcer.LoadPolicy()
	if err != nil {
		zap.L().Error("casbin Enforcer.LoadPolicy() failed",zap.Error(err))
		return err
	}

	return nil
}
