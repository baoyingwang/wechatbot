# wechatbot
本项目从https://github.com/djun/wechatbot fork而来
更改如下
1. 更改为支持openai on azure
2. 相关配置迁移到配置文件中

没有做其他改动，具体功能请直接到https://github.com/djun/wechatbot查看

# 注册openai on azure
参考:https://learn.microsoft.com/en-us/azure/cognitive-services/openai/chatgpt-quickstart?tabs=bash&pivots=rest-api
1. 对于微软员工注意使用公司邮箱
1. 已经注册openai on azure 3.5之后，进一步申请chatgpt4 - 这个[链接](https://aka.ms/oai/get-gpt4)申请


# 安装使用
````
# 获取项目
git clone https://github.com/baoyingwang/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

启动前需替换config中的api_key, endpoint等

# 启动项目
go run main.go

# 按照提示使用微信小号（别用大号，被封了咋办）
加下来，任何发送给这个微信号的消息，将被转发给openai
譬如
 + 群聊@回复
 + 私聊回复
 + 自动通过回复（https://github.com/djun/wechatbot中提到这个功能，还不知道如何使用）

