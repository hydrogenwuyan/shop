package controller

import (
	"encoding/json"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	log "github.com/sirupsen/logrus"
	"net/http"
	authsrvproto "project/shop/auth_srv/proto"
	basiccommon "project/shop/basic/common"
	inventorysrvproto "project/shop/inventory_srv/proto"
	ordersrvproto "project/shop/order_srv/proto"
)

var (
	ordersrvClient     ordersrvproto.OrderService         // 订单服务client
	authsrvClient      authsrvproto.AuthService           // token服务client
	inventorysrvClient inventorysrvproto.InventoryService // 库存服务client
)

// 初始化
func Init() {
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	ordersrvClient = ordersrvproto.NewOrderService("shop.order.srv", cl)
	inventorysrvClient = inventorysrvproto.NewInventoryService("shop.inventory.srv", cl)
	authsrvClient = authsrvproto.NewAuthService("shop.auth.srv", cl)
}

type ReqMsg struct {
	ShopId int64 `json:"shopId"`
	Num    int64 `json:"num"`
}

// 购买商品
func ShopBuy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	flag, errorCode, userId := checkoutToken(r)
	if !flag {
		http.Error(w, errorCode, 400)
	}

	reqMsg := &ReqMsg{}
	err := json.NewDecoder(r.Body).Decode(reqMsg)
	if err != nil {
		log.Errorf("orderweb: body json反序列化失败, error: %v", err)
		http.Error(w, "body json反序列化失败", 400)
		return
	}

	shopId := reqMsg.ShopId
	num := reqMsg.Num

	// 生成预订单
	scOrderCreate, err := ordersrvClient.CreateOrder(ctx, &ordersrvproto.CSOrderCreate{
		ShopId: shopId,
		UserId: userId,
		Num:    num,
	})
	if err != nil || scOrderCreate.Error.Code != 200 {
		log.WithFields(log.Fields{
			"userId": userId,
			"shopId": shopId,
			"num":    num,
			"detail": scOrderCreate.Error.Detail,
			"error":  err,
		}).Error("orderweb: 创建订单失败")
		http.Error(w, "创建订单失败", 400)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"code":  200,
		"order": scOrderCreate.Info,
	}

	// 返回JSON结构
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithFields(log.Fields{
			"userId": userId,
			"error":  err,
		}).Error("orderweb: json编码失败")
		http.Error(w, err.Error(), 500)
		return
	}
}

// 检查token
func checkoutToken(r *http.Request) (flag bool, errorCode string, uid int64) {
	flag = false
	// 验证post方法
	if r.Method != "POST" {
		log.WithFields(log.Fields{
			"method": r.Method,
		}).Error("orderweb: 不是post方法")
		errorCode = "不是post方法"
		return
	}

	// 获取cookie
	cookie, err := r.Cookie(basiccommon.UserCookieName)
	if err != nil {
		log.Errorf("orderweb: 获取cookie失败, err: %v", err)
		errorCode = "获取cookie失败"
		return
	}

	// 验证token
	scUserIdGet, err := authsrvClient.GetUserIdByToken(r.Context(), &authsrvproto.CSUserIdGet{
		Token: cookie.Value,
	})
	if err != nil || scUserIdGet.Error.Code != 200 {
		log.Errorf("orderweb: 用户未登录")
		errorCode = "用户未登录"
		return
	}

	uid = scUserIdGet.UserId
	flag = true
	return
}
