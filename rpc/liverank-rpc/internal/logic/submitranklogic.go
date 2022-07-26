package logic

import (
	"context"
	liverankmodel "github.com/leaper-one/bubblebox/model/liverank"
	"time"

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

	_, err := l.svcCtx.Model.Insert(context.Background(), &liverankmodel.BiliRanks{
		Timestamp: time.Now().Unix(),
		Buid:      in.Buid,
		RoomId:    in.Roomid,
		Rank:      in.Rank,
		GiftValue: in.Giftvalue,
		IsConcern: boolToInt64(in.Isconcern),
	})

	if err != nil {
		logx.Info(err)
		return &liverank.SubmitRankResponse{}, err
	}
	return &liverank.SubmitRankResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}

	return 0
}
