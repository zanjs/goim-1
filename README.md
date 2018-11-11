### 1 主要功能
支持单用户多设备同时在线，支持好友之间聊天，群组聊天。
### 2 所用技术
golang+mysql,用到的框架有gin(对gin进行了简单的封装)，消息发送逻辑使用写扩散。
### 3 主要逻辑
client: 客户端
connect:连接层
logic:逻辑层
mysql:存储层
#### 登录
![](leanote://file/getImage?fileId=5be818f7a0c1f52252000000)
#### 单发
![](leanote://file/getImage?fileId=5be819a0a0c1f52252000001)
#### 群发
![](leanote://file/getImage?fileId=5be819d3a0c1f52252000002)
