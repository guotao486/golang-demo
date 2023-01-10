### redis list 命令
>Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）

>一个列表最多可以包含 232 - 1 个元素 (4294967295, 每个列表超过40亿个元素)。

|序号|命令|说明|
|:----|:-----|:-----|
|1|	[BLPOP key1 key2 ] timeout| 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
|2|	[BRPOP key1 key2 ] timeout| 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
|3|	BRPOPLPUSH source destination timeout| 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
|4|	LINDEX key index| 通过索引获取列表中的元素
|5|	LINSERT key BEFORE|AFTER pivot value| 在列表的元素前或者后插入元素
|6|	LLEN key| 获取列表长度
|7|	LPOP key| 移出并获取列表的第一个元素
|8|	[LPUSH key value1 value2]| 将一个或多个值插入到列表头部
|9|	LPUSHX key value| 将一个值插入到已存在的列表头部
|10|	LRANGE key start stop| 获取列表指定范围内的元素
|11|	LREM key count value| 移除列表元素
|12|	LSET key index value| 通过索引设置列表元素的值
|13|	LTRIM key start stop| 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
|14|	RPOP key| 移除列表的最后一个元素，返回值为移除的元素。
|15|	RPOPLPUSH source destination| 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
|16|	[RPUSH key value1 value2]| 在列表中添加一个或多个值
|17|	RPUSHX key value| 为已存在的列表添加值


```
# Lpush 命令将一个或多个值插入到列表头部
# Lindex 命令用于通过索引获取列表中的元素
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> lindex lang 0
golang
127.0.0.1:6379> lindex lang 1
python
127.0.0.1:6379> lindex lang 2
java
127.0.0.1:6379>

# Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
127.0.0.1:6379> del lang
1
127.0.0.1:6379> rpush lang java python golang
3
127.0.0.1:6379> lrange lang 0 -1
java
python
golang
127.0.0.1:6379>

# Lrange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。 其中 0 表示列表的第一个元素， 1 表示列表的第二个元素，以此类推。 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推

127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> lrange lang 0 -1
golang
python
java
127.0.0.1:6379> lrange lang 1 2
python
java
127.0.0.1:6379> lrange lang 0 -2
golang
python
127.0.0.1:6379>


# Rpoplpush 命令用于移除列表的最后一个元素，并将该元素添加到另一个列表并返回
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> rpoplpush lang lang2
java
127.0.0.1:6379> lrange lang 0 -1
golang
python
127.0.0.1:6379> lrange lang2 0 -1
java
127.0.0.1:6379>

# Blpop 命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> blpop lang 1
lang
golang

# Brpop 命令移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> brpop lang 1
lang
java

# Lrem 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。
127.0.0.1:6379> lpush mylist java java java python python golang golang golang
8
127.0.0.1:6379> lrem mylist 2 java
2
127.0.0.1:6379> lrange mylist 0 -1
golang
golang
golang
python
python
java
127.0.0.1:6379>

#  Llen 命令用于返回列表的长度。 如果列表 key 不存在，则 key 被解释为一个空列表，返回 0 。 如果 key 不是列表类型，返回一个错误。
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> llen lang
3
127.0.0.1:6379>

# Ltrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> ltrim lang 1 -1
OK
127.0.0.1:6379> lrange lang 0 -1
python
java

# Lpop 命令用于移除并返回列表的第一个元素
127.0.0.1:6379> del lang
1
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> lpop lang 2
golang
python
127.0.0.1:6379> lpop lang
java

# Rpop 命令用于移除并返回列表的最后一个元素
127.0.0.1:6379> del lang
0
127.0.0.1:6379> lpush lang java python golang
3
127.0.0.1:6379> lpop lang 2
golang
python
127.0.0.1:6379> lpop lang
java

# 
```