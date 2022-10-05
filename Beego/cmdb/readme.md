用户管理
1. 登录
2. 用户增、删、改、查

用户信息
id
name
nickname
password
tel
gender
staff_id
addr
email
department
status

created_at
updated_at
deleted_at

部门管理

登录成功后显示用户列表
url -> 用户列表页面展示
Controller => Model(获取用户数据) => view => Router


用户认证
记录用户状态? 记录在哪里？
HTTP无状态 下一次请求

cookie sessio 机制
状态记录 => session
状态的跟踪 => session