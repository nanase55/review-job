package data

import (
	"context"
	"review-job/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateDoc 在es中创建文档
func (r *reviewRepo) CreateDoc(ctx context.Context, doc map[string]any) error {
	reviewID := doc["review_id"].(string)
	// 添加文档
	resp, err := r.data.esClient.Index(r.data.esClient.Idx).
		Id(reviewID). // id相同保证幂等性
		Document(doc).
		Do(ctx)
	if err != nil {
		r.log.Errorf("indexing document failed, err:%v\n", err)
		return err
	}
	r.log.Debugf("result:%#v\n", resp.Result)
	return nil
}

// UpdateDoc 在es中更新文档
func (r *reviewRepo) UpdateDoc(ctx context.Context, d map[string]any) error {
	reviewID := d["review_id"].(string)
	resp, err := r.data.esClient.Update(r.data.esClient.Idx, reviewID).
		Doc(d). // 使用结构体变量更新
		Do(ctx)
	if err != nil {
		r.log.Errorf("update document failed, err:%v\n", err)
		return err
	}
	r.log.Debugf("result:%v\n", resp.Result)
	return nil
}

// FetchMessage 从mq里读取消息
func (r *reviewRepo) FetchMessage(ctx context.Context) (*kafka.Message, error) {
	r.log.Debug("阻塞读取kafka消息")
	msg, err := r.data.kafkaReader.FetchMessage(ctx)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

// CommitMessage 提交消息
func (r *reviewRepo) CommitMessage(ctx context.Context, m *kafka.Message) error {
	return r.data.kafkaReader.CommitMessages(ctx, *m)
}
