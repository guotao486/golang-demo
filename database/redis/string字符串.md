### Redis 字符串（string）
>Redis 字符串数据类型的相关命令用于管理 redis 字符串值
### string命令

|序号	|命令	|说明|
| :------------- | :---------- | ------------ |
|1|	SET key value| 设置指定 key 的值。|
|2|	GET key |获取指定 key 的值。|
|3|	GETRANGE key start end |返回 key 中字符串值的子字符|
|4|	GETSET key value |将给定 key 的值设为 value ，并返回 key 的旧值(old value)。|
|5|	GETBIT key offset |对 key 所储存的字符串值，获取指定偏移量上的位(bit)。|
|6|	[MGET key1 key2..] |获取所有(一个或多个)给定 key 的值。|
|7|	SETBIT key offset value |对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。|
|8|	SETEX key seconds value |将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。|
|9|	SETNX key value |只有在 key 不存在时设置 key 的值。|
|10|	SETRANGE key offset value |用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始。|
|11|	STRLEN key |返回 key 所储存的字符串值的长度。|
|12|	[MSET key value key value ...] |同时设置一个或多个 key-value 对。|
|13|	[MSETNX key value key value ...] |同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。|
|14|	PSETEX key milliseconds value |这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。|
|15|	INCR key |将 key 中储存的数字值增一。|
|16|	INCRBY key increment |将 key 所储存的值加上给定的增量值（increment） 。|
|17|	INCRBYFLOAT key increment |将 key 所储存的值加上给定的浮点增量值（increment） 。|
|18|	DECR key |将 key 中储存的数字值减一。|
|19|	DECRBY key decrement key |所储存的值减去给定的减量值（decrement） 。
|20|	APPEND key value |如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。|

```
# set不存在设置，存在修改
127.0.0.1:6379> set name golang
OK
127.0.0.1:6379> get name
golang
127.0.0.1:6379> set name golang-tech-stack.com
OK
127.0.0.1:6379> get name
golang-tech-stack.com
127.0.0.1:6379>

# get存在返回，不存在返回空
127.0.0.1:6379> get name
golang-tech-stack.com
127.0.0.1:6379> get site
127.0.0.1:6379>

# mset mget 设置获得多个key的值
127.0.0.1:6379> mset site golang-tech-stack.com name golang技术栈
OK
127.0.0.1:6379> mget site name
golang-tech-stack.com
golang技术栈

# strlen
127.0.0.1:6379> set name hello
OK
127.0.0.1:6379> strlen name
5
# Setex 命令为指定的 key 设置值及其过期时间
127.0.0.1:6379> setex name 3 golang
OK
127.0.0.1:6379> get name
golang
127.0.0.1:6379> get name
# 三秒后失效过期
127.0.0.1:6379>

# Setnx（SET if Not eXists） 命令在指定的 key 不存在时，为 key 设置指定的值
127.0.0.1:6379> setnx name java
0
127.0.0.1:6379> get name
golang

# incr将 key 中储存的数字值增一
127.0.0.1:6379> set score 90
OK
127.0.0.1:6379> incr score
91
127.0.0.1:6379> get score
91

# decr将 key 中储存的数字值增一
127.0.0.1:6379> decr score
90
127.0.0.1:6379> get score
90

# incrby 将 key 所储存的值加上给定的增量值
127.0.0.1:6379> del score
1
127.0.0.1:6379> set score 100
OK
127.0.0.1:6379> incrby score 10
110
127.0.0.1:6379> get score
110
# decrby key 所储存的值减去给定的减量值
127.0.0.1:6379> decrby score 20
90
127.0.0.1:6379> get score
90

# Append 命令用于为指定的 key 追加值
127.0.0.1:6379> set name golang
OK
127.0.0.1:6379> append name -tech-stack.com
21
127.0.0.1:6379> get name
golang-tech-stack.com

# Setrange 命令用指定的字符串覆盖给定 key 所储存的字符串值，覆盖的位置从偏移量 offset 开始
127.0.0.1:6379> set name golang技术栈
OK
127.0.0.1:6379> setrange name 6 -tech-stack.com
21
127.0.0.1:6379> get name
golang-tech-stack.com
127.0.0.1:6379>
```