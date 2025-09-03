package data

import (
	"context"
	"fmt"
	"review-job/internal/biz"
	"strconv"

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

// CreateDoc 创建reviewInfos文档
func (r *reviewRepo) CreateReviewInfo(ctx context.Context, table string, doc map[string]any) error {
	var idx string
	var id int64
	var d any

	switch table {
	case "review_infos":
		d = MapToReviewDoc(doc)
		idx = r.data.esClient.ReviewInfosIdx
		id = d.(*ReviewInfoDoc).ReviewID
	case "review_reply_info":
		d = MapToReviewReplyDoc(doc)
		idx = r.data.esClient.ReviewReplyIdx
		id = d.(*ReviewReplyDoc).ReplyID
	case "review_appeal_info":
		d = MapToReviewAppealDoc(doc)
		idx = r.data.esClient.ReviewAppealIdx
		id = d.(*ReviewAppealDoc).AppealID
	default:
		return fmt.Errorf("unknown table: %s", table)
	}

	// 添加文档: id相同保证幂等性
	resp, err := r.data.esClient.Index(idx).Id(strconv.Itoa(int(id))).Document(d).Do(ctx)
	if err != nil {
		r.log.Errorf("indexing document failed, err:%v\n", err)
		return err
	}

	r.log.Debugf("result:%#v\n", resp.Result)
	return nil
}

// UpdateDoc 在es中更新文档
func (r *reviewRepo) UpdateReviewInfo(ctx context.Context, table string, doc map[string]any) error {
	var idx string
	var id int64
	var d any

	switch table {
	case "review_infos":
		d = MapToReviewDoc(doc)
		idx = r.data.esClient.ReviewInfosIdx
		id = d.(*ReviewInfoDoc).ReviewID
	case "review_reply_info":
		d = MapToReviewReplyDoc(doc)
		idx = r.data.esClient.ReviewReplyIdx
		id = d.(*ReviewReplyDoc).ReplyID
	case "review_appeal_info":
		d = MapToReviewAppealDoc(doc)
		idx = r.data.esClient.ReviewAppealIdx
		id = d.(*ReviewAppealDoc).AppealID
	default:
		return fmt.Errorf("unknown table: %s", table)
	}

	resp, err := r.data.esClient.Update(idx, strconv.Itoa(int(id))).Doc(d).Do(ctx)
	if err != nil {
		r.log.Errorf("update document failed, err:%v\n", err)
		return err
	}

	r.log.Debugf("result:%v\n", resp.Result)
	return nil
}
