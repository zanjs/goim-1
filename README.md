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
[![00d7e21cccc9050e.jpg](http://www.wailian.work/images/2018/11/12/00d7e21cccc9050e.jpg)](http://www.wailian.work/image/BVGZkp)
#### 单发
[![3496be2f9ee9d33e.jpg](http://www.wailian.work/images/2018/11/12/3496be2f9ee9d33e.jpg)](http://www.wailian.work/image/BVGV24)
#### 群发
[![7ee3ada2baf1dec0.jpg](http://www.wailian.work/images/2018/11/12/7ee3ada2baf1dec0.jpg)](http://www.wailian.work/image/BVGtLc)
