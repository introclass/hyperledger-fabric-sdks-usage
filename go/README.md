---
layout: default
title:  README
author: lijiaocn
createdate: 2018/07/16 10:45:00
changedate: 2018/07/16 17:05:02

---

* auto-gen TOC:
{:toc}

## 说明

Go语言的SDK还没有正式Release：[hyperledger/fabric-sdk-go][1]，这里使用的是:

	git checkout d489eba94c868f736015935e88da20faf9aa80d4

(@2018-07-16 10:46:41)

这里有一个使用示例：[Fabric CLI Sample Application][2]

	$ mkdir -p $GOPATH/src/github.com/securekey
	$ cd $GOPATH/src/github.com/securekey
	$ git clone https://github.com/securekey/fabric-examples.git

后面的代码主要是参考fabric-example中的实现。

## Client

github.com/hyperledger/fabric-sdk-go/pkg/client/中client的实现：

	/Users/lijiao/Work/Bin/gopath/src/github.com/hyperledger/fabric-sdk-go/pkg/client/
	▸ channel/
	▸ common/
	▸ event/
	▸ ledger/
	▸ msp/
	▸ resmgmt/

### Channel Client

channel子目录中的`example_test.go`给出了使用示例：

	/Users/lijiao/Work/Bin/gopath/src/github.com/hyperledger/fabric-sdk-go/pkg/client/
	▾ channel/
	  ▸ invoke/
	    api.go
	    api_test.go
	    chclient.go
	    chclient_test.go
	    example_test.go

演示了5个操作：

	▼ functions
	   +Example()
	   +ExampleClient_Execute()
	   +ExampleClient_InvokeHandler()
	   +ExampleClient_Query()
	   +ExampleClient_RegisterChaincodeEvent()
	   +ExampleNew()

Channel Client的定义如下，可以看到有`Execute`、`InvokeHandler`、`Query`等方法：

	▼+Client : struct
	    [fields]
	   -context : context.Channel
	   -eventService : fab.EventService
	   -greylist : *greylist.Filter
	   -membership : fab.ChannelMembership
	    [methods]
	   +Execute(request Request, options ) : Response, error
	   +InvokeHandler(handler invoke.Handler, request Request, options ) : Response, error
	   +Query(request Request, options ) : Response, error
	   +RegisterChaincodeEvent(chainCodeID string, eventFilter string) : fab.Registration, chan *fab.CCEvent, error
	   +UnregisterChaincodeEvent(registration fab.Registration)
	   -createReqContext(txnOpts *requestOptions) : reqContext.Context, reqContext.CancelFunc
	   -prepareHandlerContexts(reqCtx reqContext.Context, request Request, o requestOptions) : *invoke.RequestContext, *invoke.ClientContext, error
	   -prepareOptsFromOptions(ctx context.Client, options ) : requestOptions, error
	    [functions]
	   +New(channelProvider context.ChannelProvider, opts ) : *Client, error

Channel Client的创建过程比较复杂，原因是`New()`函数的参数构造比较麻烦。context.ChannelProvider是函数类型：

	type ChannelProvider func() (Channel, error)

它的返回值类型为：

	type Channel interface {
	    Client
	    ChannelService() fab.ChannelService
	    ChannelID() string
	}

继续把Channel展开：

	type Client fab.ClientContext
	
	    type ClientContext interface {
	        core.Providers
	        msp.Providers
	        Providers
	        msp.SigningIdentity
	    }
	
	        type Providers interface {
	            CryptoSuite() CryptoSuite
	            SigningManager() SigningManager
	        }
	        type Providers interface {
	            UserStore() UserStore
	            IdentityManagerProvider
	            IdentityConfig() IdentityConfig
	        }
	        type Providers interface {
	            LocalDiscoveryProvider() LocalDiscoveryProvider
	            ChannelProvider() ChannelProvider
	            InfraProvider() InfraProvider
	            EndpointConfig() EndpointConfig
	        }
	        type SigningIdentity interface {
	
	            // Extends Identity
	            Identity
	            ...
	        }
	
	type ChannelService interface {
	    Config() (ChannelConfig, error)
	    EventService(opts ...options.Opt) (EventService, error)
	    Membership() (ChannelMembership, error)
	    ChannelConfig() (ChannelCfg, error)
	    Transactor(reqCtx reqContext.Context) (Transactor, error)
	    Discovery() (DiscoveryService, error)
	    Selection() (SelectionService, error)
	}
	...

可以看到包含的内容非常非常多，上面还只是初步展开。

`github.com/hyperledger/fabric-sdk-go/pkg/fab/mocks`专门用来构造复杂的参数，`mockchannel.go`中构建了一个Channel变量，这样实在太麻烦。看看有没有其它方法。

	▾ hyperledger/
	  ▾ fabric-sdk-go/
	    ▸ internal/
	    ▾ pkg/
	      ▸ client/
	      ▸ common/
	      ▸ context/
	      ▸ core/
	      ▸ fab/
	      ▾ fabsdk/

在

	func New(configProvider core.ConfigProvider, opts ...Option) (*FabricSDK, error) {
		pkgSuite := defPkgSuite{}
		return fromPkgSuite(configProvider, &pkgSuite, opts...)
	}



## 参考

1. [hyperledger/fabric-sdk-go][1]
2. [Fabric CLI Sample Application][2]

[1]: https://github.com/hyperledger/fabric-sdk-go  "hyperledger/fabric-sdk-go"
[2]: https://github.com/securekey/fabric-examples/tree/master/fabric-cli "Fabric CLI Sample Application"
