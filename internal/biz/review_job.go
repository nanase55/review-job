package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

// Msg 定义kafka中消息格式
type Msg struct {
	Type     string `json:"type"`
	Database string `json:"databse"`
	Table    string `json:"table"`
	IsDdl    bool   `json:"isDdl"`
	Data     []map[string]any
}

type ReviewRepo interface {
	// 写入doc到es
	CreateDoc(context.Context, map[string]any) error
	// 更新es里的doc
	UpdateDoc(context.Context, map[string]any) error

	// 从mq读取消息
	FetchMessage(context.Context) (*kafka.Message, error)
	// 提交消息
	CommitMessage(context.Context, *kafka.Message) error
}

type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ReviewUsecase) Run(ctx context.Context) {
	uc.log.Info("开始消费 MQ 消息...")

	for {
		select {
		case <-ctx.Done():
			uc.log.Info("Context 取消，退出消费循环")
			return
		default:
			if err := uc.ConsumeAndSaveFromMQ(ctx); err != nil {
				uc.log.Warnf("消费MQ消息失败: %v", err)
				// 可以加一个短暂的延迟，避免错误时疯狂重试
				time.Sleep(time.Second)
			}
		}
	}
}

// 消费 MQ 消息并写入 ES
func (uc *ReviewUsecase) ConsumeAndSaveFromMQ(ctx context.Context) error {
	// 1. 从 mq 读取消息
	m, err := uc.repo.FetchMessage(ctx)
	if err != nil {
		return err
	}

	// 2. 反序列化消息
	msg := &Msg{}
	if err := json.Unmarshal(m.Value, msg); err != nil {
		return fmt.Errorf("json.Unmarshal failed %v", err)
	}

	// 补充！
	// 实际的业务场景可能需要在这增加一个步骤：对数据做业务处理
	// 例如：如消息转换、过滤、聚合

	// 3. 写入 ES
	switch msg.Type {
	case "INSERT":
		for idx := range msg.Data {
			if err := uc.repo.CreateDoc(ctx, msg.Data[idx]); err != nil {
				// 如果部分写入成功,因为id相同,有幂等性保证
				return fmt.Errorf("写入es失败: %#v", msg)
			}
		}
	case "UPDATE":
		for idx := range msg.Data {
			if err := uc.repo.UpdateDoc(ctx, msg.Data[idx]); err != nil {
				return fmt.Errorf("更新es失败: %#v", msg)
			}
		}
	default:
		return fmt.Errorf("unknown msg.Type")
	}

	// 提交消息
	if err := uc.repo.CommitMessage(ctx, m); err != nil {
		return fmt.Errorf("提交消息失败: %#v", msg)
	}
	return nil
}
