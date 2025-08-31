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

type ReviewInfo struct {
	ReviewID      string    `json:"review_id"`
	OrderID       string    `json:"order_id"`
	SkuID         string    `json:"sku_id"`
	SpuID         string    `json:"spu_id"`
	StoreID       string    `json:"store_id"`
	UserID        string    `json:"user_id"`
	Content       string    `json:"content"`
	Tags          string    `json:"tags,omitempty"`
	Score         int       `json:"score"`
	ServiceScore  int       `json:"service_score"`
	ExpressScore  int       `json:"express_score"`
	HasMedia      bool      `json:"has_media"`
	PicInfo       string    `json:"pic_info,omitempty"`
	VideoInfo     string    `json:"video_info,omitempty"`
	Anonymous     bool      `json:"anonymous"`
	Status        int       `json:"status"`
	IsDefault     bool      `json:"is_default"`
	HasReply      bool      `json:"has_reply"`
	OpInfo        OpInfo    `json:"op_info,omitempty"`
	GoodsSnapshot string    `json:"goods_snapshot,omitempty"`
	ExtJSON       string    `json:"ext_json,omitempty"`
	CtrlJSON      string    `json:"ctrl_json,omitempty"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
	DeleteAt      time.Time `json:"delete_at,omitempty"`
}

// MapToReviewDoc 转换为ReviewInfo
func MapToReviewDoc(m map[string]any) *ReviewInfo {
	doc := &ReviewInfo{}
	if v, ok := m["review_id"].(string); ok {
		doc.ReviewID = v
	}
	if v, ok := m["order_id"].(string); ok {
		doc.OrderID = v
	}
	if v, ok := m["sku_id"].(string); ok {
		doc.SkuID = v
	}
	if v, ok := m["spu_id"].(string); ok {
		doc.SpuID = v
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
		doc.Score, _ = strconv.Atoi(v)
	}
	if v, ok := m["service_score"].(string); ok {
		doc.ServiceScore, _ = strconv.Atoi(v)
	}
	if v, ok := m["express_score"].(string); ok {
		doc.ExpressScore, _ = strconv.Atoi(v)
	}
	if v, ok := m["has_media"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.HasMedia = (a == 1)
	}
	if v, ok := m["pic_info"].(string); ok {
		doc.PicInfo = v
	}
	if v, ok := m["video_info"].(string); ok {
		doc.VideoInfo = v
	}
	if v, ok := m["anonymous"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.Anonymous = (a == 1)
	}
	if v, ok := m["status"].(string); ok {
		doc.Status, _ = strconv.Atoi(v)
	}
	if v, ok := m["is_default"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.IsDefault = (a == 1)
	}
	if v, ok := m["has_reply"].(string); ok {
		a, _ := strconv.Atoi(v)
		doc.HasReply = (a == 1)
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
