def _get_docker_password(ctx):
    output = ctx.actions.declare_file(ctx.attr.name)

    ctx.actions.run(
        outputs = [output],
        executable = ctx.executable._docker_password,
    )

    return [
        DefaultInfo(
            files = depset([output]),
        ),
    ]

get_docker_password = rule(
    implementation = _get_docker_password,
    attrs = {
        "_docker_password": attr.label(
            executable = True,
            cfg = "host",
            default = Label("//:docker-password-test"),
        ),
    },
)

def _docker_login_impl(ctx):
    password = ctx.file._password

    output = ctx.actions.declare_file(ctx.attr.name)

    ctx.actions.run_shell(
        inputs = [password],
        outputs = [output],
        command = "cat %s | docker login --username AWS --password-stdin $SL_DEVOPS_ECR" % password.path,
    )

    return [
        DefaultInfo(
            files = depset([output]),
        ),
    ]

docker_login = rule(
    implementation = _docker_login_impl,
    attrs = {
        "_password": attr.label(
            allow_single_file = True,
            default = Label("//:docker-password"),
        ),
    },
)
