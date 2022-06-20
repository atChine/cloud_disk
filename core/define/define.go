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
var EmailPassword = "#######"

// EmailAccount 邮箱账号
var EmailAccount = "########"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间
var CodeExpire = 300

//SECRETID 表示用户的 SecretId
var SECRETID = "SecretId"

//SECRETKEY 表示用户的 SecretKey
var SECRETKEY = "SecretKey"
var CosBucket = "https://cloud-disk-#######.cos.ap-nanjing.myqcloud.com"

// PageSize 分页的默认页数
var PageSize = 20

// TimeFormat 时间格式
var TimeFormat = "2006-01-02 15:04:05"

// TokenExpire token过期时间
var TokenExpire = 3600

// RefreshTokenExpire 刷新token过期时间
var RefreshTokenExpire = 7200
