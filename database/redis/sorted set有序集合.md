<!--
 * @Author: GG
 * @Date: 2023-01-10 11:53:01
 * @LastEditTime: 2023-01-10 14:44:20
 * @LastEditors: GG
 * @Description: 
 * @FilePath: \redis\sorted set有序集合
 * 
-->
### Redis 有序集合（Sorted Set）
>Redis 有序集合和集合一样也是 string 类型元素的集合,且不允许重复的成员。

>不同的是每个元素都会关联一个 double 类型的分数。redis 正是通过分数来为集合中的成员进行从小到大的排序。

>有序集合的成员是唯一的,但分数(score)却可以重复。

>集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。 集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。

### 命令

|序号|命令|说明|
|:---|:---|:---|
|1|	[ZADD key score1 member1 score2 member2]| 向有序集合添加一个或多个成员，或者更新已存在成员的分数|
|2|	ZCARD key |获取有序集合的成员数|
|3|	ZCOUNT key min max |计算在有序集合中指定区间分数的成员数|
|4|	ZINCRBY key increment member |有序集合中对指定成员的分数加上增量 increment|
|5|	[ZINTERSTORE destination numkeys key key ...] |计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 destination 中|
|6|	ZLEXCOUNT key min max |在有序集合中计算指定字典区间内成员数量|
|7|	[ZRANGE key start stop WITHSCORES] |通过索引区间返回有序集合指定区间内的成员|
|8|	[ZRANGEBYLEX key min max LIMIT offset count] |通过字典区间返回有序集合的成员|
|9|	ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT] |通过分数返回有序集合指定区间内的成员|
|10|	ZRANK key member |返回有序集合中指定成员的索引|
|11|	[ZREM key member member ...] |移除有序集合中的一个或多个成员|
|12|	ZREMRANGEBYLEX key min max |移除有序集合中给定的字典区间的所有成员|
|13|	ZREMRANGEBYRANK key start stop |移除有序集合中给定的排名区间的所有成员|
|14|	ZREMRANGEBYSCORE key min max |移除有序集合中给定的分数区间的所有成员|
|15|	[ZREVRANGE key start stop WITHSCORES] |返回有序集中指定区间内的成员，通过索引，分数从高到低|
|16|	[ZREVRANGEBYSCORE key max min WITHSCORES] |返回有序集中指定分数区间内的成员，分数从高到低排序|
|17|	ZREVRANK key member |返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序|
|18|	ZSCORE key member |返回有序集中，成员的分数值|
|19|	[ZUNIONSTORE destination numkeys key key ...] |计算给定的一个或多个有序集的并集，并存储在新的 key 中|
|20|	ZSCAN key cursor [MATCH pattern] [COUNT count] |迭代有序集合中的元素（包括元素成员和元素分值）|