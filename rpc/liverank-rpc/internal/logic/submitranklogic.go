package logic

import (
	"context"

	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/internal/svc"
	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/types/liverank"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitRankLogic {
	return &SubmitRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  客户端提交排名
func (l *SubmitRankLogic) SubmitRank(in *liverank.SubmitRankRequest) (*liverank.SubmitRankResponse, error) {
	// todo: add your logic here and delete this line

	return &liverank.SubmitRankResponse{}, nil
}
