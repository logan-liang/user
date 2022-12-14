package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestALogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestALogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestALogic {
	return &TestALogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestALogic) TestA(in *rpc.TestReq) (*rpc.TestResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.TestResp{}, nil
}
