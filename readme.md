# 一个golang使用的banner项目，
本项目基于[https://github.com/dimiro1/banner](https://github.com/dimiro1/banner) 而创建的
原项目有以下特点:
- 支持文本文件
- 支持部分ansi颜色
- 不能把banner编译到可执行文件中

本项目有以下特点
- 支持文本文件
- 支持部分ansi颜色
- ***能把banner编译到可执行文件中***

# 使用方法：
> 打开go_banner.go文件，修改文件中定义的GO_BANNER为你自己需要显示的banner，接下来保存
> 在你的应用中使用如下代码在程序入口文件导入即可，如本例子中main那样
```
	import _ "github.com/aloxc/gobanner"
```
**特别说明下，如果你是在ide中运行可能显示效果不如意，请切换到shell或者cmd中执行你的程序**

![gobanner](https://github.com/aloxc/gobanner/blob/master/assert/gobanner.png)

可以通过[http://patorjk.com/software/taag/#p=display&f=Big&t=Banner](http://patorjk.com/software/taag/#p=display&f=Big&t=Banner)来制作你自己的banner

可以在docker试验下
```
docker pull aloxc/gobanner
docker run --name gobanner aloxc/gobanner
```
