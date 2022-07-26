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

	resp, err := l.svcCtx.Model.GetRank(context.Background(), in.Roomid, in.Timestamp)

	if err != nil {
		println("has err")
		logx.Info(err)
		return &liverank.GetRankResponse{}, err
	}

	var (
		ranks     []*liverank.Rank
		timeStamp = resp[0].Timestamp
	)

	for k, v := range resp {
		ranks = append(ranks, &liverank.Rank{
			Buid: resp[k].Buid,
			Rank: resp[k].Rank,
		})

		if v.Timestamp > timeStamp {
			timeStamp = v.Timestamp
		}
	}

	return &liverank.GetRankResponse{
		Code:      200,
		Message:   "success",
		Timestamp: timeStamp,
		Ranks:     ranks,
	}, nil

}
