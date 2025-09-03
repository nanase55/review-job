package data

import (
	"strconv"
	"time"
)

type OpInfo struct {
	Reason  string `json:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	OpUser  string `json:"op_user,omitempty"`
}

type ReviewInfoDoc struct {
	ReviewID      int64     `json:"review_id"`
	OrderID       int64     `json:"order_id"`
	SkuID         string    `json:"sku_id"`
	SpuID         int64     `json:"spu_id"`
	StoreID       string    `json:"store_id"`
	UserID        string    `json:"user_id"`
	Content       string    `json:"content"`
	Tags          string    `json:"tags,omitempty"`
	Score         int32     `json:"score"`
	ServiceScore  int32     `json:"service_score"`
	ExpressScore  int32     `json:"express_score"`
	HasMedia      int32     `json:"has_media"`
	PicInfo       string    `json:"pic_info,omitempty"`
	VideoInfo     string    `json:"video_info,omitempty"`
	Anonymous     int32     `json:"anonymous"`
	Status        int32     `json:"status"`
	IsDefault     int32     `json:"is_default"`
	HasReply      int32     `json:"has_reply"`
	OpInfo        OpInfo    `json:"op_info,omitempty"`
	GoodsSnapshot string    `json:"goods_snapshot,omitempty"`
	ExtJSON       string    `json:"ext_json,omitempty"`
	CtrlJSON      string    `json:"ctrl_json,omitempty"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
	DeleteAt      time.Time `json:"delete_at,omitempty"`
}

// MapToReviewDoc 转换为ReviewInfo
func MapToReviewDoc(m map[string]any) *ReviewInfoDoc {
	doc := &ReviewInfoDoc{}
	if v, ok := m["review_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ReviewID = int64(a)
	}
	if v, ok := m["order_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.OrderID = int64(a)
	}
	if v, ok := m["sku_id"].(string); ok {
		doc.SkuID = v
	}
	if v, ok := m["spu_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.SpuID = int64(a)
	}
	if v, ok := m["store_id"].(string); ok {
		doc.StoreID = v
	}
	if v, ok := m["user_id"].(string); ok {
		doc.UserID = v
	}
	if v, ok := m["content"].(string); ok {
		doc.Content = v
	}
	if v, ok := m["tags"].(string); ok {
		doc.Tags = v
	}
	if v, ok := m["score"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.Score = int32(a)
	}
	if v, ok := m["service_score"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ServiceScore = int32(a)
	}
	if v, ok := m["express_score"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ExpressScore = int32(a)
	}
	if v, ok := m["has_media"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.HasMedia = int32(a)
	}
	if v, ok := m["pic_info"].(string); ok {
		doc.PicInfo = v
	}
	if v, ok := m["video_info"].(string); ok {
		doc.VideoInfo = v
	}
	if v, ok := m["anonymous"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.Anonymous = int32(a)
	}
	if v, ok := m["status"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.Status = int32(a)
	}
	if v, ok := m["is_default"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.IsDefault = int32(a)
	}
	if v, ok := m["has_reply"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.HasReply = int32(a)
	}
	if v, ok := m["reason"].(string); ok {
		doc.OpInfo.Reason = v
	}
	if v, ok := m["remarks"].(string); ok {
		doc.OpInfo.Remarks = v
	}
	if v, ok := m["op_user"].(string); ok {
		doc.OpInfo.OpUser = v
	}
	if v, ok := m["goods_snapshot"].(string); ok {
		doc.GoodsSnapshot = v
	}
	if v, ok := m["ext_json"].(string); ok {
		doc.ExtJSON = v
	}
	if v, ok := m["ctrl_json"].(string); ok {
		doc.CtrlJSON = v
	}
	if v, ok := m["create_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.CreateAt = t
	}
	if v, ok := m["update_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	if v, ok := m["delete_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	return doc
}

type ReviewReplyDoc struct {
	ReplyID  int64  `json:"reply_id"`
	ReviewID int64  `json:"review_id"`
	StoreID  string `json:"store_id"`

	Content   string `json:"content,omitempty"`
	PicInfo   string `json:"pic_info,omitempty"`
	VideoInfo string `json:"video_info,omitempty"`

	CreateBy string `json:"create_by,omitempty"`
	UpdateBy string `json:"update_by,omitempty"`

	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
	DeleteAt time.Time `json:"delete_at,omitempty"`

	Version int `json:"version,omitempty"`

	ExtJSON  string `json:"ext_json,omitempty"`
	CtrlJSON string `json:"ctrl_json,omitempty"`
}

func MapToReviewReplyDoc(m map[string]any) *ReviewReplyDoc {
	doc := &ReviewReplyDoc{}

	if v, ok := m["reply_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ReplyID = int64(a)
	}
	if v, ok := m["review_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ReviewID = int64(a)
	}
	if v, ok := m["store_id"].(string); ok {
		doc.StoreID = v
	}
	if v, ok := m["content"].(string); ok {
		doc.Content = v
	}
	if v, ok := m["pic_info"].(string); ok {
		doc.PicInfo = v
	}
	if v, ok := m["video_info"].(string); ok {
		doc.VideoInfo = v
	}
	if v, ok := m["create_by"].(string); ok {
		doc.CreateBy = v
	}
	if v, ok := m["update_by"].(string); ok {
		doc.UpdateBy = v
	}
	if v, ok := m["create_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.CreateAt = t
	}
	if v, ok := m["update_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	if v, ok := m["delete_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	if v, ok := m["version"].(string); ok {
		doc.Version, _ = strconv.Atoi(v)
	}
	if v, ok := m["ext_json"].(string); ok {
		doc.ExtJSON = v
	}
	if v, ok := m["ctrl_json"].(string); ok {
		doc.CtrlJSON = v
	}
	return doc
}

type ReviewAppealDoc struct {
	AppealID  int64  `json:"appeal_id"`
	ReviewID  int64  `json:"review_id"`
	StoreID   string `json:"store_id"`
	Status    int32  `json:"status"`
	Content   string `json:"content,omitempty"`
	PicInfo   string `json:"pic_info,omitempty"`
	VideoInfo string `json:"video_info,omitempty"`
	OpInfo    OpInfo `json:"op_info,omitempty"`

	CreateBy string `json:"create_by,omitempty"`
	UpdateBy string `json:"update_by,omitempty"`

	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
	DeleteAt time.Time `json:"delete_at,omitempty"`

	Version int `json:"version,omitempty"`

	ExtJSON  string `json:"ext_json,omitempty"`
	CtrlJSON string `json:"ctrl_json,omitempty"`
}

func MapToReviewAppealDoc(m map[string]any) *ReviewAppealDoc {
	doc := &ReviewAppealDoc{}
	if v, ok := m["appeal_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.AppealID = int64(a)
	}
	if v, ok := m["review_id"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.ReviewID = int64(a)
	}
	if v, ok := m["store_id"].(string); ok {
		doc.StoreID = v
	}
	if v, ok := m["status"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.Status = int32(a)
	}
	if v, ok := m["content"].(string); ok {
		doc.Content = v
	}
	if v, ok := m["pic_info"].(string); ok {
		doc.PicInfo = v
	}
	if v, ok := m["video_info"].(string); ok {
		doc.VideoInfo = v
	}
	if v, ok := m["reason"].(string); ok {
		doc.OpInfo.Reason = v
	}
	if v, ok := m["remarks"].(string); ok {
		doc.OpInfo.Remarks = v
	}
	if v, ok := m["op_user"].(string); ok {
		doc.OpInfo.OpUser = v
	}
	if v, ok := m["video_info"].(string); ok {
		doc.VideoInfo = v
	}
	if v, ok := m["create_by"].(string); ok {
		doc.CreateBy = v
	}
	if v, ok := m["create_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.CreateAt = t
	}
	if v, ok := m["update_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	if v, ok := m["delete_at"].(string); ok {
		t, _ := time.Parse("2006-01-02 15:04:05", v)
		doc.UpdateAt = t
	}
	if v, ok := m["version"].(string); ok {
		doc.Version, _ = strconv.Atoi(v)
	}
	if v, ok := m["ext_json"].(string); ok {
		doc.ExtJSON = v
	}
	if v, ok := m["ctrl_json"].(string); ok {
		doc.CtrlJSON = v
	}
	return doc
}
