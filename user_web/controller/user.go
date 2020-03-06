package controller

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	authsrvproto "project/shop/auth_srv/proto"
	basiccommon "project/shop/basic/common"
	usersrvproto "project/shop/user_srv/proto"
	"time"
)

var (
	usersrvClient usersrvproto.UserService // 用户服务client
	authsrvClient authsrvproto.AuthService // token服务client
)

// 初始化
func Init() {
	cl := client.DefaultClient
	usersrvClient = usersrvproto.NewUserService("shop.user.srv", cl)
	authsrvClient = authsrvproto.NewAuthService("shop.auth.srv", cl)
}

type ReqMsg struct {
	Id  int64  `json:"id"`
	Pwd string `json:"pwd"`
}

// 用户登陆
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 验证post方法
	if r.Method != "POST" {
		log.WithFields(log.Fields{
			"method": r.Method,
		}).Error("userweb: 不是post方法")
		http.Error(w, "不是post方法", 400)
		return
	}

	fmt.Println("***********", r.RemoteAddr)

	reqMsg := &ReqMsg{}
	err := json.NewDecoder(r.Body).Decode(reqMsg)
	if err != nil {
		log.Fatalf("userweb: body json反序列化失败, error: %v", err)
		http.Error(w, "body json反序列化失败", 400)
		return
	}

	userId := reqMsg.Id
	pwd := reqMsg.Pwd

	// 验证用户是否存在
	scUserInfo, err := usersrvClient.QueryUserInfoByUserId(ctx, &usersrvproto.CSUserInfo{
		UserId:   userId,
		Password: pwd,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"userId": userId,
			"error":  err,
		}).Error("userweb: rpc请求失败")
		http.Error(w, "rpc请求失败", 400)
		return
	}

	if scUserInfo.Error.Code != 200 {
		log.WithFields(log.Fields{
			"userId": userId,
			"pwd":    pwd,
		}).Warn("userweb: 用户不存在")
		http.Error(w, "用户不存在", 400)
		return
	}

	// 设置token
	scTokenSet, err := authsrvClient.SetTokenByUserId(ctx, &authsrvproto.CSTokenSet{
		UserId: userId,
	})
	if err != nil || scTokenSet.Error.Code != 200 {
		log.WithFields(log.Fields{
			"userId": userId,
			"error":  err,
		}).Warn("userweb: 设置token失败")
		http.Error(w, "设置token失败", 400)
		return
	}

	// 设置cookie
	w.Header().Add("set-cookie", "application/json; charset=utf-8")
	expire := time.Now().Add(time.Hour) // 过期时间
	http.SetCookie(w, &http.Cookie{
		Name:    basiccommon.UserCookieName,
		Value:   scTokenSet.Token,
		Path:    "/",
		Expires: expire,
		MaxAge:  90000,
	})

	// 返回结果
	response := map[string]interface{}{
		"code":   200,
		"userId": scTokenSet.UserId,
		"token":  scTokenSet.Token,
	}

	// 返回JSON结构
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithFields(log.Fields{
			"userId": userId,
			"error":  err,
		}).Error("userweb: json编码失败")
		http.Error(w, err.Error(), 500)
		return
	}
}

// 用户登出
func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 验证post方法
	if r.Method != "POST" {
		log.WithFields(log.Fields{
			"method": r.Method,
		}).Error("userweb: 不是post方法")
		http.Error(w, "不是post方法", 400)
		return
	}

	// 获取cookie
	cookie, err := r.Cookie(basiccommon.UserCookieName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("userweb: 获取cookie失败")
		http.Error(w, "获取cookie失败", 400)
		return
	}
	token := cookie.Value

	// 清理token
	scTokenClear, err := authsrvClient.ClearTokenByUserId(ctx, &authsrvproto.CSTokenClear{
		Token: token,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"token": token,
			"error": err,
		}).Error("userweb: rpc请求失败")
		http.Error(w, "rpc请求失败", 400)
		return
	}

	if scTokenClear.Error.Code != 200 {
		log.WithFields(log.Fields{
			"detail": scTokenClear.Error.Detail,
			"error":  err,
		}).Error("userweb: 清理token失败")
		http.Error(w, "清理token失败", 400)
		return
	}

	// 清理cookie
	http.SetCookie(w, &http.Cookie{
		Name:    basiccommon.UserCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(0 * time.Second),
		MaxAge:  0,
	})

	// 返回结果
	response := map[string]interface{}{
		"code": 200,
	}

	// 返回JSON结构
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("userweb: json编码失败")
		http.Error(w, err.Error(), 500)
		return
	}
}
