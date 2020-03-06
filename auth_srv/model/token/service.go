package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	tokenExpiredDate = time.Hour * 24 //	token过期时间
)

type Subject struct {
	Id string `json:"id"`
}

type tokenService struct {
}

func (s *tokenService) SetToken(subject *Subject) (tokenStr string, err error) {
	claims, err := s.createTokenClaims(subject)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("authsrv: 创建claims失败")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("create token fail, err: %v", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, tokenStr)
	if err != nil {
		return "", fmt.Errorf("save token to redis, err: %v", err)
	}

	return
}

func (s *tokenService) ClearToken(token string) (err error) {
	// 解析token字符串
	claims, err := s.parseToken(token)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}

	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		Id: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	return
}

func (s *tokenService) GetUserId(token string) (userId int64, err error) {
	claims, err := s.parseToken(token)
	if err != nil {
		err = fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
		return
	}

	// 检查token是否过期
	token2, err := s.getTokenToCache(&Subject{Id: claims.Subject})
	if err != nil || token2 != token {
		log.Errorf("authsrv: token过期, err: %v", err)
		return
	}

	userId, err = strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		log.Errorf("authsrv: 解析失败, err: %v", err)
		return
	}

	return
}

// 生成 Claims
func (s *tokenService) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.Id,
		IssuedAt:  now.Unix(),
		Issuer:    "shop",
		Subject:   subject.Id,
	}

	return
}

// 保存token到缓存
func (s *tokenService) saveTokenToCache(subject *Subject, val string) (err error) {
	if err = redisClient.Set(getRedisKey(subject.Id), val, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("set token to redis fail，err: %v", err)
	}

	return
}

// 获取token
func (s *tokenService) getTokenToCache(subject *Subject) (token string, err error) {
	token, err = redisClient.Get(getRedisKey(subject.Id)).Result()
	if err != nil {
		err = fmt.Errorf("get token to redis fail，err: %v", err)
		return
	}

	return
}

// delTokenFromCache 清空token
func (s *tokenService) delTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = redisClient.Del(getRedisKey(subject.Id)).Err(); err != nil {
		return fmt.Errorf("del token to cache, err:" + err.Error())
	}
	return
}

// parseToken 解析token
func (s *tokenService) parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})

	// jwt 框架自带了一些检测，如过期，发布者错误等
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
	}

	// 检测合法
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] 不合法的token")
	}

	return mapClaimToJwClaim(claims), nil
}

// 生成reids key
func getRedisKey(id string) string {
	return fmt.Sprintf("shop.user.token.%s", id)
}

// 把jwt的claim转成claims
func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}

	return jC
}
