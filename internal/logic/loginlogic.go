package logic

import (
	"context"
	"errors"
	"ws_demo_rpc/internal/entity"

	"ws_demo_rpc/internal/svc"
	"ws_demo_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	u := &entity.User{}

	tx := l.svcCtx.DB.Where(&entity.User{
		Account: in.Account,
	}).First(u)

	if tx.Error != nil {
		return nil, errors.New("查询用户失败" + tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}

	return &user.LoginResponse{
		Name:   u.Name,
		Id:     u.Id,
		Avatar: u.Avatar,
	}, nil
}
