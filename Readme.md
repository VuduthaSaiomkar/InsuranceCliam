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
#### steps to follow to start the network

- cd InsuranceClaim 
- run the following command `docker up -d`
- docker exec -ti cli bash 
- peer channel create -o orderer.example.com:7050  -c insuranceclaim -f ./config/channel.tx --outputBlock channel.block --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
- peer channel join -b channel.block
- Repeat some for all organatiaion 

### Install and deploy chaincode

####### create package for chaincode 
- peer lifecycle chaincode package ibcs.tar.gz --path github.com/chaincode/ --lang golang --label ibcs_1
### install chaincode
- peer lifecycle chaincode install ibcs.tar.gz
- peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID insuranceclaim --name ibcs --version 1 --init-required --package-id  ibcs_1:9ee015b8815782f66d98efe8795d35fea4bb43abe6fd72692d8b397359d71438 --sequence 1
#### change variables and install chaincode  police

`CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/police.example.com/peers/peer0.police.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/police.example.com/users/Admin@police.example.com/msp
CORE_PEER_LOCALMSPID=policeMSP
CORE_PEER_ADDRESS=peer0.police.example.com:7051`

- peer lifecycle chaincode install ibcs.tar.gz
- peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID insuranceclaim --name ibcs --version 1 --init-required --package-id  ibcs_1:9ee015b8815782f66d98efe8795d35fea4bb43abe6fd72692d8b397359d71438 --sequence 1

#### change variables and install chaincode  workshop
`CORE_PEER_LOCALMSPID=workshopMSP
CORE_PEER_ADDRESS=peer0.workshop.example.com:7051
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/workshop.example.com/users/Admin@workshop.example.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/workshop.example.com/peers/peer0.workshop.example.com/tls/ca.crt`

- peer lifecycle chaincode install ibcs.tar.gz
- peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID insuranceclaim --name ibcs --version 1 --init-required --package-id  ibcs_1:9ee015b8815782f66d98efe8795d35fea4bb43abe6fd72692d8b397359d71438 --sequence 1


###### deploy the chaincode form alteast one org

- peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID insuranceclaim --name ibcs --version 1 --sequence 1 --init-required --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.police.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/police.example.com/peers/peer0.police.example.com/tls/ca.crt --peerAddresses peer0.workshop.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/workshop.example.com/peers/peer0.workshop.example.com/tls/ca.crt

-------------------
###### chaincode functions and access control




