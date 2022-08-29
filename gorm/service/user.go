/*
 * @Author: GG
 * @Date: 2022-08-22 17:18:04
 * @LastEditTime: 2022-08-29 09:27:27
 * @LastEditors: GG
 * @Description: User ation
 * @FilePath: \golang-demo\gorm\service\user.go
 *
 */
package service

import (
	"errors"
	"fmt"
	"golang-demo/gorm/models"
	"time"

	"gorm.io/gorm"
)

func UserCreateTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func UserDropTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
}

// 创建数据
func UserCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "asdasd@qweq.com",
		Age:      12,
		Birthday: time.Now(),
	}
	result := db.Create(&user)
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
	fmt.Printf("user: %v\n", user)
}

// 创建数据，但只更新select 指定字段
func UserSelectCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "asdasd@qweq.com",
		Age:      12,
		Birthday: time.Now(),
	}

	db.Select("name", "age").Create(&user)
	fmt.Printf("user: %v\n", user)
}

// 创建数据，过滤掉select 指定字段
func UserOmitCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "dasdsa@dsad.com",
		Age:      25,
		Birthday: time.Now(),
	}
	db.Omit("name", "email", "age", "birthday").Create(&user)
	fmt.Printf("user: %v\n", user)
}

// 批量创建
func UserBatchCreate(db *gorm.DB) {
	users := []models.User{
		{Name: "tom"},
		{Name: "tom1"},
		{Name: "tom2"},
	}

	// 一次性
	// db.Create(&users)
	// 分批次
	db.CreateInBatches(users, 1)
	for _, u := range users {
		fmt.Printf("u.ID: %v\n", u.ID)
	}

}

// 根据map类型创建数据
// 不会触发hook和 时间字段追踪
func UserMapCreate(db *gorm.DB) {
	// 单条
	var user = map[string]interface{}{
		"Name": "tom", "Age": 18,
	}

	db.Model(&models.User{}).Create(user)

	// 批量
	var users = []map[string]interface{}{
		{"Name": "tom", "Age": 18},
		{"Name": "tom2", "age": 20},
	}
	db.Model(&models.User{}).Create(users)
}

// 关联数据创建，如果关联值是非零值，这些关联会被 upsert，且它们的 Hook 方法也会被调用
func UserCreditCardCreate(db *gorm.DB) {
	db.Debug().Create(&models.User{
		Name: "123",
		// CreditCard: models.CreditCard{
		// 	Number: "12312321313",
		// },
	})

	// 关联数据空
	db.Debug().Create(&models.User{
		Name: "321",
		// CreditCard: models.CreditCard{},
	})

	//db.Omit("CreditCard").Create(&user)

	// 跳过所有关联
	//db.Omit(clause.Associations).Create(&user)
}

// 测试默认值
func UserDefaultCreate(db *gorm.DB) {

	db.Create(&models.User{})
}

// 查询单条数据
func UserFind(db *gorm.DB) {
	user := models.User{}
	// 获取第一条记录（主键升序）
	db.First(&user)
	fmt.Printf("user: %v\n", user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	user2 := models.User{}
	db.Take(&user2)
	fmt.Printf("user2: %v\n", user2)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	user3 := models.User{}
	db.Last(&user3)
	fmt.Printf("user3: %v\n", user3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	// result.RowsAffected // 返回找到的记录数
	// result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 避免ErrRecordNotFound错误
	user4 := models.User{}
	db.Limit(1).Find(&user4)
	fmt.Printf("user4: %v\n", user4)
}

// first 和 last 主键排序，分别查询第一条和最后一条
// 只有在目标struct是指针或者通过db.Model()指定 model时有效
// 如果model没有主键，按第一个字段进行排序
func UserQueryFirst(db *gorm.DB) {

	// 有效
	user := &models.User{}
	db.First(user)

	// 有效
	result := map[string]interface{}{}
	db.Model(&models.User{}).First(result)

	// 无效
	result2 := map[string]interface{}{}
	db.Table("users").First(result2)

	// table 配合 take 有效
	result3 := map[string]interface{}{}
	db.Table("users").Take(result3)

	// 没有主键，会根据第一个字段排序，code
	type language struct {
		Code string
		Name string
	}
	db.First(&language{})
}

// first 根据主键查询
func UserQueryFirstById(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)
	fmt.Printf("user: %v\n", user)
	// SELECT * FROM `users` WHERE `users`.`id` = 4 ORDER BY `users`.`id` LIMIT 1

	user2 := &models.User{}
	db.Debug().First(user2, "4")
	fmt.Printf("user2: %v\n", user2)
	//SELECT * FROM `users` WHERE `users`.`id` = '4' ORDER BY `users`.`id` LIMIT 1

	users := &[]models.User{}
	db.Debug().First(users, []int{1, 2, 4})
	fmt.Printf("users: %v\n", users)
	//SELECT * FROM `users` WHERE `users`.`id` IN (1,2,4) ORDER BY `users`.`id` LIMIT 1
}

// 检索全部对象
func UserQueryFind(db *gorm.DB) {
	// 全部对象
	users := &[]models.User{}
	rows := db.Debug().Find(users)
	fmt.Printf("users: %v\n", users)
	fmt.Printf("rows: %v\n", rows.RowsAffected)
	fmt.Printf("rows.Error: %v\n", rows.Error)
	// SELECT * FROM `users`

	// 单条对象
	user := &models.User{}
	db.Debug().Find(user)
	fmt.Printf("user: %v\n", user)
	// SELECT * FROM `users`
}

func UserQueryWhereFind(db *gorm.DB) {
	users := &[]models.User{}

	// db.Debug().Where("name = ?", "tom").Find(users)
	//SELECT * FROM `users` WHERE name = 'tom'

	// db.Debug().Where("123").Find(users)
	//SELECT * FROM `users` WHERE `users`.`id` = '123', 默认是主键

	// db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	// db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	lastWeek := time.Now()
	db.Debug().Where("created_at < ?", lastWeek).Where("age = ?", 18).Find(&users)
	// SELECT * FROM `users` WHERE created_at < '2022-08-23 17:56:05.974'

	// BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	fmt.Printf("users: %v\n", users)
}

// 更新数据
func UserSave(db *gorm.DB) {
	user := &models.User{}
	db.First(user)
	fmt.Printf("user: %v\n", user)
	user.Name = "tom"
	user.Age = 18
	db.Debug().Save(user)
	fmt.Printf("user: %v\n", user)
}

// 更新单列
func UserUpdate(db *gorm.DB) {
	user := &models.User{}
	row := db.Debug().Model(user).Where("age = ?", 18).Update("name", "tomabc123")
	fmt.Printf("row.RowsAffected: %v\n", row.RowsAffected)
	//UPDATE `users` SET `name`='tomabc123',`updated_at`='2022-08-24 10:54:51.73' WHERE age = 18
	fmt.Printf("user: %v\n", user)

	user2 := &models.User{}
	db.Debug().Model(user2).Where("age = 0").Update("name", "tom0")
	//UPDATE `users` SET `name`='tom0',`updated_at`='2022-08-24 10:54:51.867' WHERE age = 0
	fmt.Printf("user: %v\n", user)

}

// 更新多列
func UserUpdates(db *gorm.DB) {
	user := &models.User{Name: "tomabc123", Email: "adasd@asdsa.com"}
	db.Debug().Where("age = ?", 18).Updates(user)
	//UPDATE `users` SET `name`='tomabc123',`updated_at`='2022-08-24 10:54:51.73' WHERE age = 18
	fmt.Printf("user: %v\n", user)

}

// 删除数据
// struct 拥有gorm.DetealAt字段，自动拥有软删除功能
func UserDelete(db *gorm.DB) {
	user := models.User{}
	// 根据struct主键或者主键参数
	db.Debug().Delete(&user, 1)
	// 额外删除条件
	db.Debug().Where("name = ?", "tom0").Delete(&user, 1)
}

// 查找软删除的数据
func UserUnscoped(db *gorm.DB) {
	users := []models.User{}
	db.Debug().Unscoped().Find(&users)
	//SELECT * FROM `users`
	db.Debug().Find(&users)
	//SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	// fmt.Printf("users: %v\n", users)
}

// 永久删除
func UserDrop(db *gorm.DB) {
	user := models.User{}
	db.Debug().Unscoped().Where("name = ?", "tom0").Delete(&user, 1)
	//DELETE FROM `users` WHERE name = 'tom0' AND `users`.`id` = 1
	fmt.Printf("user: %v\n", user)
}

// 原生sql
func UserRaw(db *gorm.DB) {
	type result struct {
		ID   uint
		Age  uint
		Name string
	}

	res := result{}
	db.Debug().Raw("select id,age,name from users where age > ?", 10).Scan(&res)
	// select id,age,name from users where age > 10
	fmt.Printf("res: %v\n", res)

	var age int
	db.Debug().Raw("select sum(age) from users").Scan(&age)
	// select sum(age) from users
	fmt.Printf("age: %v\n", age)

	// 修改
	db.Debug().Exec("update users set name = ? where id = ?", "tom123", 1)
	// update users set name = 'tom123' where id = 1

	// 不执行并打印sql
	user := models.User{}
	// stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	// SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	stmt := db.Session(&gorm.Session{DryRun: true}).Find(&user, 1).Statement
	// SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL
	fmt.Printf("stmt.SQL.String(): %v\n", stmt.SQL.String())
	// SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	fmt.Printf("stmt.Vars: %v\n", stmt.Vars)
}

// 关联数据新增
// User belong to Company
// User has Many CreditCard
// has Many 和 has one 的区别在于一个是[]struct 一个是struct
func UserRelatedCreate(db *gorm.DB) {

	// 新增一条user company 2条CreditCard 数据
	user := &models.User{
		Name: "12312",
		Age:  99,
		Company: models.Company{
			Name: "abc",
		},
		CreditCard: []models.CreditCard{
			{Number: "1232132312"},
			{Number: "123213123123"},
		},
		Languages: []models.Language{
			{Name: "zh"},
			{Name: "en"},
		},
	}

	db.Debug().Create(user)
	// db.Debug().Save(user)
}

// 关联数据更新
// User belong to Company
// Company 不是新数据的情况下新增User
func UserRelatedCreate2(db *gorm.DB) {
	company := models.Company{}
	db.First(&company, 2)
	// 只更新了User数据
	user := &models.User{
		Name:    "12312",
		Age:     99,
		Company: company,
	}

	db.Debug().Create(user)
}

// User has Many CreditCard
// 新增一条User 的 CreditCard信息
func UserRelatedCreate3(db *gorm.DB) {
	user := models.User{}
	db.Debug().Limit(1).Find(&user, 1)
	fmt.Printf("user: %v\n", user)
	card := models.CreditCard{
		Number: "111111",
		User:   user,
	}
	db.Debug().Create(&card)
}

// 修改数据
func UserRelatedUpdate(db *gorm.DB) {

	user := models.User{}
	db.Debug().First(&user, 2)

	user.Name = "22222"
	user.Age = 9

	card := models.CreditCard{}
	db.Debug().First(&card)
	card.Number = "222-222-222"
	user.CreditCard = []models.CreditCard{}
	user.CreditCard = append(user.CreditCard, card)

	// db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Updates(&user)
	db.Debug().Save(&user)
}

// 多对多关系，跳过upsert
func UserRelateSkipUpsert(db *gorm.DB) {
	user := &models.User{
		Name: "12312",
		Age:  99,
		Company: models.Company{
			Name: "abc",
		},
		Languages: []models.Language{
			{Name: "zh"},
			{Name: "en"},
		},
	}
	db.Omit("Languages").Create(&user)
}

// 通过关联查询数据
func UserRelateAssociation(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 1)

	languages := &[]models.Language{}

	// 查询
	// db.Debug().Model(user).Association("Languages").Find(languages)

	// 条件查询
	codes := []string{"en"}
	db.Debug().Model(user).Where("name in ?", codes).Association("Languages").Find(languages)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("len(user.Languages): %v\n", len(user.Languages))
	fmt.Printf("len(*languages): %v\n", len(*languages))
}

// 新增一条关联数据
func UserRelateAssociationAppend(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)
	// many2many 新增一条
	db.Model(&user).Association("Languages").Append(&models.Language{Name: "DE111"})

	// belong to 替换关联
	db.Model(&user).Association("Company").Append(&models.Company{Name: "DE"})
}

// 替换一条关联数据，元数据还在，只替换关联关系
func UserRelateAssociationReplace(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)

	l := &models.Language{}
	db.Debug().First(l, 7)
	// replace(model1,model2)
	// replace([]model{model1,model2})
	db.Model(user).Association("Languages").Replace(l)
}

// 删除关联关系
func UserRelateAssociationDelete(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)

	// belong to
	c := &models.Company{}
	db.Debug().First(c, 5)
	db.Debug().Model(user).Association("Company").Delete(c)

	cc := &[]models.CreditCard{}
	db.Debug().Model(user).Association("CreditCard").Find(cc)
	db.Debug().Model(user).Association("CreditCard").Delete(cc)
	// many2many
	l := &[]models.Language{}
	db.Debug().Model(user).Association("Languages").Find(l)
	db.Debug().Model(user).Association("Languages").Delete(l)
}

// 清空关联关系
func UserRelateAssociationClear(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)
	db.Debug().Model(user).Association("CreditCard").Clear()
}

// 关联数据计数
func UserRelateAssociationCount(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)
	db.Model(user).Association("Languages").Count()
}

// 删除源数据并删除关联数据，满足软删除条件默认软删除
func UserRelateAssociationDrop(db *gorm.DB) {
	user := &models.User{}
	db.Debug().First(user, 4)
	//has one、has many、many2many
	// company 没有被删除，因为是belong to
	// CreditCard 被删除了
	db.Debug().Select("Company", "CreditCard").Delete(user)
}

// 预加载 join，利用inner join加载关联数据，适合一对一关系
func UserRelateJoins(db *gorm.DB) {
	user := &models.User{}
	db.Debug().Joins("Company").First(user, 1)
	fmt.Printf("user: %v\n", user)
	fmt.Printf("user.Company: %v\n", user.Company)
}

// 预加载
func UserRelatePreload(db *gorm.DB) {
	c := []models.Company{}
	rows := db.Debug().Preload("Users").Find(&c)
	fmt.Printf("rows.RowsAffected: %v\n", rows.RowsAffected)

	for _, company := range c {
		fmt.Printf("c: %v\n", company)
		fmt.Printf("company.Users: %v\n", company.Users)
	}
}

func UserSelectDemo(db *gorm.DB) {
	user := []models.User{}

	db.Debug().Model(user).Find(user)

}
