package logic

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"
	"MiniDY/app/videos/model"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishVideoReq, file multipart.File, fileName string) (*types.PublishVideoResp, error) {
	// todo: add your logic here and delete this line
	//从上下文ctx拿到UserID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 生成一个唯一的文件名
	timestamp := time.Now()
	fileName += fmt.Sprintf("%v-%v", req.Title, timestamp)

	//文件保存路径 暂定为本地
	savePath := filepath.Join("VideoData", fileName)
	// 创建保存文件的目录
	err := os.MkdirAll("VideoData", os.ModePerm)
	if err != nil {
		return nil, err
	}
	// 创建目标文件
	destinationFile, err := os.Create(savePath)
	if err != nil {
		return nil, err
	}
	defer destinationFile.Close()

	// 将上传的文件内容拷贝到目标文件
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		return &types.PublishVideoResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		}, nil
	}
	//将视频信息保存到数据库中
	//将数据库的事务处理封装成一个函数调用，通过将 fn 函数作为参数传递给 Trans 方法，可以在 fn 函数中执行一系列的数据库操作，这些操作将在一个事务中执行，并且出现任何错误时，事务将自动回滚。
	if err := l.svcCtx.VideoModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//创建新的视频实例对象
		video := new(model.Video)
		if req.Title == "" {
			return errors.New("视频标题不能为空")
		}
		video.Title = req.Title
		video.AuthorId = userId
		video.PlayUrl = savePath

		insertResult, err := l.svcCtx.VideoModel.Insert(l.ctx, video)
		if err != nil {
			return errors.Wrapf(dyerr.ErrDBerror, "new video Insert err:%v,video:%+v", err, video)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(dyerr.ErrDBerror, "new video insertResult.LastInsertId err:%v,user:%+v", err, video)
		}
		video.Id = lastId
		return nil

	}); err != nil {
		return &types.PublishVideoResp{
			StatusCode: 1,
			StatusMsg:  "插入视频数据失败" + err.Error(),
		}, nil
	}

	return &types.PublishVideoResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "投稿视频" + dyerr.SUCCESS,
	}, nil
}
