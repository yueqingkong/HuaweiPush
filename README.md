## 简介
使用华为推送服务给App推送日志

## 功能
- [x] 系统推送
- [ ] 页面跳转

## 使用
```
    huaweiPush:= push.HuaWeiPush{}
    huaweiPush.Init("***","****")
    users:= make([]string,0)
    users = append(users, "****")
    huaweiPush.Push("title","content",users)
```