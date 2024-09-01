package gorm_learn

import (
	"fmt"
	"github.com/thoas/go-funk"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var GormDb *gorm.DB

func GetGorm() *gorm.DB {
	var once sync.Once
	once.Do(func() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{},
		)

		db, err := gorm.Open(mysql.Open("root:root001@tcp(127.0.0.1:3306)/tests?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic(err)
		}
		GormDb = db
	})
	return GormDb
}

func TestPing(t *testing.T) {
	db := GetGorm()
	err := db.Exec("select 1").Error
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

// 添加单条数据
func TestGormCreate(t *testing.T) {
	db := GetGorm()
	user := User{
		Username:  "test",
		Email:     "test@qq.com",
		Uid:       GetUid(),
		Telephone: "12345678901",
	}
	res := db.Create(&user) //创建用户
	t.Log(res.RowsAffected, res.Error)
	t.Logf("%+v", user)
}

// 批量添加数据
func TestGormCreateBatch(t *testing.T) {
	db := GetGorm()
	var (
		users []User
	)
	for i := 1; i <= 10; i++ {
		user := User{}
		user.Username = fmt.Sprintf("test%d", i)
		user.Email = fmt.Sprintf("test%d@qq.com", i)
		user.Uid = GetUid()
		users = append(users, user)
	}
	res := db.Create(&users) //批量创建
	t.Log(res.RowsAffected, res.Error)
}

func GetUid() string {
	now := time.Now()
	timestamp := now.Unix()
	timestamp = timestamp << 3
	return fmt.Sprintf("%d%d", timestamp, funk.RandomInt(1000, 9999))
}

func TestGormUpdate(t *testing.T) {
	db := GetGorm()
	var user User
	db.Model(User{}).First(&user, 1)
	fmt.Printf("%+v", user)
	user.Email = "gg@qq.com"
	user.Telephone = "123123"
	res := db.Save(user)
	t.Log(res.RowsAffected)
}

func TestGormUpdate2(t *testing.T) {
	db := GetGorm()
	res := db.Model(User{}).
		Where("id", 2).
		Update("email", "gg2@qq.com")
	t.Log(res.RowsAffected)

	res = db.Model(User{}).Where("id", 3).Updates(map[string]interface{}{
		"email":     "gg3@qq.com",
		"telephone": "123123",
	})
	t.Log(res.RowsAffected)
}

func TestGormDelete(t *testing.T) {
	db := GetGorm()
	res := db.Where("id", 4).Delete(User{})
	t.Log(res.RowsAffected)

	res = db.Delete(User{}, 5)
	t.Log(res.RowsAffected)

	users := []User{{ID: 6}, {ID: 7}}
	res = db.Delete(&users)
	t.Log(res.RowsAffected)
}

func TestSelectOne(t *testing.T) {
	db := GetGorm()
	var (
		user  User
		user2 User
		user3 User
		user4 User
	)

	//SELECT * FROM `user` ORDER BY `user`.`id` LIMIT 1
	db.Debug().First(&user) //获取第一个数据
	t.Log(user)

	//SELECT * FROM `user` LIMIT 1
	db.Debug().Take(&user2) //获取随机一条
	t.Log(user2)

	//SELECT * FROM `user` ORDER BY `user`.`id` DESC LIMIT 1
	db.Debug().Last(&user3) //获取最后一个数据
	t.Log(user3)

	//SELECT * FROM `user` WHERE `user`.`id` = 2
	db.Debug().Take(&user4, 2) //获取id=2的数据
	t.Log(user4)
}

func PrintUserId(users []User) {
	for _, user := range users {
		fmt.Println(user.ID, user.Username)
	}
}

func TestSelectMany(t *testing.T) {
	db := GetGorm()
	var users []User
	//SELECT * FROM `user`
	db.Model(User{}).Debug().
		Find(&users)
	t.Logf("%+v", users)
	PrintUserId(users)

	//SELECT * FROM `user` WHERE id > 10 and id <= 12
	db.Model(User{}).Debug().
		Find(&users, "id > ? and id <= 12", 10)
	t.Logf("%+v", users)
	PrintUserId(users)

	//SELECT * FROM `user` WHERE `id` NOT IN (1,2)
	db.Model(User{}).Debug().
		Not("id", []int{1, 2}).
		Find(&users)
	t.Logf("%+v", users)
	PrintUserId(users)

	//SELECT * FROM `user` WHERE id = 1 OR id = 2
	db.Model(User{}).Debug().
		Or("id = ?", 1).Or("id = ?", 2).Find(&users)
	t.Logf("%+v", users)
	PrintUserId(users)
}

func TestChain(t *testing.T) {
	db := GetGorm()
	var users []User
	//SELECT * FROM `user` WHERE id > 1 ORDER BY id desc LIMIT 3 OFFSET 1
	db.Model(User{}).Debug().
		Where("id > ?", 1).
		Order("id desc").
		Limit(4).Offset(1).Find(&users)
	PrintUserId(users)

	//SELECT count(*) as cnt, username FROM `user` GROUP BY `username` HAVING username != 'test10'
	var list []map[string]interface{}
	db.Model(User{}).Debug().
		Select("count(*) as cnt, username").
		Group("username").
		Having("username != ?", "test10").
		Scan(&list)
	//fmt.Printf("%+v", list)
	for _, item := range list {
		fmt.Printf("%+v\n", item)
	}
}

func TestJoin(t *testing.T) {
	db := GetGorm()
	//SELECT u.id,u.username,u.telephone,uc.company_name FROM user as u left join user_company as uc on u.id = uc.userId
	var users []map[string]any
	db.Table("user as u").Debug().
		Select("u.id,u.username,u.telephone,uc.company_name").
		Joins("left join user_company as uc on u.id = uc.userId").
		Scan(&users)
	t.Logf("%+v", users)
}

func TestRow(t *testing.T) {
	db := GetGorm()
	var user User
	db.Raw("select * from user where id = ?", 1).Scan(&user)
	t.Logf("%+v", user)
}

func TestTransaction(t *testing.T) {
	db := GetGorm()
	//第一种方式
	//res := db.Transaction(func(tx *gorm.DB) error {
	//	var user = User{
	//		Username:  GetUserName(),
	//		Email:     "go@qq.com",
	//		Uid:       GetUid(),
	//		Telephone: "12345678901",
	//	}
	//	err := tx.Create(&user).Error
	//	if err != nil {
	//		return err //errors.New("err test")
	//	}
	//
	//	err = tx.Create(&UserCompany{
	//		CompanyName: GetCompany(),
	//		UserId:      user.ID,
	//	}).Error
	//	if err != nil {
	//		return err
	//	}
	//	//return errors.New("err test")
	//	return nil
	//})
	//t.Log(res)

	//第二种方式
	tx := db.Begin()
	var user = User{
		Username:  GetUserName(),
		Email:     "go@qq.com",
		Uid:       GetUid(),
		Telephone: "12345678901",
	}
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		t.Log(err)
	}
	for i := 0; i < 2; i++ {
		err = tx.Create(&UserAddress{
			Address: GetAddress(),
			UserId:  user.ID,
		}).Error
		if err != nil {
			tx.Rollback()
			t.Log(err)
			break
		}
	}

	if err == nil {
		tx.Commit()
	}

}

func GetUserName() string {
	return "user-" + funk.RandomString(6)
}

func GetAddress() string {
	return "address-" + funk.RandomString(6)
}

func GetCompany() string {
	return "company-" + funk.RandomString(6)
}

func TestHaveOne(t *testing.T) {
	db := GetGorm()
	var user User
	db.Model(User{}).Debug().
		Preload("UserCompany").
		First(&user, 18)
	t.Logf("%+v", user)
	t.Log(user.UserCompany)
}

func TestHaveMany(t *testing.T) {

	db := GetGorm()
	var user User
	db.Model(User{}).Debug().
		Preload("UserAddress").
		Preload("UserCompany").
		First(&user, 18)
	t.Logf("%+v", user)
	t.Log(user.UserAddress)
	t.Log(user.UserCompany)
}
