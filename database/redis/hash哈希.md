<!--
 * @Author: GG
 * @Date: 2023-01-09 17:18:27
 * @LastEditTime: 2023-01-10 10:38:04
 * @LastEditors: GG
 * @Description: 
 * @FilePath: \redis\hash.md
 * 
-->
### hash命令
>Redis hash 是一个 string 类型的 field（字段） 和 value（值） 的映射表，hash 特别适合用于存储对象。

>Redis 中每个 hash 可以存储 232 - 1 键值对（40多亿）

|序号	|命令	|说明|
| :------------- | :----------: | ------------: |
|1|	[HDEL key field1 field2]| 删除一个或多个哈希表字段|
|2|	HEXISTS key field| 查看哈希表 key 中，指定的字段是否存在。|
|3|	HGET key field| 获取存储在哈希表中指定字段的值。|
|4|	HGETALL key| 获取在哈希表中指定 key 的所有字段和值|
|5|	HINCRBY key field increment| 为哈希表 key 中的指定字段的整数值加上增量 increment 。|
|6|	HINCRBYFLOAT key field increment| 为哈希表 key 中的指定字段的浮点数值加上增量 increment 。|
|7|	HKEYS key| 获取所有哈希表中的字段|
|8|	HLEN key| 获取哈希表中字段的数量|
|9|	[HMGET key field1 field2]| 获取所有给定字段的值|
|10|	[HMSET key field1 value1 field2 value2 ]| 同时将多个 field-value (域-值)对设置到哈希表 key 中。|
|11|	HSET key field value| 将哈希表 key 中的字段 field 的值设为 value 。|
|12|	HSETNX key field value| 只有在字段 field 不存在时，设置哈希表字段的值。|
|13|	HVALS key| 获取哈希表中所有值。|
|14|	HSCAN key cursor [MATCH pattern] [COUNT count]| 迭代哈希表中的键值对。|


```
# Hset 命令用于为哈希表中的字段赋值 
# Hget 命令用于返回哈希表中指定字段的值
127.0.0.1:6379> del site
1
127.0.0.1:6379> hset site name "golang-tech-stack.com"
1
127.0.0.1:6379> hset site users 300000
1
127.0.0.1:6379> hget site name
golang-tech-stack.com
127.0.0.1:6379> hget site users
300000
127.0.0.1:6379>

# Hmset 命令用于同时将多个 field-value (字段-值)对设置到哈希表中
# Hmget 命令用于返回哈希表中，一个或多个给定字段的值
127.0.0.1:6379> del site
1
127.0.0.1:6379> hmset site name "golang技术栈" url "golang-tech-stack.com" author "老郭"
OK
127.0.0.1:6379> hgetall site
name
golang技术栈
url
golang-tech-stack.com
author
老郭

# hkeys 获取所有哈希表中的字段 
# hvals 获取哈希表中所有值
127.0.0.1:6379> hkeys site
name
users
127.0.0.1:6379> hvals site
golang-tech-stack.com
300000
127.0.0.1:6379>


# hdel 删除一个或多个哈希表字段
127.0.0.1:6379> del site
1
127.0.0.1:6379> hmset site name "golang技术栈" url "golang-tech-stack.com" author "老郭"
OK
127.0.0.1:6379> hdel site name
1
127.0.0.1:6379> hdel site url
1
127.0.0.1:6379> hgetall site
author
老郭

# Hexists 命令用于查看哈希表的指定字段是否存在
127.0.0.1:6379> hmset site name "golang技术栈" url "golang-tech-stack.com" author "老郭"
OK
127.0.0.1:6379> hexists site name
1
127.0.0.1:6379> hexists site ip
0
127.0.0.1:6379>

# Hincrby 命令用于为哈希表中的字段值加上指定增量值 

127.0.0.1:6379> del score
1
127.0.0.1:6379> hset score english 100
1
127.0.0.1:6379> hincrby score english 10
110
127.0.0.1:6379> hincrby score english -10
100
127.0.0.1:6379>

# Hincrbyfloat 命令用于为哈希表中的字段值加上指定浮点数增量值

127.0.0.1:6379> del score
1
127.0.0.1:6379> hset score english 90.5
1
127.0.0.1:6379> hincrbyfloat score english 0.5
91
127.0.0.1:6379>

```