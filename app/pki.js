var certs = require('selfsigned');
var fs = require('fs');

exports.create = function(basePath) {

    var pems = certs.generate([{
        name: 'commonName',
        value: 'localhost'
    }], {
        keySize: 2048,
        days: 365,
        algorithm: 'sha256'
    });

    createCertDirectory(basePath);
    writeCertsToFile(basePath, pems);

}

function createCertDirectory(basePath) {

    if(!fs.existsSync(`${basePath}/certs`)) {
        fs.mkdirSync(`${basePath}/certs`)
    }

}

function writeCertsToFile(basePath, certs) {

    fs.writeFileSync(`${basePath}/certs/server.crt`, certs.cert, {
        mode: 400
    });
    fs.writeFileSync(`${basePath}/certs/server.key`, certs.private, {
        mode: 400
    });

}
