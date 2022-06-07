# 字节青训营项目

<a href="https://github.com/Mr-Jacks520" target="https://github.com/Mr-Jacks520"><img src="https://s2.loli.net/2022/06/07/z9jE61fxHBPich7.png" ></a>![手指.png](https://s2.loli.net/2022/06/07/4nOs3K85TfLUq7a.png)

## 项目结构：

**参考：[Gin-Vue-Admin](https://github.com/flipped-aurora/gin-vue-admin.git)**

```shell
├── api
├── config
├── core
├── global
├── initialize
├── middleware
├── model
│   ├── request
│   └── response
├── resource
├── router
├── service
└── utils
```

| 文件夹          | 说明             | 描述                                     |
|--------------|----------------|----------------------------------------|
| `api`        | 接口/controller层 | 相当于controller层这里参考的下面那个项目的架构,比较复杂      |
| `config`     | 配置包            | config.yaml对应的配置结构体                    |
| `core`       | 核心文件           | 核心组件(zap, viper, server)的初始化           |
| `global`     | 全局对象           | 全局对象                                   |
| `initialize` | 初始化            | router,redis,gorm,validator, timer的初始化 |
| `middleware` | 中间件层           | 用于存放 `gin` 中间件代码                       |
| `model`      | 模型层            | 模型对应数据表                                |
| `--request`  | 入参结构体          | 接收前端发送到后端的数据。                          |
| `--response` | 出参结构体          | 返回给前端的数据结构体                            |
| `resource`   | 静态资源文件夹        | 负责存放静态文件                               |
| `router`     | 路由层            | 路由层                                    |
| `service`    | service层       | 存放业务逻辑问题                               |
| `utils`      | 工具包            | 工具函数封装                                 |

## 接口完成：

- [x] 基础接口
- [x] 拓展接口一
- [x] 拓展接口二

## 快速开始：

1. 配置服务运行端口以及数据库

   ~~~yaml
   # config.yaml
   # system configuration
   system:
     env: 'develop'
     addr: 8080	#<<<<<<<<<<<<<<<<<设置服务运行端口
     db-type: 'mysql'
     oss-type: 'local'	#<<<<<<<<<<<<<<<<<本地功能还未实现,若有OSS服务可替换例如aliyun-oss,huawei-oss等
     use-redis: false
     use-multipoint: false
     # IP限制次数 一个小时15000次
     iplimit-count: 15000
     #  IP限制一个小时
     iplimit-time: 3600
     
   # mysql connect configuration
   mysql:
     path: 'localhost'	#<<<<<<<<<<<<<<<<<设置MYSQL地址
     port: '3306'	#<<<<<<<<<<<<<<<<<设置MySQL运行端口
     db-name: 'douyin'	#<<<<<<<<<<<<<<<<<设置MySQL数据库名，必须预先建好
     config: 'charset=utf8&parseTime=True&loc=Local'
     username: '数据库用户名'	#<<<<<<<<<<<<<<<<<填充
     password: '数据库用户密码'	#<<<<<<<<<<<<<<<<<填充
     max-idle-conns: 10
     max-open-conns: 100
     log-mode: "info"
     log-zap: false
   
   # aliyun oss configuration
   aliyun-oss:
     endpoint: ''
     access-key-id: ''
     access-key-secret: ''
     bucket-name: ''
     bucket-url: ''
     base-path: ''
   ~~~

2. 确认加载配置文件路径

   ~~~go
   // utils/constant.go
   package utils
   
   const (
   	ConfigFile = "config.yaml"
   )
   ~~~

3. 运行服务

   ~~~shell
   go run ./main.go
   ~~~

## 项目展示：

<a href="https://qingyin-video.oss-cn-chengdu.aliyuncs.com/%E6%BC%94%E7%A4%BA%E8%A7%86%E9%A2%91.mp4" target="_blank"><img src="https://s2.loli.net/2022/06/07/onXAjk7NvfMKY8w.png" >项目演示地址（每天9:00-11:00开放）</a>

**如需体验请联系[Mr-Jacks520](https://github.com/Mr-Jacks520)**



![debug.png](https://s2.loli.net/2022/06/03/ALIwj9O4cRXbDsZ.png)


# 6月4日基础接口开发完毕(BUG另说)

![视频Feed.jpg](https://s2.loli.net/2022/06/04/aKgMRbBjx2Wirs3.jpg)

![作品列表.jpg](https://s2.loli.net/2022/06/04/BjxFDNLvt9dZV27.jpg)

![上传.jpg](https://s2.loli.net/2022/06/04/YvW9dpT2R8Oq5Fr.jpg)

![发布列表.jpg](https://s2.loli.net/2022/06/04/8tSIZV73xPHMQRd.jpg)

