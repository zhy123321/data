package models

import (
	"Data_Attestatio/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

/*
 * 创建用户保存方法，用于用户数据保存
 */
func (u User) UserSave()(int64,error){
	//用户密码加密
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	//fmt.Println(u.Password)
	//fmt.Println(u)
	//用户数据保存
	row, err := db_mysql.Db.Exec("insert into user(phone , password)"+"values(?,?)",u.Phone,u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return id, nil
}

/*
 * 创建用户登录方法，用于用户数据比对
 */
func (u User) UserData()(*User, error){
	//由于用户密码经过哈希加密，所以要将前端密码经过哈希脱敏
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	//查询用户数据
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?", u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
