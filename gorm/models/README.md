## Gorm Models  模型声明 

> Product.go使用gorm约定models


> User.go 使用自定义配置


### 字段权限配置
```
type User struct {
 Name string `gorm:"<-:create"` // 允许读和创建
 Name string `gorm:"<-:update"` // 允许读和更新
 Name string `gorm:"<-"`        // 允许读和写（创建和更新）
 Name string `gorm:"<-:false"`  // 允许读，禁止写
 Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
 Name string `gorm:"->;<-:create"` // 允许读和写
 Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
 Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
}
```

### 时间追踪，纳秒、毫秒、秒、Time
> GORM 约定使用CreateAt、UpdateAt 来自动追踪创建/修改时间
> 如果使用不同的字段可以配置 `autoCreateTime` `autoUpdateTime` 标签
> 如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int，并在标签加上`:nano` `:milli`

```
type user struct {
	Name string
	CreateTime time.Time  `gorm:"autoCreateTime"`
	UpdateAt int64 		`gorm:"autoUpdateTime:nano"`
}
```

### 嵌入结构体

对于匿名字段，GORM 会将其字段包含在父结构体中，例如：

```
type user struct {
	gorm.Model
	Name string
}

// 等效于

type user struct {
	ID 	uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeleteAt `gorm:"index"`
	Name string
}

```
于正常的结构体字段，你也可以通过标签 embedded 将其嵌入，例如：
```
type Author struct {
	Name string
	Email string
}

type Blog struct {
	ID	uint
	Author Author `gorm:"embedded"`
}

// 等效于

type Blog struct {
	ID uint
	Name string
	Email string
}
```
并且，您可以使用标签 embeddedPrefix 来为 db 中的字段名添加前缀，例如：
```
type Author struct {
	Name string
	Email string
}

type Blog struct {
	ID	uint
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
}

// 等效于

type Blog struct {
	ID uint
	AuthorName string
	AuthorEmail string
}
```
字段标签

| 标签名                 | 说明                                                                                                |
| ---------------------- | --------------------------------------------------------------------------------------------------- |
| column                 | 指定db列名                                                                                          |
| type                   | 列数据类型，bool、int、uint、float、string、time、bytes                                             |
| size                   | 指定列大小 `size:256`                                                                               |
| primaryKey             | 指定列为主键                                                                                        |
| unique                 | 指定列为唯一                                                                                        |
| default                | 指定列的默认值                                                                                      |
| precision              | 指定列的精度                                                                                        |
| scale                  | 指定列的大小                                                                                        |
| not null               | 指定列非空                                                                                          |
| autoIncrement          | 指定列自增                                                                                          |
| autoIncrementIncrement | 自增步长                                                                                            |
| embedded               | 嵌套字段                                                                                            |
| embeddedPrefix         | 嵌套字段前缀                                                                                        |
| autoCreateTime         | 自动填充time类型的字段，int类型可使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli |
| index                  | 根据字段创建索引，详情查询gorm 索引                                                                 |
| uniqueIndex            | 创建唯一索引                                                                                        |
| check                  | 约束 `check age>18`，更多详情查询gorm 约束                                                          |
| <-                     | 设置字段写入权限， `<-:create` 只创建,`<-:update`只修改，`<-:false`无写入，`<-`创建和更新           |
| ->                     | 设置字段读权限，`->:false`无读权限                                                                  |
| -                      | 忽略字段，无读写                                                                                    |
| comment                | 迁移时为字段添加注释                                                                                |
