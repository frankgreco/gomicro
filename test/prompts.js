let assert  = require('yeoman-assert');
let prompts = require('../app/prompts');

describe('prompts', () => {

    it('getValues', function() {

        assert.deepEqual(prompts.getValues({
            name            : 'Frank B Greco Jr',
            email           : 'frank@petrasphere.io',
            vcs             : 'github.com',
            user            : 'frankgreco',
            project         : 'gomicro',
            singular        : 'call',
            plural          : 'calls',
            db              : 'mysql',
            scheme          : 'https',
            auth            : true,
            orchestration   : 'kubernetes'
        }), {
            name                : 'Frank B Greco Jr',
            email               : 'frank@petrasphere.io',
            vcs                 : 'github.com',
            repo                : 'frankgreco',
            project             : 'gomicro',
            nounSingularUpper   : 'Call',
            nounSingularLower   : 'call',
            nounPluralUpper     : 'Calls',
            nounPluralLower     : 'calls',
            db                  : 'mysql',
            scheme              : 'https',
            auth                : true,
            orchestration       : 'kubernetes'
        }, 'return values do not match expected')
    });

});
