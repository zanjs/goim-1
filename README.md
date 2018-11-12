### 1 主要功能
支持单用户多设备同时在线，支持好友之间聊天，群组聊天。
### 2 所用技术
golang+mysql,用到的框架有gin(对gin进行了简单的封装)，消息发送逻辑使用写扩散。
### 3 拆包粘包以及通信协议
拆包粘包使用了自己写的一个拆包粘包算法，主要也是采用了TLV的思想，把一个消息割分成三段，第一段用两个字节表示消息的类型，第二段也是用两个字节来表示消息的长度，第三段则是消息的实际内容。详情可以查看：  
https://www.jianshu.com/p/e7c016efb09d  
通信协议使用Google的Protocol buffers作为通信协议。
### 4 消息唯一id生成
其思想借鉴了Leaf——美团点评分布式ID生成系统的Leaf-segment数据库双buffer优化方案，其实他的核心思想是，每次从数据库拿取一个号段，用完了，再去数据库拿，当用尽去数据库拿的时候，会有一小会的阻塞，对这一情况做了一些优化。

刚开始实现的时候，和美团的方案一样，利用两个buffer，Leaf服务内部有两个号段缓存区segment。当前号段已下发10%时，如果下一个号段未更新，则另启一个更新线程去更新下一个号段。当前号段全部下发完后，如果下个号段准备好了则切换到下个号段为当前segment接着下发，循环往复。

最后想了想，其实没必要这么复杂，用一个channal,一边起一个goroutine，先从数据库拿取一个号段，然后生成id放到channel里面，如果号段用尽，再从数据库里面取，如此往复，当channel里面满时，goroutine会阻塞。一边用的时候从里面拿就行。
https://www.jianshu.com/p/9295e1babf37
### 3 主要逻辑
client: 客户端  
connect:连接层  
logic:逻辑层  
mysql:存储层  
#### 登录
[![3496be2f9ee9d33e.jpg](http://www.wailian.work/images/2018/11/12/3496be2f9ee9d33e.jpg)](http://www.wailian.work/image/BVGV24)
#### 单发
[![00d7e21cccc9050e.jpg](http://www.wailian.work/images/2018/11/12/00d7e21cccc9050e.jpg)](http://www.wailian.work/image/BVGZkp)
#### 群发
[![7ee3ada2baf1dec0.jpg](http://www.wailian.work/images/2018/11/12/7ee3ada2baf1dec0.jpg)](http://www.wailian.work/image/BVGtLc)
### 日志
![9f644dcd04b20287.jpg](http://www.wailian.work/images/2018/11/12/9f644dcd04b20287.jpg)