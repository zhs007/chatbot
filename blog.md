# ChatBot Development Log

### 2019-09-12

这一版的ChatBot，将AppServ和ChatBot内核分离，而且支持同一个App的多个账号。  
后面ChatBot内核肯定也是能水平扩展的（多开）。

核心模块还是插件体系。  

核心协议还是``grpc``。