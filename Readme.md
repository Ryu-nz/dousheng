#### 基本介绍

使用Go实现短视频播放软件的后端接口，基于传统MVC架构，功能包含视频Feed流观看视频、视频投稿、个人信息模块、点赞与评论模块及关注与粉丝关系模块

根据提前定义好的各个功能对应的接口，以及提供极简版抖音的安装包来支持功能验证和调试

接口文档：[抖音极简版 (apifox.cn)](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)

需求文档及Demo软件：[极简抖音App使用说明 - 青训营版 - 飞书云文档 (feishu.cn)](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)

#### 项目架构

<img src="../Images/Readme/image-20220719150747167.png" alt="image-20220719150747167" style="zoom: 50%;" />

#### 技术选型

- Gin：Web框架
- Viper：配置管理
- Zap：日志管理
- JWT：用户鉴权
- GORM：ORM框架
- Mysql：持久层数据库

#### 功能实现

- 视频：视频推送、视频投稿、发布列表
- 用户：用户注册、用户登录、用户信息
- 点赞：点赞操作、点赞列表
- 评论：评论操作、评论列表
- 关注：关注操作、关注列表、粉丝列表

#### 目录结构

**.**                                                                                                                                                                                         

├── **config**                                                                                                                                                                                

│  └── config.go                                                                                                                                                                         

├── **controller**                                                                                                                                                                            

│  ├── basic.go                                                                                                                                                                          

│  ├── comment.go                                                                                                                                                                        

│  ├── favorite.go                                                                                                                                                                       

│  ├── feed.go                                                                                                                                                                           

│  ├── jwt.go                                                                                                                                                                            

│  ├── publish.go                                                                                                                                                                        

│  ├── request.go                                                                                                                                                                        

│  ├── response.go                                                                                                                                                                       

│  └── user.go                                                                                                                                                                           

├── **global**                                                                                                                                                                                

│  └── globalVar.go                                                                                                                                                                      

├── go.mod                                                                                                                                                                                

├── go.sum                                                                                                                                                                                

├── **Init**                                                                                                                                                                                  

│  ├── config.go                                                                                                                                                                         

│  ├── logger.go                                                                                                                                                                         

│  ├── minio.go                                                                                                                                                                          

│  ├── mysql.go                                                                                                                                                                          

│  └── router.go                                                                                                                                                                         

├── **middlewares**                                                                                                                                                                           

│  └── log.go                                                                                                                                                                            

├── readme.md                                                                                                                                                                             

├── **router**                                                                                                                                                                                

│  ├── comment.go                                                                                                                                                                        

│  ├── favorite.go                                                                                                                                                                       

│  ├── feed.go                                                                                                                                                                           

│  ├── publish.go                                                                                                                                                                        

│  └── user.go                                                                                                                                                                           

├── server.go                                                                                                                                                                             

├── settings.yaml                                                                                                                                                                         

└── **utils**                                                                                                                                                                                 

​    ├── getTime.go                                                                                                                                                                        

​    ├── ip.go                                                                                                                                                                             、

​    └── minio.go         