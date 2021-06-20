load("@rules_pkg//:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "workspace_tar",
    mode = "0755",
    package_dir = "/",
    srcs = [
        "WORKSPACE",
    ],
    strip_prefix = "/",
    visibility = ["//visibility:public"],
)

