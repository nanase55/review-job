package server

import (
	"context"
	"review-job/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type JobServer struct {
	uc     *biz.ReviewUsecase
	log    *log.Helper
	cancel context.CancelFunc
}

func NewJobServer(uc *biz.ReviewUsecase, logger log.Logger) *JobServer {
	return &JobServer{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// 实现kratos 的transport.server 接口
// app.run() 会启动一个goroutine 并阻塞等待
func (js *JobServer) Start(ctx context.Context) error {
	js.log.Info("Job Server start...")

	// 创建一个可控制的 context
	bizCtx, cancel := context.WithCancel(ctx)
	js.cancel = cancel // 保存 cancel 函数

	// 传入可控制的 context 给业务逻辑
	js.uc.Run(bizCtx)
	return nil
}

func (js *JobServer) Stop(context.Context) error {
	js.log.Info("job server stopping...")

	// 调用 cancel 取消业务逻辑
	if js.cancel != nil {
		js.cancel()
	}

	js.log.Info("job server stopped")
	return nil
}
