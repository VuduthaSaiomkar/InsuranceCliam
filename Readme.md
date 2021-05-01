# Insurance Claim Through Blockchain 

## building Blockchain network to avoid duplicate claims on Insurance sector 

#### Follow below steps to create a network 

#### Project Description 
 > I have taken hyperledger fabric as Blockchain medium to track the insurance cliams between mulitple organizations ,why hyperledger fabric ? because here we are trageting transaction data not any crypto-tokens are required.

- First identifying the required Organization needs to particpate in blockchain network 
- Below are Following list of organizations
    - Workshop
    - Police 
    - Insurance
- Generate required files for creating blockchain network 
    - Config.yaml
    - crypto-config.yaml
    - docker-compose.yaml

- Generate Requried Certficates for organziation to partipate in network
 - `./bin/cryptogen generate --config=./crypto-config.yaml` 
 - `./bin/configtxgen -profile DemoProfile -channelID byfn-sys-channel -outputBlock ./config/genesis.block`  generate genesis block using above command 
 - `./bin/configtxgen -profile  ClaimsChannel  -outputCreateChannelTx ./config/insurancechannel.tx -channelID insuranceclaim`
 - above command is used to create channel .one thing to remember is channel name must be lower case
 - 



