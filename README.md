#Go-DDD领域驱动代码实践(Domain Driver Design)  
##一级目录结构
###interfaces(用户接口层):
主要存放用户接口层与前端交互、展现数据相关的代码  
###application(应用层):
主要存放应用层服务组合和编排相关的代码  
###Domain(领域层):
主要存放领域层核心业务逻辑相关的代码  
###Infrastructure(基础层):
主要存放基础资源服务相关的代码  
##各层目录结构
###用户接口层
Assembler:   
实现 DTO 与领域对象之间的相互转换和数据交换  
Dto:  
它是数据传输的载体，内部不存在任何业务逻辑，我们可以通过 DTO 把内部的领域对象与外界隔离  
Facade:   
提供较粗粒度的调用接口，将用户请求委派给一个或多个应用服务进行处理  
###应用层
Event:   
这层目录主要存放事件相关的代码   
Service:   
这层的服务是应用服务   
###领域层
Aggregate(聚合):   
它是聚合软件包的根目录，可以根据实际项目的聚合名称命名，比如订单聚合   
Entity（实体）:   
它存放聚合根、实体、值对象以及工厂模式（Factory）相关代码   
Event（事件）:   
它存放事件实体以及与事件活动相关的业务逻辑代码   
Service（领域服务）:   
它存放领域服务代码   
Repository（仓储）:   
它存放所在聚合的查询或持久化领域对象的代码，通常包括仓储接口和仓储实现方法,例子中的以内存memory为例，可以更换为mysql、mongo、redis等其它。
###基础层
Config： 
主要存放配置相关代码   
Util： 
主要存放平台、开发框架、消息、数据库、缓存、文件、总线、网关、第三方类库、通用算法等基础代码，你可以为不同的资源类别建立不同的子目录   
###代码树
+---applicaiton   
|   +---event   
|   |   +---publish   
|   |   \---subscribe   
|   \---service   
+---domain   
|   \---aggregate   
|       +---entity   
|       +---event   
|       +---repository   
|       |   \---memory  
|       \---service  
+---infrastructure   
|   +---config   
|   \---util   
|       +---api   
|       +---driver   
|       +---eventbus   
|       \---mq   
\---interface   
+---assembler   
+---dto  
\---facade   
##后记
写这个的原因是因为面试的时候被人问到怎么设计代码结构，用传统MVC模式写后被科普了DDD模式，   
回家研究了一下，参考 [DDD 实战课](url:https://zq99299.github.io/note-book2/ddd/#%E6%8E%A8%E8%8D%90%E9%98%85%E8%AF%BB])  写了demo以加深理解



