# CHATROOM
**TCP实现的一个聊天室**

# 项目结构
```
chatroom/
├── client          客户端
│    ├── main       界面入口
│    ├── model
│    ├── process    client端的收发消息
│    └── utils      工具包/读取数据并验证
│ 
├── common         
│  ├── message   消息类型 
│  └── user     公共的结构体    
│ 
├── sever    
│    ├── main      登录/注册功能
│    ├── model     验证数据库中用户信息
│    ├── process   群发消息/上线提醒
└────└── utils
```
**build sever包创建服务器，并运行client 模拟用户登录连接**
