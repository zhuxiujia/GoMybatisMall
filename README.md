# GoMybatisMall
GoMybatis based  mall project

注意：本项目架构拷贝自生产环境源码的改版，经过生产环境检验！

基于GoMybatis的商城 微服务 项目
* 适用于给 初学GoMybatis开发者作为 示例或参考用或者快速开发一个GoMybatis项目.
* 适用于Java服务端开发者学习适用go项目（jvm开发者一般都是公司性质，注重扩展性和简单化，适用于习惯SSM组合的开发者，即SpringBoot,SpringMVC,Mybatis）

### 设计
+ 本项目使用基本的微服务设计（Controller-Service Discovery-Rpc-Service）
+ 关于网关（自行设计controller 作为网关做鉴权，注意本项目作为案例未加入，但是可以把controller层视为网关）
+ 关于链路追踪（需要自行封装链路追踪，本项目基于rpc ，使用方式上没差别。注意本项目未加入）
+ 关于熔断（需要自行封装熔断，本项目基于rpc ，使用方式上没差别。注意本项目未加入）

### 架构:
* 必须的库
+ log（日志推荐使用golang官方（官方最稳定，其他库需要考虑安全性），仅写入到静态文件）
+ consul(服务发现，consul是支持spring cloud的go语言分布式服务发现) 
+ easy_mvc(https://github.com/zhuxiujia/easy_mvc) （拟mvc设计）
+ easyrpc(https://github.com/zhuxiujia/easyrpc) （封装定义,支持事务传播，基于go官方rpc库封装，兼容性也好）
+ easyrpc_discovery(https://github.com/zhuxiujia/easyrpc_discovery) （微服务发现）
+ GoMybatis 
+ Mysql5.7
+ redis

* 可选的库
+ GoFastdfs(分布式文件存储-引用其他作者的,你也可以选择其他的例如 阿里云oss) 
+ statikFs(静态文件编译，用于打包外部配置文件，mapper xml文件 进程序（实际上就是生成一串string）-引用其他作者的) 

## 编译打包
+ DockerFile镜像打包
+ jinkins自动打包（仅提供步骤脚本）

### 包管理
+ mod (golang官方推荐)

### DDD（领域驱动设计）分层:
+ app/ app/admin接口层，也可以看做网关鉴权层（基于JWT token（使用BCryptPasswordEncoder同时兼容java和go）配合easy_mvc过滤器）
+ core/服务层,也是核心业务层
+ common 公共库（Domain model,Rpc model,VO model）
+ consul/存放下载好的consul可执行文件
+ controller/网关层
+ database/dao数据库层
+ db/ 存放mysql表结构

### 项目接口分层
+ app(提供app接口)
+ admin(后台接口)



