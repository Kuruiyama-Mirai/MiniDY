package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"
	"MiniDY/app/interaction/cmd/rpc/interactsservice"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (*types.CommentListResp, error) {
	// todo: add your logic here and delete this line
	//从ctx拿到用户Id
	//userId := ctxdata.GetUidFromCtx(l.ctx)

	rpcResp, err := l.svcCtx.InteractsRpc.VideoCommentList(l.ctx, &interactsservice.DouyinCommentListRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		return &types.CommentListResp{
			StatusCode:  1,
			StatusMsg:   "获取评论列表失败" + err.Error(),
			CommentList: nil,
		}, nil
	}
	var resp []*types.Comment
	if len(rpcResp.CommentList) > 0 {
		for _, commentId := range rpcResp.CommentList {
			commentResp, _ := l.svcCtx.CommentModel.FindOne(l.ctx, commentId)
			var tempComment types.Comment
			userInfo, _ := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: commentResp.UserId,
			})
			var tempuser types.User
			_ = copier.Copy(&tempuser, userInfo.User)
			tempComment.Id = commentId
			tempComment.UserInfo = tempuser
			tempComment.Content = commentResp.Content
			tempComment.Create_data = commentResp.CreateTime.Format("mm-dd")

			resp = append(resp, &tempComment)
		}
	}
	return &types.CommentListResp{
		StatusCode:  int32(dyerr.OK),
		StatusMsg:   "评论" + dyerr.SUCCESS,
		CommentList: resp,
	}, nil
}
