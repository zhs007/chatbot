# chatbot

### 概要

golang的chatbot框架。  
特性如下：

- 支持多聊天工具，现在支持telegram和QQ（coolq）。
- 分布式弹性架构，主进程和各聊天工具对接进程彻底分离，用户管理进程等也可以独立配置，各服务节点之间通过grpc通信，不限开发语言。
- 插件式结构，可以非常方便的通过插件扩展chatbot功能。
- 支持command扩展，可以方便的处理各种指令参数。
- 支持文件处理器扩展，可以方便的添加新的文件处理代码。
- 支持多语言。
- 支持配置文件响应指令。
- 开发简单，易上手；大部分节点都无需二次开发，即可直接使用。
- 调试友好，自带debug模式，直接通过聊天，就可以得到足够的调试信息。
- 运维友好，所有节点都直接通过Docker部署，提供第三方服务部署脚本。

计划添加的功能：

- 跨聊天工具的统一账号管理；通过token绑定账号。
- 支持自然语言聊天。

### 使用

暂略

### 开发 Core Node

暂略

### 接入 telegram

接入telegram非常简单，您可以直接使用官方镜像。  

安装或更新官方镜像：

``` sh
docker pull zerrozhao/telegrambot
```

然后需要一个telegram的配置文件，格式如下：

``` yaml
telegramtoken: 1234567

token: 1234567
servaddr: 127.0.0.1:123
username: ada
```

这个文件需要放在当前目录的cfg下，然后我们可以通过下面的脚本来启动telegrambot节点。

``` sh
docker container stop telegrambot
docker container rm telegrambot
docker run -d \
    --name telegrambot \
    --link adachatbot:adachatbot \
    -v $PWD/cfg:/app/telegrambot/cfg \
    -v $PWD/logs:/app/telegrambot/logs \
    zerrozhao/telegrambot
```

这里没有直接用``--rm``，而是先 ``stop`` ，再 ``rm`` ，主要是方便查看错误日志。

### 接入 CoolQ

1. 注册QQ号，绑定手机号。
2. 用手机版的QQ登录手机号，开启账号锁定。
3. 使用coolq目录下的 init.sh 脚本初始化 coolq，这个脚本基本上可以不用修改。
4. 第一次启动 coolq，用 coolq目录下的 start.sh 脚本启动 coolq，注意参数，建议修改端口、密码、qq号等参数。
5. 用浏览器连接导出的9000端口，用VNC密码登录进去，启动CoolQ，第一次启动，会需要短信验证码，前面步骤1、2一定要执行，这种方式最省事了！
6. 到data/app/io.github.richardchien.coolqhttpapi/config/修改配置文件，qq号.ini。  
``` ini
serve_data_files = yes
use_ws = yes
ws_host = 0.0.0.0
ws_port = 6700
```
7. 重启coolq。

然后，才可以开始部署coolqbot。  
还是直接使用官方镜像吧。

``` sh
docker pull zerrozhao/coolqbot
```

修改coolqbot的配置文件。

``` yaml
ooolqtoken:
coolqsecret:
coolqservurl: ws://127.0.0.1:5700
coolqhttpservaddr: 

token: 1234567
servaddr: 127.0.0.1:123
username: ada

debug: true
preloaduserinfo: true
```

接下来，我们可以启动coolqbot了。

``` sh
docker container stop coolqbot
docker container rm coolqbot
docker run -d \
    --name coolqbot \
    --link adachatbot:adachatbot \
    --link cqhttp:cqhttp \
    -p 7234:7234 \
    -v $PWD/cfg:/app/coolqbot/cfg \
    -v $PWD/logs:/app/coolqbot/logs \
    zerrozhao/coolqbot
```

必要的参数要改，如果开在一台服务器上，建议用 ``--link``，可以省掉外网端口占用。



### 版本更新

##### v0.3

- 支持channel、group等。

##### v0.2

- 支持QQ。

##### v0.1

- 核心架构完成。
- 完成telegram节点。
- plugin、command、fileprocessor架构完成。
- 多语言框架完成。
- 第一个机器人，Ada的telegram完成，[这里](https://t.me/@ada_heyalgo_bot)。