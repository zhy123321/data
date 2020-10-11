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
	fmt.Println(u.Password)
	//用户数据保存
	row, err := db_mysql.Db.Exec("insert into user(phone, password)" + "values(?,?)",u.Phone,u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}