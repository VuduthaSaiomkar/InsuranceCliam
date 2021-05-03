'use strict';

const FabricCAServices = require('fabric-ca-client');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);
const defaultsPath = path.resolve(__dirname, 'defaults.json');
const defaultsJSON = fs.readFileSync(defaultsPath, 'utf8');
const defaults = JSON.parse(defaultsJSON);
let Org = defaults["Organisation"]
const caInfo = ccp.certificateAuthorities['ca-workshop'];
console.log(caInfo)
const caTLSCACerts = caInfo.tlsCACerts.pem;
const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName)
const walletPath = path.join(process.cwd(), 'wallet');
var adminIdentity;


async function enrolluser(username, secert) {
    try {
        const wallet = await  Wallets.newFileSystemWallet(walletPath);
        const userExists = await wallet.get(username);
                if (userExists) {
                    console.log('An identity for the ' + username + ' already exists in the wallet');
                    return { "success": true, "message": "Login success" }
                }else{
                    return {success:false,"message":"please register use first"}
                }
                } catch (err) {
        return { "success": false, "message": `error connecting to network ${err}` }
    }
}

async function registeruser(username,secret,org) {
    try {
        console.log("entered into register user")
        const wallet = await  Wallets.newFileSystemWallet(walletPath);
        var adminExists = await wallet.get('admin');
        if (adminExists) {
            const provider = wallet.getProviderRegistry().getProvider(adminExists.type);
            const adminUser = await provider.getUserContext(adminExists, 'admin');    
            const secret = await ca.register({
                affiliation: 'org1.department1',
                enrollmentID: username,
                role: 'client'
            }, adminUser);
            const enrollment = await ca.enroll({
                enrollmentID: username,
                enrollmentSecret: secret
            });
            const x509Identity = {
                credentials: {
                    certificate: enrollment.certificate,
                    privateKey: enrollment.key.toBytes(),
                },
                mspId: 'workshopMSP',
                type: 'X.509',
            };
            await wallet.put(username, x509Identity);
        } else {
            return { "success": false, "message": `not able to register admin into network ${err}` }
        }
    } catch (error) {
        console.error(`Failed to register user ${username}: ${error}`);
    }
}

async function enrolladmin() {
    try {
        const wallet = await  Wallets.newFileSystemWallet(walletPath);
        const adminExists = await wallet.get('admin');
        if (adminExists) {
            return { "success": true, "message": `Already Admin exists` }
        }
        const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: 'workshopMSP',
            type: 'X.509',
        };
        await wallet.put('admin', x509Identity);
        return { "success": true, "message": `successfully registered Admin` }
    } catch (error) {
        console.error(`Failed to enroll admin user "admin": ${error}`);
        return { "success": false, "message": `Failed to enroll admin user "admin": ${error}` }
    }
}




exports.enrolluser = enrolluser;
exports.registeruser = registeruser;
exports.enrolladmin = enrolladmin;
