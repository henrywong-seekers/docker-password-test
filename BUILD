load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/henrywong-seekers/docker-password-test
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["docker-password.go"],
    importpath = "github.com/henrywong-seekers/docker-password-test",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_aws_aws_sdk_go_v2//aws/external:go_default_library",
        "@com_github_aws_aws_sdk_go_v2//service/ecr:go_default_library",
    ],
)

go_binary(
    name = "docker-password-test",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
