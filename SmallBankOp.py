#encoding=utf-8
"""
NEU bLockchain first man
@autho=wanggang
功能 描述了Smallbank的所有操作
"""

#solidity 版本
"""
pragma solidity ^0.4.0;
 
contract SmallBank {
    
    //uint constant MAX_ACCOUNT = 10000;
    //uint constant BALANCE = 10000;
    //bytes20 constant accountTab = "account";
    //bytes20 constant savingTab = "saving";
    //bytes20 constant checkingTab = "checking";
    
    // 活期储蓄账户
    mapping(string=>uint) savingStore;
    // 支票账户
    mapping(string=>uint) checkingStore;
 
    // 将支票账户并到储蓄账户
    function almagate(string arg0, string arg1) public {
       uint bal1 = savingStore[arg0];
       uint bal2 = checkingStore[arg1];
       
       checkingStore[arg0] = 0;
       savingStore[arg1] = bal1 + bal2;
    }
 
    // 得到账户总余额
    function getBalance(string arg0) public constant returns (uint balance) {
        uint bal1 = savingStore[arg0];
        uint bal2 = checkingStore[arg0];
        
        balance = bal1 + bal2;
        return balance;
    }
    
    // 更新支票账户余额
    function updateBalance(string arg0, uint arg1) public {
        uint bal1 = checkingStore[arg0];
        uint bal2 = arg1;
        
        checkingStore[arg0] = bal1 + bal2;
    }
    
    // 更新储蓄账户余额
    function updateSaving(string arg0, uint arg1) public {
        uint bal1 = savingStore[arg0];
        uint bal2 = arg1;
        
        savingStore[arg0] = bal1 + bal2;
    }
    // arg0向arg1支付arg2元
    function sendPayment(string arg0, string arg1, uint arg2) public {
        uint bal1 = checkingStore[arg0];
        uint bal2 = checkingStore[arg1];
        uint amount = arg2;
        
        bal1 -= amount;
        bal2 += amount;
        
        checkingStore[arg0] = bal1;
        checkingStore[arg1] = bal2;
    }
    // 写支票
    function writeCheck(string arg0, uint arg1) public {
        uint bal1 = checkingStore[arg0];
        uint bal2 = savingStore[arg0];
        uint amount = arg1;
        
        if (amount < bal1 + bal2) {
            checkingStore[arg0] = bal1 - amount - 1;
        } 
        else {
            checkingStore[arg0] = bal1 - amount;
        }
    }
}
"""
import sys
"""
命名规则
savingStore  活期账户  应该是字典（key：value）
checkingStore  支票账户  是个字典（key1：value1）
方法
1.将支票账户并到储蓄账户  字典中支票账户转换到字典中的储蓄账户，支票账户的账号置为0，存储账户的余额=自身余额+支票余额
2.得到账户总余额  总余额=支票账户+存储账户
3.更新支票账户余额
4.储蓄账户余额
5.转账
6.写支票
"""
print("Small Bank")
savingStore = {"wanggang": 1000, "wangzengyan": 500,"wangluange":10000}
checkingStore = {"gang": 300, "yan": 50,"wanggang": 10}
def ChecktoSave(arg0,arg1):
    savingStore[arg1] += checkingStore[arg0]
    checkingStore[arg0]=0
    return checkingStore[arg0],savingStore[arg1]
def getBalance(arg0):
    checkbalance=checkingStore[arg0]
    savebalance=savingStore[arg0]
    allbalance=checkbalance+savebalance
    return allbalance
def updatecheckaccount(arg0,money):
    checkingStore[arg0]+=money
    return checkingStore[arg0]
def updatesaveaccount(arg0,money):
    savingStore[arg0] += money
    return savingStore[arg0]
def sendPayment(arg0,arg1,money):
    if checkingStore[arg0]>money:
        Frombal = checkingStore[arg0]-money
        Tobal = checkingStore[arg1]+money
        print("转账成功")
        return Frombal,Tobal
    else:
        print("支票额度不足")
        sys.exit()#跳出程序
def writeCheck(arg0, money): #个人理解给别人写支票
        bal1 = checkingStore[arg0];
        bal2 = savingStore[arg0];
        amount =money;
        if (amount <=bal1+bal2) :
            bal2 -= (amount - bal1)
            bal1 =0;
            return bal1, bal2
        elif (amount > bal1+bal2):
            print("转账不成功")
            sys.exit()
        else:
            checkingStore[arg0] = bal1 - amount;
            return checkingStore[arg0],savingStore[0]


if __name__ == '__main__':
    #print(ChecktoSave("gang","wanggang"))  #输出(0,1300) 支票转存储账户
    #print(getBalance("wanggang")) #输出总金额
    #print(updatecheckaccount("gang",300))
    #print(updatesaveaccount("wangzengyan", 300))
    #print(sendPayment("gang","yan",10000))
    print(writeCheck("wanggang",  11))


