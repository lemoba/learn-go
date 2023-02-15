package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func init() {
	dsn := "lemoba:Ranen520..@tcp(121.4.32.131:3310)/test?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Failed to create a connection to database: ", err.Error())
	}
	db = conn
}

type User struct {
	gorm.Model
	UserName  string    `gorm:"column:username"`
	LoginTime time.Time `gorm:"column:login_time"`
	Email     string    `gorm:"column:email"`
}

func MulUpdate(users []User) []uint {
	db.Create(&users)
	// db.CreateInBatches(&users, 1) // 分批更新

	var Ids []uint

	for _, user := range users {
		Ids = append(Ids, user.Model.ID)
	}
	return Ids
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	fmt.Println("BeforeCreate")
//	return
// }
//
// func (u *User) AfterCreate(tx *gorm.DB) (err error) {
//	fmt.Println("AfterCreate")
//	return
// }

func create() {
	fmt.Println("create...")
	// user := User{UserName: "david", LoginTime: time.Now(), Email: "david@qq.com"}
	// result := db.Create(&user)
	//
	// fmt.Println(result.RowsAffected)
	// fmt.Println(result.Error)
	// fmt.Println(user.Id)

	var users = []User{
		{UserName: "tom", LoginTime: time.Now(), Email: "tom@qq.com"},
		{UserName: "tom2", LoginTime: time.Now(), Email: "tom1@qq.com"},
		{UserName: "tom3", LoginTime: time.Now(), Email: "tom2@qq.com"},
	}

	// 批量更新
	Ids := MulUpdate(users)
	fmt.Println(Ids)

	// 选择指定字段更新
	// res := db.Select("UserName", "Email").Create(&user)

	// 忽略指定字段更新
	// res := db.Omit("Email").Create(&user)

	// fmt.Println(res.RowsAffected)
	// fmt.Println(user.ID)
}

func query(user User) {
	// fmt.Println("query.....")

	// db.First(&user) // 获取第一条记录(主键升序)
	// db.Take(&user) // 获取一条记录，没有指定排序字段(默认主键升序)
	// result := db.Last(&user)                              // 获取最后一条记录(主键降序)
	// re := errors.Is(result.Error, gorm.ErrRecordNotFound) // 检查 ErrRecordNotFound 错误
	// fmt.Println(re)

	// db.First(&user, 20)   where id = 20 // 根据ID查找
	// db.First(&user, "20") // 根据ID查找
	// db.First(&user, "username = ?", "tom") // 指定查询条件

	// db.Where("username = ?", "tom").First(&user) // where username='tom'

	// fmt.Println(user)
}
func queryMul(users []User) {
	// db.Find(&users, []int{18, 19, 20}) // where id IN (18, 19, 20)
	// var nameList []string
	//
	// for _, user := range users {
	// 	nameList = append(nameList, user.UserName)
	// }
	// fmt.Println(nameList)

	// db.Where("username <> ?", "tom").Find(&users) // where username <> 'tom'

	// db.Where("username IN ? ", []string{"tom", "tom2"}).Find(&users) // where username IN ("tom", "tom2")

	// db.Where("username LIKE ?", "%om%").Find(&users) // where username LIKE '%1%'

	// db.Where("username = ? AND email LIKE ?", "tom", "%tom@qq.com%").Find(&users)

	// lastWeek := time.Now().Format("2006-01-02 15:04:05")
	// db.Where("created_at < ?", lastWeek).Find(&users)

	fmt.Println(users)
}
func main() {
	user := User{}
	query(user)
	users := []User{}
	queryMul(users)
}
