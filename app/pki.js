var certs   = require('selfsigned');
var fs      = require('fs');
var path    = require('path');
var mkdirp  = require('mkdirp');

fs.mkdirParent = function(dirPath, mode, callback) {
    fs.mkdir(dirPath, mode, function(error) {
        if (error) {
            fs.mkdirParent(path.dirname(dirPath), mode, callback);
            fs.mkdirParent(dirPath, mode, callback);
        }
        callback && callback(error);
    });
}

exports.create = function(basePath) {

    var pems = certs.generate([{
        name: 'commonName',
        value: 'localhost'
    }], {
        keySize: 2048,
        days: 365,
        algorithm: 'sha256'
    });

    mkdirp.sync(`${basePath}/certs`)
    writeCertsToFile(basePath, pems)

}

function writeCertsToFile(basePath, certs) {

    fs.writeFileSync(`${basePath}/certs/server.crt`, certs.cert, {
        mode: 400
    });
    fs.writeFileSync(`${basePath}/certs/server.key`, certs.private, {
        mode: 400
    });

}
