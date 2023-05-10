package system

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model/request"
	"github.com/lsy88/jsonwizard/model/system"
	"github.com/lsy88/jsonwizard/utils"
	"strconv"
)

type JwtService struct{}

//拉黑jwt
func (j *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.JW_DB.DB.Create(&jwtList).Error
	return
}

//判断jwt是否在黑名单内部
func (j *JwtService) IsBlacklist(jwt string) bool {
	var jwtBlack system.JwtBlacklist
	global.JW_DB.DB.Where("jwt = ?", jwt).Limit(1).Find(&jwtBlack)
	if jwtBlack != (system.JwtBlacklist{}) {
		return true
	}
	return false
}

//从redis中获取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.JW_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	duration, _ := utils.ParseDuration(global.JW_CONFIG.JWT.ExpiresTime)
	timer := duration
	err = global.JW_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//@function: SetRedisUserInfo
//@description: 向redis设置用户信息
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisUserInfo(userInfo request.UserCache) (err error) {
	_, err = global.JW_REDIS.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		rdb.HSet(context.Background(), strconv.Itoa(int(userInfo.ID)), "id", userInfo.ID)
		//roleJson, _ := json.Marshal(userInfo.RoleIds)
		rdb.HSet(context.Background(), strconv.Itoa(int(userInfo.ID)), "type", userInfo.Type)
		return nil
	})
	return err
}

//从redis获取用户信息
func (jwtService *JwtService) GetRedisUserInfo(id int) (userInfo request.UserCache, err error) {
	//var userInfoRedis request.UserCacheRedis
	err = global.JW_REDIS.HGetAll(context.Background(), strconv.Itoa(id)).Scan(&userInfo)
	if err != nil || userInfo.ID == 0 {
		return userInfo, errors.New(err.Error() + "查询用户缓存失败")
	}
	return
}

//删除redis用户信息
func (jwtService *JwtService) DelRedisUserInfo(id int) (err error) {
	return global.JW_REDIS.HDel(context.Background(), strconv.Itoa(id)).Err()
}
