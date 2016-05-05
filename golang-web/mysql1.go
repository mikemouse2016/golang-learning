package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"time"
        "fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:devchj@theL0805@tcp(localhost:3306)/stay?charset=utf8mb4")
	checkErr(err)

	//查询数据
//	rows, err := db.Query("select follower_id, count(follower_id) as count_follower from stay.app_user_follower where create_time > '2015-01-01 00:00:00' group by follower_id ORDER BY count_follower desc limit 0,10000")  //挤眼

	rows, err := db.Query("select user_id, count(user_id) as count_send from stay.app_user_follower where create_time > '2015-01-01 00:00:00' group by user_id ORDER BY count_send desc limit 0,10000") //送出的挤眼
	checkErr(err)
        score := make(map[int]int)
	for rows.Next() {
		var id int
                var count int
		err = rows.Scan(&id, &count)
		checkErr(err)
                sum := count/100
                _,ok := score[sum] // 假如key存在,则name = 李四 ，ok = true,否则，ok = false
                if ok{
                    score[sum] += 1
                }else{
                    score[sum]=1
                }
	}
        for key,value := range score {
               fmt.Println(key, value)
 	}
       db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
