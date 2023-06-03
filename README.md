

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



## 计划

 - 登录接口(预计一月左右)
 - 造一个vue界面(预计一月左右)
 - 中间件处理请求(预计一月左右)
 - 百万并发抢购功能(预计二月左右)

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
