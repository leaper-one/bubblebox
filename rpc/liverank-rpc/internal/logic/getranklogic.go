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

	resp, err := l.svcCtx.Model.FindOne(context.Background(), in.Roomid)

	if err != nil {
		logx.Info("Has err")
		logx.Info(err)
		return nil, err
	}

	logx.Info(resp)

	return nil, err

	//resp, err := l.svcCtx.Model.GetRank(context.Background(), 123, 0)
	//
	////logx.Info(
	//
	//if err != nil {
	//	println("has err")
	//	logx.Info(err)
	//	return &liverank.GetRankResponse{}, err
	//}
	//
	//var (
	//	ranks     []*liverank.Rank
	//	timeStamp int64 = resp[0].Timestamp
	//)
	//
	//for k, v := range resp {
	//	ranks[k].Buid = v.Buid
	//	ranks[k].Rank = v.Rank
	//
	//	if v.Timestamp > timeStamp {
	//		timeStamp = v.Timestamp
	//	}
	//}
	//
	//logx.Info(ranks)
	//
	//println("获取排名成功")
	//return &liverank.GetRankResponse{
	//	Code:      200,
	//	Message:   "success",
	//	Timestamp: timeStamp,
	//	Ranks:     ranks,
	//}, nil

	//return &liverank.GetRankResponse{}, nil
}
