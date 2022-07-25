package svc

import (
	"github.com/leaper-one/bubblebox/model/liverank"
	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	//"github.com/leaper-one/bubblebox/rpc/liverank-rpc/liverank"
)

type ServiceContext struct {
	Config config.Config
	Model  liverank.BiliRanksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  liverank.NewBiliRanksModel(sqlx.NewMysql(c.DataSource)),
	}
}
