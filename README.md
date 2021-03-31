# DeFi 挖矿自动复投脚本

源自湾区SWB https://imwsb.org


### 目前仅支持 Heco 链的 [MDEX](https://mdex.com/#/liquidity) 及 BSC 链的 [pancakeswap](https://pancakeswap.finance/farms) LP对复投挖矿


## 使用方法

在对应的平台完成交易对配对并质押挖矿

#### 1. 访问 [release](https://github.com/MycoinTool/revinest/releases) 下载对应平台的软件包 (支持 Windwos/Mac OSX/Linux 系统)
#### 2. 解压打开并执行 reinvest 
#### 3. 选择相应的挖矿平台，回车
#### 4. 填入要挖的ID，回车
<img width="500" alt="readme-2" src="https://user-images.githubusercontent.com/81501838/112823170-86377e80-90bb-11eb-9114-89114d29f8e3.png">

#### 5. 在 Private Key: 后贴入钱包私钥（如下图）
<img width="500" alt="readme-2" src="https://user-images.githubusercontent.com/81501838/112820573-5044cb00-90b8-11eb-9bf4-c12a5fa3545d.jpg">

#### 6. 填入复投频率 （单位/分钟），回车后开始自动复投。请保持电脑不会进入睡眠模式
<img width="500" alt="readme-3" src="https://user-images.githubusercontent.com/81501838/112820842-9f8afb80-90b8-11eb-9477-c57ae8d12e83.png">



## 收益和 GAS 费参考
经过两天的稳定性测试，数据如下：

> 2021/03/28 05:44:23 开始投入 0.927 ETH + 相应的 MDX 配对 ETH/MDX LP，每2小时复投一次
> 
> 2021-03-29 15:51:28 收获 8.671708 MDX，消耗 GAS 费 0.022166928 HT，约合RMB 1.96元



## 免责声明
1. 脚本为 golang 编写，代码开源可自行修改、编译、运行
2. 脚本需要你提供私钥才能正常运行，请谨慎使用



## 其它
我们的网站，专注区块链小工具  https://mycointool.com/zh/  欢迎访问!

其它矿池的支持开发中...
