exports.get = function(params) {

    return [
        createTransfer("auth/auth.go", params.auth),
        createTransfer("auth/basic.go", params.auth),
        createTransfer("auth/token.go", params.auth),
        createTransfer("cmd/root.go"),
        createTransfer("cmd/test.go"),
        createTransfer("cmd/start.go"),
        createTransfer("cmd/version.go"),
        createTransfer("database/driver.go"),
        createTransfer("handler/handler.go"),
        createTransfer("cmd/root.go"),
        createTransfer("handler/plural.go", `handler/${params.nounPluralLower}.go`),
        createTransfer("handler/singular.go", `handler/${params.nounSingularLower}.go`),
        createTransfer("handler/util.go"),
        createTransfer("models/model.go"),
        createTransfer("route/logger.go"),
        createTransfer("route/router.go"),
        createTransfer("route/routes.go"),
        createTransfer("server/server.go"),
        createTransfer("utils/error.go"),
        createTransfer("utils/flag.go"),
        createTransfer(".gitignore"),
        createTransfer("Dockerfile"),
        createTransfer("glide.yaml"),
        createTransfer("main.go"),
        createTransfer("Makefile"),
        createTransfer("swagger.json"),
        createTransfer(`schemas/${params.db}.sql`, 'schema.sql', ['mysql', 'postgres'].includes(params.db)),
        createTransfer('docker-compose.yaml', !['sqlite'].includes(params.db)),
        createTransfer("basic.csv", params.auth),
        createTransfer("token.csv", params.auth),
        createTransfer("deploy/docker-compose.yaml", (params.orchestration == 'docker swarm')),
        createTransfer("deploy/kubernetes.yaml", (params.orchestration == 'kubernetes'))
    ]

}

function createTransfer(from, to, condition) {

    if(from == 'docker-compose.yaml') {
        console.log(from)
        console.log(to)
        console.log(condition)
    }

    if(typeof to == 'boolean') {
        condition = to;
        to = undefined;
    }

    if(from == 'docker-compose.yaml') {
        console.log(from)
        console.log(to)
        console.log(condition)
    }

    if(typeof condition !== 'undefined' && !condition) {
        return
    }

    return {
        from: from,
        to: to ? to : from
    }

}
