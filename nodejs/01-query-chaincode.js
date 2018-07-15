/*
 * client.js
 * Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

var fs = require('fs');
var Fabric_Client = require('fabric-client');

//创建一个Client
Fabric_Client.newDefaultKeyValueStore({ path: '/tmp/xx/' }).then((state_store) => {
    client=new Fabric_Client();
    client.setStateStore(state_store)

    //设置用户信息    
    var userOpt = {
        username: 'Admin@member1.example.com',
        mspid: 'peers.member1.example.com',
        cryptoContent: { 
            privateKey: './msp/keystore/09dd09cf530d8f0fa6cb383b5b409ae8e895d32d31f75823f3bdb3c1f3ee180a_sk',
            signedCert: './msp/signcerts/Admin@member1.example.com-cert.pem'
        }
    }

    return client.createUser(userOpt)

}).then((user)=>{

    //设置要连接的Channel
    var channel = client.newChannel('mychannel');

    //设置要连接的Peer
    var peer = client.newPeer(
        'grpcs://peer0.member1.example.com:7051',
        {
            pem: fs.readFileSync('./tls/ca.crt', { encoding: 'utf8' }),
            clientKey: fs.readFileSync('./tls/client.key', { encoding: 'utf8' }),
            clientCert: fs.readFileSync('./tls/client.crt', { encoding: 'utf8' }),
            'ssl-target-name-override': 'peer0.member1.example.com'
        }
    );

    channel.addPeer(peer);

    //调用chaincode
    const request = {
        chaincodeId: 'mycc',   //chaincode名称
        fcn: 'query',          //调用的函数名
        args: ['key1']         //参数
    };

    // send the query proposal to the peer
    return channel.queryByChaincode(request);

}).then((response)=>{
    console.log('Response is', response.toString());
})
