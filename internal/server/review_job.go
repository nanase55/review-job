package server

import (
	"context"
	"review-job/internal/biz"
)

type JobServer struct {
	uc   *biz.ReviewUsecase
	stop chan struct{}
}

func NewJobServer(uc *biz.ReviewUsecase) *JobServer {
	return &JobServer{
		uc:   uc,
		stop: make(chan struct{}),
	}
}

// 实现kratos 的transport.server 接口
func (js *JobServer) Start(ctx context.Context) error {
	return nil
}

func (js *JobServer) Stop(ctx context.Context) error {
	return nil
}
