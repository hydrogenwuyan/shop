# shop

基于micor的购物系统

## 启动顺序

| 优先级 | 服务名称 | 
| -------- | -------- |
| 0      | config_srv     | 
| 1      | auth_srv     | 
| 2      | user inventory order  |

## 需求整理
1.用户登陆
2.用户购买商品，生成订单
3.用户确认订单，付款成功

## 相关组件
mysql etcd redis