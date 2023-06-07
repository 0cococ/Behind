

### 声明：本库思路来自于'b站:六个六科技'

# 原生Go通用后端模板 · Behind

> The only people who have anything to fear from free software are those whose products are worth even less. 
>
> <p align="right">——David Emery</p>

Behind，是一款Go通用后端模板框架可以自己进行扩展。

![](https://img.shields.io/badge/language-java-brightgreen.svg)

## 作者的碎碎念

本库为初学项目,我会努力把本库完善,本库重要代码都加上了注释,以便于大家学习.


## 找我

QQ:3560000009

## 支持

注册接口：
1. 判断用户是否注册
2. 判断手机号长度(小于11位直接return)
3. 判断用户名是否nil(nil则随机生成一个10位随机名)
4. 判断密码长度(小于6位直接return)
5. 创建用户信息写入到mysql

登录接口:
1. 判断手机号长度(小于11位直接return)
2. 判断密码长度(小于6位直接return)
3. 判断手机号是否注册
4. 判断密码是否正确
5. 进行token返回

查询接口:
1.token通过jwt校验完成返回对应用户信息


## 如何使用

### Step 1.初始化：

额，下载直接运行就行



### 相关API

```Go
注册接口：
/api/auth/register(注册账户)
参数1：name(用户名)
参数2：telephone(手机号)
参数3：password(密码)
```

更多其他操作看api.go文件大概就知道了。



## 计划

 - vue界面(预计一月左右)
 - 百万并发(预计二月左右)

## 感谢

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)

### License

> ```
> 没有许可证
> ```
