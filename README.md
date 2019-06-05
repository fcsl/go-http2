# golang http2预研
1. golang中只要给HTTP服务配置TLS证书，则自动支持HTTP协议，如代码main.go所示
2. h2c.go是非加密版的HTTP2服务器，client.go是相应的客户端。

#HTTP2的特性

- 多路复用——多个请求响应可以在一个TCP连接中交叉传输，可解决HTTP1.1中的"队头阻塞"问题
- 请求头压缩
- 请求之间可以设置依赖关系或优先级——但是是否起作用还得看服务器的实现
- 流量控制——类似TCP的流量控制
- 服务端数据推送

# 调试

由于HTTP2一般都是TLS加密的，因此抓包调试比普通的HTTP请求麻烦，下面介绍Wireshark抓包：

1. 设置环境变量SSLKEYLOGFILE=**~/sslkey.log**
2. Wireshark的Preferences > Protocols > SSL，在(Pre)-Master-Secret log filename 中写入**~/sslkey.log**
3. 用chrome请求对应的URL(需要打开开发者工具)

# 注意事项

HTTP2一般都要基于TLS(TLS1.2及以上)，因为主流的浏览器、curl、SpringBoot等都不支持非HTTPS的HTTP2服务