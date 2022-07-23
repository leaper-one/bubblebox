package liverrank

import (
	"gorm.io/gorm"
)

type (
	BiliRank struct {
		gorm.Model
		Buid      int64 `gorm:"default:0"`
		RoomId    int64 `gorm:"default:0"`
		Rank      int64 `gorm:"default:0"`
		GiftValue int64 `gorm:"default:0"`
		IsConcern bool  `gorm:"default:false"`
	}
)
