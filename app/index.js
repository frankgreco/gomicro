var Generator = require('yeoman-generator');
var chalk = require('chalk');
var leftpad = require('left-pad');
var fs = require('fs');
var pluralize = require('pluralize');
var cmd = require('command-exists').sync
var certs = require('selfsigned');
var fs = require('fs');

module.exports = class extends Generator {

    constructor(args, opts) {
        super(args, opts);
    }

    initializing() {
        var errorMessage = "";
        var hasSoftError = false;
        var hasHardError = false;
        var baseMessage = `\n` +
            chalk.yellow(`=====================================================================\n`) +
            chalk.gray(leftpad(`Welcome to`, 35)) + chalk.bold.blue(` gomicro`) + `!\n` +
            chalk.gray(`This Yeoman generator aims to scaffold a robust RESTful microservice.\n`) +
            '\n' + chalk.red(leftpad(`[WARNING] this project is under active development!\n`, 63)) +
            chalk.red(leftpad(`results may not be as expected.`, 52)) + '\n\n';

        if(!cmd('go')) {
            hasSoftError = true;
            errorMessage += chalk.yellow(leftpad(`[WARNING] go not installed or not in PATH\n`, 55))
        }

        if(!cmd('glide')) {
            hasSoftError = true;
            errorMessage += chalk.yellow(leftpad(`[WARNING] glide not installed or not in PATH\n`, 58))
        }

        if(!process.env.GOPATH || !fs.lstatSync(process.env.GOPATH).isDirectory()) {
            hasHardError = true;
            errorMessage += chalk.red(leftpad(`[ERROR] GOPATH not properly configured`, 51)) + '\n'
        }

        if(hasHardError) {
            this.log(`${baseMessage}` + (errorMessage == "" ? '' : '\n') + `${errorMessage}` + (errorMessage == "" ? '' : '\n') + chalk.white(leftpad(`Please fix the above errors and try again`, 55)) +
                 chalk.yellow(`\n=====================================================================`)
             )
            process.exit(1)
        }
        this.log(`${baseMessage}` + (errorMessage == "" ? '' : '\n') + `${errorMessage}` + (errorMessage == "" ? '' : '\n') + chalk.white(leftpad(`Let\'s get started!`, 43)) +
             chalk.yellow(`\n=====================================================================`)
        )
    }

    prompting() {
        this.log(chalk.blue(`\n=====================================================================\n`) + chalk.gray(leftpad('Go workspaces like to be created a certain way', 57)) + '\n' + chalk.gray(leftpad(`(e.g. `, 17)) + chalk.green(`$GOPATH/src/github.com/frankgreco/gohttp`) + chalk.gray(`)`) + '\n' + chalk.gray(leftpad('To do this, let\'s get some information about your project', 63)) + chalk.blue(`\n=====================================================================\n`));
        return this.prompt([{
            type    : 'input',
            name    : 'name',
            message : 'your name',
            default : 'Frank B Greco Jr',
            store: true
        }, {
            type    : 'input',
            name    : 'email',
            message : 'your email',
            default : 'frank@petrasphere.io',
            store: true
        }, {
            type    : 'input',
            name    : 'vcs',
            message : 'vcs',
            default : 'github.com',
            store: true
        }, {
            type    : 'input',
            name    : 'user',
            default : 'frankgreco',
            store   : true,
            message : function(answers) {
                return `${answers.vcs} username`
            },
        }, {
            type    : 'input',
            name    : 'project',
            message : 'project name',
            default : 'gohttp',
            store   : true
        }, {
            type    : 'input',
            name    : 'singular',
            message : 'resource noun (singular)',
            default : 'person',
            store   : true
        }, {
            type    : 'input',
            name    : 'plural',
            message : 'resource noun (plural)',
            store   : true,
            default : function(answers) {
                return pluralize(answers.singular)
            }
        }, {
            type    : 'checkbox',
            name    : 'schemes',
            message : 'http schemes',
            choices : [
                {
                    name: 'http',
                    checked: true
                },
                {
                    name: 'https'
                }
            ],
            default : ['http'],
            store   : true,
            validate: function(answers){
                return answers.length < 1 ? 'choose at least one scheme' : true
            }
        }, {
            type    : 'list',
            name    : 'db',
            message : 'database type',
            choices : [ 'mysql', 'postgres', 'sqlite', 'mongodb'],
            default : 'mysql',
            store   : true,
        }]).then((answers) => {
            return;
        });
    }

    writing() {

        var basePath = `${process.env.GOPATH}/src/${this.config.get("promptValues").vcs}/${this.config.get("promptValues").user}/${this.config.get("promptValues").project}`

        var self = this;

        this.log(`\n` +
            chalk.green(`=====================================================================\n`) +
            chalk.gray(leftpad(`Creating certificates...`, 47))
        )

        //generating certificates
        var pems = certs.generate([
            { name: 'commonName', value: 'localhost' }
        ], {
            keySize: 2048,
            days: 365,
            algorithm: 'sha256'
        })

        if(!fs.existsSync(`${basePath}/certs`)) {
            fs.mkdirSync(`${basePath}/certs`)
        }
        var crt = fs.writeFileSync(`${basePath}/certs/server.crt`, pems.cert, { mode: 400 })
        var key = fs.writeFileSync(`${basePath}/certs/server.key`, pems.private, { mode: 400 })

        this.log(chalk.gray(leftpad(`Creating your workspace...`, 49)) +
            chalk.green(`\n=====================================================================\n`)
        )

        var cap = function capitalizeFirstLetter(string) {
            return string.charAt(0).toUpperCase() + string.slice(1);
        }
        var params = {
            name: this.config.get("promptValues").name,
            email: this.config.get("promptValues").email,
            vcs: this.config.get("promptValues").vcs,
            repo: this.config.get("promptValues").user,
            project: this.config.get("promptValues").project,
            nounSingularUpper: cap(this.config.get("promptValues").singular),
            nounSingularLower: this.config.get("promptValues").singular,
            nounPluralUpper: cap(this.config.get("promptValues").plural),
            nounPluralLower: this.config.get("promptValues").plural,
            db: this.config.get("promptValues").db,
            http: this.config.get("promptValues").schemes.includes('http')
            https: this.config.get("promptValues").schemes.includes('https')
        }

        var templateFiles = [
            {from: "cmd/root.go",           to: "cmd/root.go"},
            {from: "cmd/start.go",          to: "cmd/start.go"},
            {from: "cmd/version.go",        to: "cmd/version.go"},
            {from: "database/driver.go",    to: "database/driver.go"},
            {from: "handler/handler.go",    to: "handler/handler.go"},
            {from: "handler/plural.go",     to: `handler/${this.config.get("promptValues").plural}.go`},
            {from: "handler/singular.go",   to: `handler/${this.config.get("promptValues").singular}.go`},
            {from: "handler/util.go",       to: "handler/util.go"},
            {from: "models/model.go",       to: "models/model.go"},
            {from: "route/logger.go",       to: "route/logger.go"},
            {from: "route/router.go",       to: "route/router.go"},
            {from: "route/routes.go",       to: "route/routes.go"},
            {from: "server/server.go",      to: "server/server.go"},
            {from: "utils/error.go",        to: "utils/error.go"},
            {from: "utils/flag.go",         to: "utils/flag.go"},
            {from: ".gitignore",            to: ".gitignore"},
            {from: "Dockerfile",            to: "Dockerfile"},
            {from: "glide.yaml",            to: "glide.yaml"},
            {from: "main.go",               to: "main.go"},
            {from: "Makefile",              to: "Makefile"},
            {from: "swagger.json",          to: "swagger.json"},
        ]

        templateFiles.forEach(function(file) {
            self.fs.copyTpl(
                self.templatePath(file.from),
                self.destinationPath(`${basePath}/${file.to}`),
                params
            );
        });

        if(['mysql', 'postgres'].includes(params.db)) {
            this.fs.copyTpl(
                this.templatePath(`schemas/${params.db}.sql`),
                this.destinationPath(`${basePath}/schema.sql`),
                params
            );
        }

        if(!['sqlite'].includes(params.db)) {
            this.fs.copyTpl(
                this.templatePath('docker-compose.yaml'),
                this.destinationPath(`${basePath}/docker-compose.yaml`),
                params
            );
        }

    }

    end() {
        var answers = this.config.get("promptValues")
        this.log(`\n` +
            chalk.magenta(`=====================================================================\n`) +
            chalk.gray(leftpad(`You\'re almost done! Your workspace has been created here:`, 63)) + `\n` +
            chalk.white.bold(leftpad(`$GOPATH/src/${answers.vcs}/${answers.user}/${answers.project}/`, 52)) + `\n` +
            chalk.gray(leftpad(`To complete your setup, run the following commands in your workspace:`, 55)) + `\n` +
            chalk.white(leftpad('$ ', 17)) + chalk.cyan(`make `) + chalk.gray(leftpad(`(use your own database)`, 35)) + `\n` +
            chalk.white(leftpad('$ ', 17)) + chalk.cyan(`make local-dev `) + chalk.gray(leftpad(`(or, create a local database)`, 31)) + `\n` +
            chalk.white(leftpad('$ ', 17)) + chalk.cyan(`./${answers.project} --help`) + chalk.gray(leftpad(`(example usage)`, 18)) +
            chalk.magenta(`\n=====================================================================`)
        )
    }

};
