package define

import (
	"github.com/golang-jwt/jwt/v4"
)

// UserClaim 用户加密声明
type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

// JwtKey 再次加密token
var JwtKey = "cloud_disk_key"

// EmailPassword 邮箱密码
var EmailPassword = "MIRWCVJMBVSVUCZN"

// EmailAccount 邮箱账号
var EmailAccount = "at911477183@163.com"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间
var CodeExpire = 300

//SECRETID 表示用户的 SecretId
var SECRETID = "AKIDue9OO9eutd7LK3XlkTxCBvPdcMq54nHX"

//SECRETKEY 表示用户的 SecretKey
var SECRETKEY = "6xZv1sUqxoQnumOlsjKgneFbiZAWcS2N"
var CosBucket = "https://cloud-disk-1312527462.cos.ap-nanjing.myqcloud.com"

// PageSize 分页的默认页数
var PageSize = 20

//时间格式
var TimeFormat = "2006-01-02 15:04:05"
