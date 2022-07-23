package logic

import (
	"context"

	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/internal/svc"
	"github.com/leaper-one/bubblebox/rpc/liverank-rpc/types/liverank"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkConcernLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkConcernLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkConcernLogic {
	return &MarkConcernLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  关注某个直播间的某个已有投稿的用户
func (l *MarkConcernLogic) MarkConcern(in *liverank.MarkConcernRequest) (*liverank.MarkConcernResponse, error) {
	// todo: add your logic here and delete this line

	return &liverank.MarkConcernResponse{}, nil
}
