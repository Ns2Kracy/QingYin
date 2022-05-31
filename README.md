# 项目结构

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

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`     | 接口/controller层             | 相当于controller层这里参考的下面那个项目的架构,比较复杂 |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `resource`   | 静态资源文件夹          | 负责存放静态文件                |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `utils`      | 工具包                  | 工具函数封装            |

## 项目架构参考

[相关项目](https://github.com/flipped-aurora/gin-vue-admin.git)

## 5.31-6.1测试访问接口结果：

![image-20220601002408274](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20220601002408274.png)

**并未实现业务逻辑，只是进行了伪数据测试**

