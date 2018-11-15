# 服务计算学习之路4——用beego框架构建cloudgo简单应用


## 选择beego框架
+ 原因？
  + beego基于MVC架构
  + beego相对于其它框架([最好的6个go Web框架](https://blog.csdn.net/dev_csdn/article/details/78740990))，它的开发文档是中文的。

+ beego的安装
  + beego[官网](https://beego.me)介绍很详细，不过感觉beego官网不太稳定，有时会崩。
  + 使用命令`go get github.com/astaxie/beego`但我在centos7下要用`go get  -u github.com/astaxie/beego`才行
  + 如果要用beego做复杂点应用建议安装beego官方推荐的bee工具：`go get github.com/beego/bee`，但本次作业还没有用到。

### 代码实现
beego把功能分成很多模块，继承后重写一些方法就ok了，由于代码很短就直接贴上来了

```go
package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller //这相当于继承beego里的Controller类
}

func (controller *MainController) Get() { //重写Get方法
	username := controller.Ctx.Input.Param(":name")      //获取路由信息
	controller.Ctx.WriteString("Helloword! " + username) // 没用beego的模板，直接往网页写东西
}

func main() {
	beego.Router("/cloudgo/:name", &MainController{}) //设置路由，传入controller处理函数
	beego.Run(":9000")                                //在9000端口上运行
}
```

### curl测试
![](https://img-blog.csdnimg.cn/20181115174709952.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2tlMTk1MDUyMzQ5MQ==,size_16,color_FFFFFF,t_70)
分别是开启端口监听前和开启后curl测试的结果

### 压力测试
+ Centos7要先安装ab

  ```
  yum -y install httpd-tools
  ```

+ ab测试：

  ![](https://img-blog.csdnimg.cn/20181115175817878.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2tlMTk1MDUyMzQ5MQ==,size_16,color_FFFFFF,t_70)
  ![](https://img-blog.csdnimg.cn/20181115175829995.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2tlMTk1MDUyMzQ5MQ==,size_16,color_FFFFFF,t_70)

+ 参数解析：
   -n  即requests，用于指定压力测试总共的执行次数。
   -c  即concurrency，用于指定的并发数。
   其它还有很多参数，这次没用到就先别说了。

### 代码on Github
+ [代码传送门](https://github.com/kesongyue/ServiceComputer/tree/master/Homework3-cloudgo)
