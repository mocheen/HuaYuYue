package jwt

import (
	"gateway/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(userID int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "38384-SearchEngine",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtSecret = []byte(conf.Conf.Server.JwtSecret)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	var jwtSecret = []byte(conf.Conf.Server.JwtSecret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
