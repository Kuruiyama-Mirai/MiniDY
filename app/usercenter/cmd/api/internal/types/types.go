// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id               int64  `json:"id"`
	Username         string `json:"username"`
	Follow_count     int64  `json:"follow_count"`
	Follower_count   int64  `json:"follower_count"`
	Is_follow        bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	Total_favorited  int64  `json:"total_favorited"`
	Work_count       int64  `json:"work_count"`
	Favorite_count   int64  `json:"favorite_count"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type GetUserInfoReq struct {
	UserID int64 `json:"user_id"`
}

type GetUserInfoResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserInfo   User   `json:"userInfo"`
}
