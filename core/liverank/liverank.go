package liverrank

type (
	BiliRank struct {
		//gorm.Model
		ID        int64 `gorm:"primary_key"`
		Timestamp int64 `gorm:"default:0"`
		Buid      int64 `gorm:"default:0"`
		RoomId    int64 `gorm:"default:0"`
		Rank      int64 `gorm:"default:0"`
		GiftValue int64 `gorm:"default:0"`
		IsConcern bool  `gorm:"default:false"`
	}
)
