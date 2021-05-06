"""
`go_repository_dependencies` loads all the relevant Go dependencies for the project
"""
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repository_dependencies():
    """
    `go_repository_dependencies` loads all the relevant Go dependencies for the project
    """
    ## Begin gRPC dependecies ##

    GRPC_SUM = "h1:SfXqXS5hkufcdZ/mHtYCh53P2b+92WQq/DZcKLgsFRs="
    GRPC_VERSION = "v1.31.1"

    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "%s" % GRPC_SUM,
        version = "%s" % GRPC_VERSION,
    )

    NET_SUM = "h1:MXfv8rhZWmFeqX3GNZRsd6vOLoaCHjYEX3qkRo3YBUA="
    NET_VERSION = "v0.0.0-20200904194848-62affa334b73"

    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "%s" % NET_SUM,
        version = "%s" % NET_VERSION,
    )

    TEXT_SUM = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k="
    TEXT_VERSION = "v0.3.3"

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "%s" % TEXT_SUM,
        version = "%s" % TEXT_VERSION,
    )

    ## End gRPC dependencies ##

    ## Begin grpc-gateway dependencies ##

    GLOG_COMMIT = "23def4e6c14b4da8ac2ed8007337bc5eb5007998"

    go_repository(
        name = "com_github_golang_glog",
        commit = "%s" % GLOG_COMMIT,
        importpath = "github.com/golang/glog",
    )

    YAML_COMMIT = "25d852aebe32c875e9c044af3eef9c7dc6bc777f"

    go_repository(
        name = "com_github_ghodss_yaml",
        commit = "%s" % YAML_COMMIT,
        importpath = "github.com/ghodss/yaml",
    )

    YAML_V2_COMMIT = "0b1645d91e851e735d3e23330303ce81f70adbe3"

    go_repository(
        name = "in_gopkg_yaml_v2",
        commit = "%s" % YAML_V2_COMMIT,
        importpath = "gopkg.in/yaml.v2",
    )

    ## End grpc-gateway dependencies ##
