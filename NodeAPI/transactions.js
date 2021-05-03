'use strict';

const { Gateway,Wallets } = require('fabric-network');
let enroll = require("./enrollment.js")
var query = require('./query.js')
const fs = require('fs');
const path = require('path');
const request = require('request');
const defaultsPath = path.resolve(__dirname, 'connection.json');
const defaultsJSON = fs.readFileSync(defaultsPath, 'utf8');
const defaults = JSON.parse(defaultsJSON);

const ccpPath = path.resolve(__dirname, 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp =  JSON.parse(ccpJSON);
const walletPath = path.join(process.cwd(), 'wallet');

async function fndoesUserExists(_user) {
    console.log('step-1',_user)
    const wallet = await  Wallets.newFileSystemWallet(walletPath);
    let userExists = await wallet.get(_user);
    if (userExists) {
        return true;
    } else {
        return false
    }
}

async function submitInvoke(user, _fcn, _args) {

    console.info("invoking With admin : ", user)
    if (await fndoesUserExists(user)) {
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        const identity = await wallet.get(user);
        try {
            const gateway = new Gateway();
            //console.log("step-1")
            await gateway.connect(ccp, { wallet, identity: user, discovery: { enabled: true, asLocalhost: false} });
            const network = await gateway.getNetwork('insuranceclaim');
            // Get the contract from the network.
            const contract = network.getContract('ibc');
            let result = await contract.submitTransaction(_fcn, _args)
             console.log("result tx id:", result)
            console.log('Transaction has been submitted');
            await gateway.disconnect();
            //return { success: true, message: "Transaction successfully sumbmitted. " + result.toString() }
            return result.toString();
            } catch (error) {
            
            return { success: false, message: `Failed to submit transaction: ${error}` }
        }

    }
    return { success: false, message: `Failed to submit transaction` }
}




exports.submitInvoke = submitInvoke
