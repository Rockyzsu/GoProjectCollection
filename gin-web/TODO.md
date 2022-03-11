<h1>Gin Web TODO LIST (Change Logs)</h1>

> 这也是我的学习路线, 新手可以根据每个步骤从头尝试, 可看对应commit id.
> 主要是一些常用的功能, 如果你有一些实用或好玩的功能, 可以推荐给我学习哦~

## 待完成 OR 更新日志

### 2022-01

- [x] API接口签名(防重放攻击、防数据篡改)
- [x] 密码输错次数过多需输入验证码
- [x] 数据字典
- [x] 提升go版本
- [x] 文档(swagger && slate)
- [x] 错误处理github.com/pkg/errors

### 2021-10

- [x] 将部分公共代码移入go-helper包, 各个模块按需自行配置, 从tag v1.2.0开始(如果你有多个项目, 需要不同的后端, 那么这样会大大节省效率)
- [x] 工作流切换为状态机(finite state machine), 常见的审批流程完全够用, 而不是去设计很复杂的工作流
- [x] 中英文双语更友好, 将界面展示用到的中文全部替换
- [x] 菜单图标使用ali iconfont, 随心替换自己想要的图标

### 2020-12

- [x] 全局接口幂等性中间件
- [ ] 发现一些接口自动化测试工具, 好用的话会加入到项目中
- [x] 请求参数绑定优化(数字类型兼容, 如{"key":1}和{"key":"1"}, 现在定义ReqNumber可以正常解析)
- [x] Mysql Binlog: drop/truncate table同步刷新到redis
- [x] 前端代码构建是自动上传到阿里云OSS对象存储, 加快网页访问速度

### 较早之前

- [x] Gorm连接Mysql
- [x] Gin搭建简易web服务器
- [x] 全局日志管理
- [x] 跨域中间件
- [x] 全局异常中间件
- [x] Jwt认证中间件(jwt无状态与session有状态的理解)
- [x] 用户密码加密(hash比对不可逆, 且同一明文每次生成的密文都不一样)
- [x] 接口访问日志
- [x] 前端页面搭建
- [x] 系统管理菜单: 菜单/角色/用户/接口管理
- [x] 后端参数校验
- [x] Casbin资源访问授权中间件
- [x] 配置管理以及配置文件打包到二进制中
- [x] 全局事务管理中间件
- [x] 速率限制中间件
- [x] Docker镜像打包发布
- [x] Mysql Binlog监听器, 并将数据低延迟刷到Redis
- [x] Gorm时间字段转换为本地时间
- [x] 工作流管理
- [x] 以请假审批为例完成工作流测试(wangwu提交=>zhangsan一级审批=>admin二级审批=>流程结束)
- [x] Redis查询优化(结合gorm源码, 实现redis查询json时可以Preload关联表)
- [x] 上传组件(大文件分块上传/多文件、文件夹上传, 根据实际项目需要再扩展断点续传等)
- [x] 全局代码审计(用户访问哪些接口, 做了哪些操作, 操作内容是什么, 根据实际需求定制)
- [x] 上传组件新增ZIP解压示例
- [x] 用户登录通过RSA非对称加密传输用户密码, 避免密码泄漏风险
- [x] 上线在线演示
- [x] 最近发现一款对象存储服务[Minio](https://github.com/minio/minio), 准备研究下
- [x] 消息中心(活跃用户上线时新增消息表, 不活跃用户不管, 有效降低数据量)
- [x] 消息中心修改为websocket长连接保证实时性
- [x] 轻量级日志收集工具Grafana loki(ELK太重, 本地测试至少需要4G内存才基本跑起来)