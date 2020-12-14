package data

import (
	"crypto/md5"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/rs/xid"
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
	"io"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        uint64
	Uuid      string
	Name      string
	Email     string
	CreatedAt int64
}

func CreateSession(user models.User) (*Session, error) {
	guid := xid.New()
	expireTime := time.Now().Add(10 * time.Hour).Unix()
	s := &Session{
		Id:        user.ID,
		Uuid:      guid.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: expireTime,
	}
	return s, nil

}

func CreateUser(name, email, password string) error {
	err := db.GetDB().Create(&models.User{Name: name, Email: email, Password: password}).Error
	return err
}

func UserIsExisted(email string) (user models.User, err error) {
	err = db.GetDB().Where("account = ?", email).First(&user).Error
	return
}

func SetHashValues(keyFiled []interface{}) error {
	conn := db.GetRedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("HMSET", keyFiled...)
	return err
}

func GetHashValues(keyFiled []interface{}) ([]string, error) {
	conn := db.GetRedisPool().Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HMGET", keyFiled...))
}

func SetKeyTtl(key string, expires int) error {
	conn := db.GetRedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, expires)
	return err
}

func IsExisted(key string) (int, error) {
	conn := db.GetRedisPool().Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXISTS", key))
}

func MdSalt(p string) string {
	salt := []byte("$%*&%99")
	hashmd := md5.New()
	io.WriteString(hashmd, p)
	password := fmt.Sprintf("%x", hashmd.Sum(salt))
	return password
}
