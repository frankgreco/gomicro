let helpers = require('yeoman-test');
let path    = require('path');
let assert  = require('yeoman-assert');

describe('gomicro', () => {

    describe('#indexOf()', function() {
        it('should return -1 when the value is not present', function() {
            assert.equal(-1, [1,2,3].indexOf(4));
        });
    });

});
