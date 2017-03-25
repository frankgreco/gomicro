var pluralize = require('pluralize');

var name = {
    type    : 'input',
    name    : 'name',
    message : 'your name',
    default : 'Frank B Greco Jr',
    store   : true
}

var email = {
    type    : 'input',
    name    : 'email',
    message : 'your email',
    default : 'frank@petrasphere.io',
    store   : true
}

var vcs = {
    type    : 'input',
    name    : 'vcs',
    message : 'vcs',
    default : 'github.com',
    store   : true
}

var vcs_user = {
    type    : 'input',
    name    : 'user',
    default : 'frankgreco',
    store   : true,
    message : function(answers) {
        return `${answers.vcs} username`
    }
}

var project_name = {
    type    : 'input',
    name    : 'project',
    message : 'project name',
    default : 'gomicro',
    store   : true
}

var resource_singular = {
    type    : 'input',
    name    : 'singular',
    message : 'resource noun (singular)',
    default : 'person',
    store   : true
}

var resource_plural = {
    type    : 'input',
    name    : 'plural',
    message : 'resource noun (plural)',
    store   : true,
    default : function(answers) {
        return pluralize(answers.singular)
    }
}

var http_scheme = {
    type    : 'list',
    name    : 'scheme',
    message : 'http scheme',
    choices : [ 'http', 'https'],
    default : 'http',
    store   : true
}

var db_type = {
    type    : 'list',
    name    : 'db',
    message : 'database type',
    choices : [ 'mysql', 'postgres', 'sqlite', 'mongodb'],
    default : 'mysql',
    store   : true
}

function capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

exports.get = function() {

    return [
        name,
        email,
        vcs,
        vcs_user,
        project_name,
        resource_singular,
        resource_plural,
        http_scheme,
        db_type
    ]

}

exports.getValues = function(promptAnswers) {

    return {
        name                : promptAnswers.name,
        email               : promptAnswers.email,
        vcs                 : promptAnswers.vcs,
        repo                : promptAnswers.user,
        project             : promptAnswers.project,
        nounSingularUpper   : capitalize(promptAnswers.singular),
        nounSingularLower   : promptAnswers.singular,
        nounPluralUpper     : capitalize(promptAnswers.plural),
        nounPluralLower     : promptAnswers.plural,
        db                  : promptAnswers.db,
        scheme              : promptAnswers.scheme
    }

}
