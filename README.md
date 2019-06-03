#### http redirect 重定向
```
客户端client 请求server1,server1将请求重定向给server2，server2再将请求重定向给server3，server3相应服务.
     +--------------+           +----------------+        
     |              |           |                |        
     |   client     +<--------->+  server1       |        
     |              |           |                |        
     +--------------+           +----------------+   
            |
            |                   +----------------+        
            |                   |                |        
            ------------------->+  server2       |        
                                |                |        
                                +----------------+   

```

<img src="./static/http_redirect.png">



#### http 反向代理
```
客户端client 请求server1,server1将请求重定向给server2，server2再将请求重定向给server3，server3相应服务.

     +--------------+           +----------------+         +----------------+
     |              |           |                |         |                |
     |   client     +---------->+ forward proxy  +--------^+ server1/2/3    |
     |              |           |                |         |                |
     +--------------+           +----------------+         +----^-----------+
                                                                |
                                                                |
    +---------------+                                           |
    |               |                                           |
    |   client      +<------------------------------------------+
    |               |
    +---------------+
```

<img src="./static/http_proxy.png">




