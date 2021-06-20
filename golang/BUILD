load("@bazel_gazelle//:def.bzl", "gazelle")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

# gazelle:prefix github.com/veganafro/mono/server
# gazelle:go_grpc_compilers @com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-grpc-gateway:go_gen_grpc_gateway, @io_bazel_rules_go//proto:go_grpc

gazelle(name = "gazelle")

pkg_tar(
    name = "build_tar",
    mode = "0755",
    package_dir = "/",
    srcs = [
        "BUILD",
    ],
    strip_prefix = "/",
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "repositories_tar",
    mode = "0755",
    package_dir = "/",
    srcs = [
        "repositories.bzl",
    ],
    strip_prefix = "/",
    visibility = ["//visibility:public"],
)
