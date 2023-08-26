package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"
	"MiniDY/app/interaction/cmd/rpc/interactsservice"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCommentLogic {
	return &PostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostCommentLogic) PostComment(req *types.CommentActionReq) (*types.CommentActionResp, error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)

	rpcResp, err := l.svcCtx.InteractsRpc.VideoCommentAction(l.ctx, &interactsservice.DouyinCommentActionRequest{
		UserId:      userId,
		VideoId:     req.VideoId,
		CommentText: &req.CommentText,
		CommentId:   &req.CommentId,
	})
	if err != nil {
		return &types.CommentActionResp{
			StatusCode: 1,
			StatusMsg:  "评论失败" + dyerr.SUCCESS,
			Comment:    nil,
		}, nil
	}
	commentResp, err := l.svcCtx.CommentModel.FindOne(l.ctx, rpcResp.Comment)
	if err != nil {
		return &types.CommentActionResp{
			StatusCode: 1,
			StatusMsg:  "评论失败" + err.Error(),
			Comment:    nil,
		}, nil
	}
	userInfo, _ := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
		UserId: commentResp.UserId,
	})
	var tempuser types.User
	_ = copier.Copy(&tempuser, userInfo.User)

	var resp *types.Comment
	resp.Id = commentResp.Id
	resp.UserInfo = tempuser
	resp.Content = commentResp.Content
	resp.Create_data = commentResp.CreateTime.Format("mm-dd")

	if req.ActionType == 1 {
		return &types.CommentActionResp{
			StatusCode: int32(dyerr.OK),
			StatusMsg:  "评论" + dyerr.SUCCESS,
			Comment:    resp,
		}, nil
	} else {
		return nil, nil
	}
}
