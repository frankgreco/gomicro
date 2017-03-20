var Generator = require('yeoman-generator');
var chalk = require('chalk');
var leftpad = require('left-pad');
var fs = require('fs');
var pluralize = require('pluralize')

module.exports = class extends Generator {

    constructor(args, opts) {
        super(args, opts);
    }

    initializing() {
        this.log(`\n` +
            chalk.yellow(`=================================================================\n`) +
            chalk.gray(leftpad(`Welcome to`, 33)) + chalk.bold.blue(` gohttp`) + `!\n` +
            chalk.gray(`This Yeoman generator aims to scaffold a robust http web service.\n`) +
            chalk.white(leftpad(`Let\'s get started!`, 41)) +
            chalk.yellow(`\n=================================================================`)
        )

        // if(!process.env.GOPATH || !fs.lstatSync(process.env.GOPATH).isDirectory()) {
        //     this.log(`\n` +
        //         chalk.red(`=================================================================\n`) +
        //         chalk.gray(leftpad(`Welcome to`, 33)) + chalk.bold.blue(` gohttp`) + `!\n` +
        //         chalk.gray(`This Yeoman generator aims to scaffold a robust http web service.\n`) +
        //         chalk.white(leftpad(`Let\'s get started!`, 41)) +
        //         chalk.red(`\n=================================================================`)
        //     )
        //     process.exit(1)
        // }
    }

    prompting() {
        this.log(chalk.blue(`\n=================================================================\n`) + chalk.gray(leftpad('Go workspaces like to be created a certain way', 55)) + '\n' + chalk.gray(leftpad(`(e.g. `, 15)) + chalk.green(`$GOPATH/src/github.com/frankgreco/gohttp`) + chalk.gray(`)`) + '\n' + chalk.gray(leftpad('To do this, let\'s get some information about your project', 61)) + chalk.blue(`\n=================================================================\n`));
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

    end() {
        var answers = this.config.get("promptValues")
        this.log(`\n` +
            chalk.magenta(`=================================================================\n`) +
            chalk.gray(leftpad(`You\'re all set! Find your project here:`, 51)) + `\n` +
            chalk.white.bold(leftpad(`$GOPATH/${answers.repo}/${answers.user}/${answers.project}/`, 50)) +
            chalk.magenta(`\n=================================================================`)
        )
    }

    writing() {

        var self = this;

        this.log(`\n` +
            chalk.green(`=================================================================\n`) +
            chalk.gray(leftpad(`Creating your workspace...`, 45)) +
            chalk.green(`\n=================================================================\n`)
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
            nounPluralLower: this.config.get("promptValues").plural
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
            "database/mysql.go",
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

    // install() {
    //     this.log(`\n` +
    //         chalk.green(`=================================================================\n`) +
    //         chalk.gray(leftpad(`Installing dependencies...`, 45)) +
    //         chalk.green(`\n=================================================================`)
    //     )
    //     this.spawnCommand('glide', ['update']);
    //     this.spawnCommand('glide', ['install']);
    //     this.spawnCommand('go', ['build']);
    // }

};
