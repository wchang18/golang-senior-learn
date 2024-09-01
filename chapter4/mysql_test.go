package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"testing"
)

var mysqlDB *sql.DB

func GetDb() *sql.DB {
	var once sync.Once
	once.Do(func() {
		db, err := sql.Open("mysql", "root:root001@tcp(127.0.0.1:3306)/tests?charset=utf8mb4")
		if err != nil {
			panic(err)
		}
		mysqlDB = db
	})

	//if mysqlDB == nil {
	//	db, err := sql.Open("mysql", "root:root001@tcp(127.0.0.1:3306)/tests?charset=utf8mb4")
	//	if err != nil {
	//		panic(err)
	//	}
	//	mysqlDB = db
	//}

	return mysqlDB
}

func TestDb(t *testing.T) {
	db := GetDb()
	defer db.Close()
	err := db.Ping()
	fmt.Println(err)
}

func TestInsert(t *testing.T) {
	db := GetDb()
	defer db.Close()
	//insert into s_user(username,email) values ('jack','qq.com')
	res, err := db.Exec("insert into s_user(username,email) values(?,?)", "mat", "jack@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
	num, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(num)
}

func TestUpdate(t *testing.T) {
	db := GetDb()
	defer db.Close()
	//update s_user set username='tom' where id=2
	res, err := db.Exec("update s_user set email=? where id=?", "tom@qq.com", 2)
	if err != nil {
		t.Fatal(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(num)
}

func TestDelete(t *testing.T) {
	db := GetDb()
	defer db.Close()
	res, err := db.Exec("delete from s_user where id=?", 1)
	if err != nil {
		t.Fatal(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(num)
}

func TestSelectOne(t *testing.T) {
	db := GetDb()
	defer db.Close()

	row := db.QueryRow("select * from s_user where id = ?", 3)

	var user User
	row.Scan(&user.Id, &user.Username, &user.Email, &user.CreateAt)
	t.Logf("%+v", user)
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	CreateAt string `db:"create_at"`
}

func TestSelect(t *testing.T) {
	db := GetDb()
	defer db.Close()
	rows, err := db.Query("select * from s_user")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreateAt)
		if err != nil {
			t.Fatal(err)
		}
		users = append(users, user)
	}
	t.Logf("%+v", users)
}

func TestOrderBy(t *testing.T) {
	db := GetDb()
	defer db.Close()
	rows, err := db.Query("select * from s_user order by username desc")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreateAt)
		if err != nil {
			t.Fatal(err)
		}
		users = append(users, user)
	}
	t.Logf("%+v", users)
}

func TestLimit(t *testing.T) {
	db := GetDb()
	defer db.Close()
	rows, err := db.Query("select id,username from s_user order by id limit 0,10")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username)
		if err != nil {
			t.Fatal(err)
		}
		users = append(users, user)
	}
	t.Logf("%+v", users)
}

func TestGroupBy(t *testing.T) {
	db := GetDb()
	defer db.Close()
	rows, err := db.Query("select count(*) as count,email from s_user group by email")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var (
		count int64
		email string
	)
	for rows.Next() {
		err = rows.Scan(&count, &email)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%d %s", count, email)
	}
}

func TestCount(t *testing.T) {
	db := GetDb()
	defer db.Close()
	row := db.QueryRow("select count(*) as count from s_user ")

	var count int64
	row.Scan(&count)
	t.Log(count)
}
