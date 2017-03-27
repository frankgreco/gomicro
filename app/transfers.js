exports.get = function(params) {

    return [
        createTransfer("auth/auth.go", params.auth),
        createTransfer("auth/basic.go", params.auth),
        createTransfer("auth/token.go", params.auth),
        createTransfer("cmd/root.go"),
        createTransfer("cmd/start.go"),
        createTransfer("cmd/version.go"),
        createTransfer("database/driver.go"),
        createTransfer("handler/handler.go"),
        createTransfer("cmd/root.go"),
        createTransfer("handler/plural.go", `handler/${params.plural}.go`),
        createTransfer("handler/singular.go", `handler/${params.singular}.go`),
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
        createTransfer("deploy/docker-compose.yaml", params.orchestration == 'docker swarm'),
        createTransfer("deploy/kubernetes.yaml", params.orchestration == 'kubernetes')
    ]

}

function createTransfer(from, to, condition) {

    condition = condition || to
    to = typeof to == 'boolean' ? undefined : to

    if(condition && !condition) {
        return
    }

    return {
        from: from,
        to: to ? to : from
    }

}
