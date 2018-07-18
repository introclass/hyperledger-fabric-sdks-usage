---
layout: default
title:  README
author: lijiaocn
createdate: 2018/07/15 14:14:00
changedate: 2018/07/18 10:24:57

---

* auto-gen TOC:
{:toc}

## 适用版本

release-1.1

`注意将msp和tls，以及代码中的url，替换成你自己的。`

使用说明见：[《超级账本HyperLedger：Fabric Node.js SDK的使用》](http://www.lijiaocn.com/%E7%BC%96%E7%A8%8B/2018/04/25/hyperledger-fabric-sdk-nodejs.html)

使用过程中遇到的一些问题，记录在这里：[Fabric的Node.js SDK使用时遇到的问题](https://www.lijiaocn.com/问题/2018/07/15/hyperledger-fabric-nodejs-problem.html)

## 准备合适的node版本

当前支持的node版本是v8.9.0~v9.0，v9.0以上版本不支持(2018-07-15 14:11:01)。

在mac上可以用brew安装node8：

	$ brew install node@8
	$ echo 'export PATH="/usr/local/opt/node@8/bin:$PATH"' >> ~/.bash_profile
	$ source ~/.bash_profile
	$ node --version
	v8.11.3

或者直接下载安装：[nodejs download][1]

## 用npm管理依赖包

创建文件package.json：

	$ cat package.json
	{
	    "dependencies": {
	        "fabric-ca-client": "1.1.2",
	        "fabric-client": "1.1.2",
	        "grpc": "^1.6.0"
	    },
	    "author": "Anthony O'Dowd",
	    "license": "Apache-2.0",
	    "keywords": [
	        "Hyperledger",
	        "Fabric",
	        "Car",
	        "Sample",
	        "Application"
	    ]
	}

用npm安装依赖包：

	npm config set registry https://registry.npm.taobao.org  (设置淘宝提供的镜像源)
	npm install

## Example

调用mychannel中的mycc合约的query接口，参数为`key`:

	node ./01-query-chaincode.js

## 参考

1. [nodejs download][1]

[1]: https://nodejs.org/en/  "nodejs download" 
