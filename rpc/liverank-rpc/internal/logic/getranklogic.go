package logic

import (
	"context"

	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/internal/svc"
	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/types/liverank"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankLogic {
	return &GetRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  客户端获取排名
func (l *GetRankLogic) GetRank(in *liverank.GetRankRequest) (*liverank.GetRankResponse, error) {
	// todo: add your logic here and delete this line

	return &liverank.GetRankResponse{}, nil
}
