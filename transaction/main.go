package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"time"
)

type Doctor struct {
	ID      int64
	Name    string
	Age     int
	Sex     int
	AddTime time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:19970904@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return
	}
	//延迟到函数结束关闭链接
	defer db.Close()


	//-------7、事务\--------
	tx, _ := db.Begin()
	result4, _ := tx.Exec("update doctor_tb set age = age + 1 where name = ?", "钟南山")
	result5, _ := tx.Exec("update doctor_tb set age = age + 1 where name = ?", "叶子")

	//影响行数，为0则失败
	i4, _ := result4.RowsAffected()
	i5, _ := result5.RowsAffected()
	if i4 > 0 && i5 > 0 {
		//2条数据都更新成功才提交事务
		err = tx.Commit()
		if err != nil {
			fmt.Println("事务提交失败", err)
			return
		}
		fmt.Println("事务提交成功")
	} else {
		//否则回退事务
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回退事务失败", err)
			return
		}
		fmt.Println("回退事务成功")
	}
}
