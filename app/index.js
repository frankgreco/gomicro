let Generator   = require('yeoman-generator');
let pki         = require('./pki');
let prompts     = require('./prompts');
let transfers   = require('./transfers');
let messages    = require('./messages');
let preflight   = require('./preflight');

module.exports = class extends Generator {

    constructor(args, opts) {

        super(args, opts);

    }

    initializing() {

        var errors = preflight.execute()

        this.log(messages.init(errors))

        if(errors.length > 0) {
            process.exit(1)
        }
    }

    prompting() {

        this.log(messages.prompt())
        return this.prompt(prompts.get())

    }

    certificates() {

        let params = prompts.getValues(this.config.get("promptValues"))
        let basePath = `${process.env.GOPATH}/src/${params.vcs}/${params.repo}/${params.project}`

        if(params.scheme == 'https') {
            this.log(messages.certs())
            pki.create(basePath)
        }

    }

    writing() {

        let self = this;
        let params = prompts.getValues(this.config.get("promptValues"))
        let basePath = `${process.env.GOPATH}/src/${params.vcs}/${params.repo}/${params.project}`

        this.log(messages.write())

        transfers.get(params).forEach(file => {
            self.fs.copyTpl(
                self.templatePath(file.from),
                self.destinationPath(`${basePath}/${file.to}`),
                params
            );
        });

    }

    install() {

        let params = prompts.getValues(this.config.get("promptValues"))
        let basePath = `${process.env.GOPATH}/src/${params.vcs}/${params.repo}/${params.project}`

        this.destinationRoot(basePath)

        // this.log('install was executed')
        // this.spawnCommand('cd', [
        //     basePath,
        //     '&&',
        //     'make'
        // ]);
        this.log(this.contextRoot)

    }

    end() {

        this.log(messages.end(this.config.get("promptValues")))

    }

}
