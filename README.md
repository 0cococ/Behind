![xx](assets/banner.png)

### 声明：本库只允许非商业使用和学习

# 原生Go通用后端模板 · Behind

> The only people who have anything to fear from free software are those whose products are worth even less. 
>
> <p align="right">——David Emery</p>

Behind，是一款Go通用后端模板框架可以自己进行扩展。

![](https://img.shields.io/badge/language-java-brightgreen.svg)

## 作者的碎碎念

无聊的时候写的练手项目。


## 找我

QQ:3560000009

## 支持

注册接口：
1. 判断用户是否注册
2. 判断手机号长度(小于11位直接return)
3. 判断用户名是否nil(nil则随机生成一个10位随机名)
4. 判断密码长度(小于6位直接return)

## 如何使用

### Step 1.初始化：

额，下载直接运行就行



### 相关API

```Go
/api/auth/register(注册账户)
参数1：name(用户名)
参数2：telephone(手机号)
参数3：password(密码)
```

更多其他操作看api.go文件大概就知道了。




## 如何参与开发？

### 应用分2个模块

- app模块，用户操作与UI模块
- core模块，此模块为MagicHands的核心模块，也就是Sdk。

如需要参与开发请直接pr就可以了，相关教程请Google或者看 [如何在 GitHub 提交第一个 pull request](https://chinese.freecodecamp.org/news/how-to-make-your-first-pull-request-on-github/)

### PR须知

1. 中英文说明都可以，但是一定要详细说明问题
2. 请遵从原项目的代码风格、设计模式，请勿个性化。
3. PR不分大小，有问题随时欢迎提交。

## 计划

 - opencv实现图色识别(预计一月左右)
 - root权限实现自动化(预计半年左右)
 - 将要实现自动化的app加载到我们app中实现同一个进程的自动化(预计半年左右)
 - 利用esp32蓝牙hid协议实现自动化(预计半年左右)
 - 利用android的adb提权原理实现免插数据线实现adb自动化(预计半年左右)

## 感谢

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)

### License

> ```
> MIT License
> ```

>Copyright (c) 2023 pingan

>Permission is hereby granted, free of charge, to any person obtaining a copy
>of this software and associated documentation files (the "Software"), to deal
>in the Software without restriction, including without limitation the rights
>to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
>copies of the Software, and to permit persons to whom the Software is
>furnished to do so, subject to the following conditions:

>The above copyright notice and this permission notice shall be included in all
>copies or substantial portions of the Software.

>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
>IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
>FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
>AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
>LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
>OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
>SOFTWARE.
