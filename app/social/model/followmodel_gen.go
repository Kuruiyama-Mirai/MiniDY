// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	followFieldNames          = builder.RawFieldNames(&Follow{})
	followRows                = strings.Join(followFieldNames, ",")
	followRowsExpectAutoSet   = strings.Join(stringx.Remove(followFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followRowsWithPlaceHolder = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFollowIdPrefix       = "cache:follow:id:"
	cacheFollowToUserIdPrefix = "cache:follow:toUserId:"
	cacheFollowUserIdPrefix   = "cache:follow:userId:"
)

type (
	followModel interface {
		Insert(ctx context.Context, data *Follow) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follow, error)
		FindOneByToUserId(ctx context.Context, toUserId int64) (*Follow, error)
		FindOneByUserId(ctx context.Context, userId int64) (*Follow, error)
		Update(ctx context.Context, data *Follow) error
		Delete(ctx context.Context, id int64) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindOneByUserIdAndFollowId(ctx context.Context, userId, followId int64)(*Follow, error)
		FindAllFollow(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string)([]int64, error)
		FindAllFollower(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string)([]int64, error)
		FindAllFriends(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string)([]int64, error)
	}

	defaultFollowModel struct {
		sqlc.CachedConn
		table string
	}

	Follow struct {
		Id         int64        `db:"id"`          // 关注表的主键id
		UserId     int64        `db:"user_id"`     // 关注人的id
		ToUserId   int64        `db:"to_user_id"`  // 被关注人id
		Status     int64        `db:"status"`      // 是否为互相关注 1为相互关注 0不是
		CreateTime time.Time    `db:"create_time"` // 创建时间
		UpdateTime time.Time    `db:"update_time"` // 更新时间
		DeletedAt  sql.NullTime `db:"deleted_at"`  // 逻辑删除
	}
)

func newFollowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFollowModel {
	return &defaultFollowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`follow`",
	}
}

func (m *defaultFollowModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, id)
	followToUserIdKey := fmt.Sprintf("%s%v", cacheFollowToUserIdPrefix, data.ToUserId)
	followUserIdKey := fmt.Sprintf("%s%v", cacheFollowUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, followIdKey, followToUserIdKey, followUserIdKey)
	return err
}

func (m *defaultFollowModel) FindOne(ctx context.Context, id int64) (*Follow, error) {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, id)
	var resp Follow
	err := m.QueryRowCtx(ctx, &resp, followIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) FindOneByToUserId(ctx context.Context, toUserId int64) (*Follow, error) {
	followToUserIdKey := fmt.Sprintf("%s%v", cacheFollowToUserIdPrefix, toUserId)
	var resp Follow
	err := m.QueryRowIndexCtx(ctx, &resp, followToUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `to_user_id` = ? limit 1", followRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, toUserId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) FindOneByUserId(ctx context.Context, userId int64) (*Follow, error) {
	followUserIdKey := fmt.Sprintf("%s%v", cacheFollowUserIdPrefix, userId)
	var resp Follow
	err := m.QueryRowIndexCtx(ctx, &resp, followUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", followRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) Insert(ctx context.Context, data *Follow) (sql.Result, error) {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, data.Id)
	followToUserIdKey := fmt.Sprintf("%s%v", cacheFollowToUserIdPrefix, data.ToUserId)
	followUserIdKey := fmt.Sprintf("%s%v", cacheFollowUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, followRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.ToUserId, data.Status, data.DeletedAt)
	}, followIdKey, followToUserIdKey, followUserIdKey)
	return ret, err
}

func (m *defaultFollowModel) Update(ctx context.Context, newData *Follow) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, data.Id)
	followToUserIdKey := fmt.Sprintf("%s%v", cacheFollowToUserIdPrefix, data.ToUserId)
	followUserIdKey := fmt.Sprintf("%s%v", cacheFollowUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.ToUserId, newData.Status, newData.DeletedAt, newData.Id)
	}, followIdKey, followToUserIdKey, followUserIdKey)
	return err
}

func (m *defaultFollowModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFollowIdPrefix, primary)
}

func (m *defaultFollowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowModel) tableName() string {
	return m.table
}

func(m *defaultFollowModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error{
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s )
	})
}

func (m *defaultFollowModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func(m* defaultFollowModel) FindOneByUserIdAndFollowId(ctx context.Context, userId, followId int64)(*Follow ,error){
	followToUserIdPrefix := fmt.Sprintf("%s%v",cacheFollowUserIdPrefix,userId)
	var resp Follow
	//根据用户索引来查
	err := m.QueryRowIndexCtx(ctx,&resp,followToUserIdPrefix,m.formatPrimary,func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `to_user_id` = ? ",followRows,m.table)
		if err := conn.QueryRowCtx(ctx,&resp, query,userId,followId); err != nil{
			return nil, err
		}
		return resp.Id, nil
	},m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
//查找所有关注者
func (m *defaultFollowModel)FindAllFollow(ctx context.Context, builder squirrel.SelectBuilder, orderBy string)([]int64, error){

	builder = builder.Columns(followRows)

	if orderBy == ""{
		builder = builder.OrderBy("id DESC")
	}else{
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Follow
	var followIdList []int64	
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	for i := range resp{
		followIdList = append(followIdList, resp[i].ToUserId)
	}
	switch err {
	case nil:
		return followIdList, nil
	default:
		return nil, err
	}

}
//查找所有粉丝
func (m *defaultFollowModel)FindAllFollower(ctx context.Context, builder squirrel.SelectBuilder, orderBy string)([]int64, error){

	builder = builder.Columns(followRows)

	if orderBy == ""{
		builder = builder.OrderBy("id DESC")
	}else{
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Follow
	var followerIdList []int64	
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	for i := range resp{
		followerIdList = append(followerIdList, resp[i].UserId)
	}
	switch err {
	case nil:
		return followerIdList, nil
	default:
		return nil, err
	}

}

//查找所有好友
func (m *defaultFollowModel)FindAllFriends(ctx context.Context, builder squirrel.SelectBuilder, orderBy string)([]int64, error){

	builder = builder.Columns(followRows)

	if orderBy == ""{
		builder = builder.OrderBy("id DESC")
	}else{
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("status = ?",1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Follow
	var friendsIdList []int64	
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	for i := range resp{
		friendsIdList = append(friendsIdList, resp[i].ToUserId)
	}
	switch err {
	case nil:
		return friendsIdList, nil
	default:
		return nil, err
	}

}