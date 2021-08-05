# go语言调用c语言的方法

首先需要生成c的动态链接库或静态链接库，如果已经存在则参考__注意__部分进行配置。

如果代码是C++，则需要先使用C进行包装一次。

代码如下：
demo.h
```
class demo{
public:
    demo();
    ~demo();
    void hello();
}
```

demo.cpp
```
#include<stdio.h>
#include "demo.h"
demo::demo(){

}
demo::~demo(){
}
void demo:hello(){
    printf('hello world\n');
}
```
C语言封装，demo_c.h
```
#ifdef __cplusplus
extern "C" {
#endif
    void hello();
#ifdef __cplusplus
}
#endif
```
demo_c.c
```
#include "demo.h"
#include "demo_c.h"

void hello() {
    demo* demoObj = new demo();
    demoObj->hello();
}
```
然后进行编译，在当前代码的目录下，g++ -c demo.cpp，g++ -c demo_c.c，ar -crs libdemo.a demo.o demo_c.o

最后是go的代码，注意`import "C"` 和上面的注释行不能有空行，下面路径的demo是c的头文件，库的目录，-ldemo是动态链接库的名称：
```
package main

// #cgo CFLAGS: -I./demo
// #cgo LDFLAGS: -lstdc++ -L./demo -ldemo
// #include "demo_c.h"
import "C"

func main(){
    C.hello()
}
```

### 注意
1. 在编译时，需要将动态链接库放在`ld`可扫描的路径，例如在linux下面，可以在`/etc/ld.so.conf.d/`增加对应的conf文件，在conf文件中配置`so`所在的文件夹；
2. 在运行时，需要将动态链接库放在运行扫描的路径，例如在linux下面，可以设置环境变量`LD_LIBRARY_PATH`，如：`export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:{the path of the lib}`；s