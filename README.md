# moegraphdb
>用户关系内存数据库 

## 简介

支持用户关系
```
0:没有关系
1：user1 关注了 user2
2：user2 关注了 user1
3：互粉的好友

```

## restful接口

### like

GET /like?vid=1

>用户1所有关注的人的id（数组）

POST /like?vid=1&beliked=2

>用户1关注了用户2

DELETE /like?vid=1&beliked=2

>用户1取消关注了用户2



### fans

GET /fans?vid=1

>用户1所有的粉丝的id（数组）

POST /like?vid=1&fan=2

>用户1获得了粉丝用户2

DELETE /like?vid=1&fan=2

>用户1失去了粉丝用户2

### user

GET /user?vid=1

>用户1的更多key/value信息

POST /user?vid=1
```
{
 "key":"value"
}
```
>重新设置用户1的key/value信息

DELETE /user?vid=1
```
{
 "key":"value"
}
```
>更新用户1的一部分key/value信息

### relate
GET /fans?vid1=1&vid2=2

>用户1与用户2的关系

POST /like?vid=1&fan=2&relate=(0-3)

>设置用户1与用户2的关系

DELETE /like?vid=1&fan=2

>让用户1与用户2没有任何关系