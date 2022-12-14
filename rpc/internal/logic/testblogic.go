package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestBLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestBLogic {
	return &TestBLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestBLogic) TestB(in *rpc.TestReq) (*rpc.TestResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.TestResp{}, nil
}
