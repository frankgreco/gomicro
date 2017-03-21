var Generator = require('yeoman-generator');
var chalk = require('chalk');
var leftpad = require('left-pad');
var fs = require('fs');
var pluralize = require('pluralize');
var cmd = require('command-exists').sync

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
            },
        }, {
            type    : 'list',
            name    : 'db',
            message : 'database type',
            choices : [ 'mongodb', 'mysql'],
            defulat : 'mongodb',
            store   : true,
        }]).then((answers) => {
            return;
        });
    }

    writing() {

        var self = this;

        this.log(`\n` +
            chalk.green(`=====================================================================\n`) +
            chalk.gray(leftpad(`Creating your workspace...`, 45)) +
            chalk.green(`\n=====================================================================\n`)
        )

        var cap = function capitalizeFirstLetter(string) {
            return string.charAt(0).toUpperCase() + string.slice(1);
        }

        var basePath = `${process.env.GOPATH}/src/${this.config.get("promptValues").vcs}/${this.config.get("promptValues").user}/${this.config.get("promptValues").project}`

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
            db: this.config.get("promptValues").db
        }

        var templateFiles = [
            "Makefile",
            "table.sql",
            "Dockerfile",
            "glide.yaml",
            "swagger.json",
            "docker-compose.yaml",
            ".gitignore",
            "main.go",
            "route/logger.go",
            "route/router.go",
            "route/routes.go",
            "cmd/root.go",
            "cmd/start.go",
            "cmd/version.go",
            "utils/error.go",
            "utils/flag.go",
            "server/server.go",
            "handler/handler.go",
            "handler/util.go"
        ];

        templateFiles.forEach(function(file) {
            self.fs.copyTpl(
                self.templatePath(file),
                self.destinationPath(`${basePath}/${file}`),
                params
            );
        });

        this.fs.copyTpl(
            this.templatePath('database/' + (params.db == 'mysql' ? 'mysql' : 'mongo') + '.go'),
            this.destinationPath(`${basePath}/database/driver.go`),
            params
        );

        this.fs.copyTpl(
            this.templatePath('database/mysql.go'),
            this.destinationPath(`${basePath}/database/driver.go`),
            params
        );

        this.fs.copyTpl(
            this.templatePath('handler/singular.go'),
            this.destinationPath(`${basePath}/handler/${this.config.get("promptValues").singular}.go`),
            params
        );

        this.fs.copyTpl(
            this.templatePath('handler/plural.go'),
            this.destinationPath(`${basePath}/handler/${this.config.get("promptValues").plural}.go`),
            params
        );

        this.fs.copyTpl(
            this.templatePath('models/model.go'),
            this.destinationPath(`${basePath}/models/${this.config.get("promptValues").singular}.go`),
            params
        );

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
            chalk.white(leftpad('$ ', 17)) + chalk.cyan(`./${answers.project} --help`) + chalk.gray(leftpad(`(example usage)`, 17)) +
            chalk.magenta(`\n=====================================================================`)
        )
    }

};
