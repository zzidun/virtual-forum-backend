post    /adminlogin
post    /admins/
delete  /admins/:id
put     /admins/:id
get     /admins/
get     /admins/:id

post    /admingroups/
delete  /admingroups/:id
put     /admingroups/:id
get     /admingroups/
get     /admingroups/:id

post    /bans/
delete  /bans/:id
put     /bans/:id
get     /bans/
get     /bans/:id

post    /categories/
delete  /categories/:id
put     /categories/:id
get     /categories/
get     /categories/:id


post    /categoryers/:id
delete  /categoryers/:id
put     /categoryers/:id
get     /categoryers/
get     /categoryers/:id

post    /login
post    /register

get     /categories/
get     /categories/:id

post    /posts/
delete  /posts/:id
put     /posts/:id
get     /posts/
get     /posts/:id


post    /comments/
delete  /comments/:id
put     /comments/:id
get     /comments/
get     /comments/:id


post    /shields/
delete  /shields/:id
put     /shields/:id
get     /shields/
get     /shields/:id

post    /collects/
delete  /collects/:id
put     /collects/:id
get     /collects/
get     /collects/:id

post    /follows/
delete  /follows/:id
put     /follows/:id
get     /follows/
get     /follows/:id


代码

请求

| 功能 | 代码 |
| --- | --- |
| 管理员登陆 | 1000 |
| 管理员注销 | 1001 | 
| 小管理员创建 | 1002 |
| 管理员组创建 | 1010 |
| 管理员组权限 | 1011 |
| 封禁ip | 1020 |
| 封禁用户 | 1021 |
| 查看版块列表 | 1030 |
| 创建版块 | 1031 |
| 删除版块 | 1032 |
| 设置版主 | 1033 |
| 删除帖子 | 1040 |
| 删除楼层 | 1041 |
| 查询版块数量 | 1050 |
| 查询总发帖数量 | 1051 |
| 查询版块发帖数量 | 1052 |
| 查询总楼层数量 | 1053 |
| 查询版块楼层数量 | 1054 | 
| 查询帖子楼层数量 | 1055 | 
| 查询总用户数量 | 1056 | 
| 查询版块用户数量 | 1057 | 
| 查询登陆失败ip信息 | 1060 | 
| 查询登陆失败用户信息 | 1061 |

回复
| 功能 | 代码 |
| --- | --- |
| 操作成功 | 0000 |
| 未登陆 | 0001 |
| 无对应权限 | 0002 |
| 含有具体消息 | 0003 |





管理员登陆
{
    200,
    {

    }
}