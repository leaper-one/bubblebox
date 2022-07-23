package liverank

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BiliRanksModel = (*customBiliRanksModel)(nil)

type (
	// BiliRanksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBiliRanksModel.
	BiliRanksModel interface {
		biliRanksModel
	}

	customBiliRanksModel struct {
		*defaultBiliRanksModel
	}
)

// NewBiliRanksModel returns a model for the database table.
func NewBiliRanksModel(conn sqlx.SqlConn, c cache.CacheConf) BiliRanksModel {
	return &customBiliRanksModel{
		defaultBiliRanksModel: newBiliRanksModel(conn, c),
	}
}
