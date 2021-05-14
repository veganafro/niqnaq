"""
`go_grpc_deps` loads all the relevant Go dependencies for gRPC

`go_grpc_gateway_deps` loads all the relevant Go dependencies for grpc-gateway
"""
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_grpc_deps():
    """
    `go_grpc_deps` loads all the relevant Go dependencies for gRPC
    """
    ## Begin gRPC dependecies - 05/13 ##

    GRPC_SUM = "h1:o1bcQ6imQMIOpdrO3SWf2z5RV72WbDwdXuK0MDlc8As="
    GRPC_VERSION = "v1.36.0"

    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "%s" % GRPC_SUM,
        version = "%s" % GRPC_VERSION,
    )

    NET_SUM = "h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg="
    NET_VERSION = "v0v0.0.0-20210119194325-5f4716e94777"

    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "%s" % NET_SUM,
        version = "%s" % NET_VERSION,
    )

    TEXT_SUM = "h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ="
    TEXT_VERSION = "v0.3.5"

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "%s" % TEXT_SUM,
        version = "%s" % TEXT_VERSION,
    )

    ## End gRPC dependencies ##

def go_grpc_gateway_deps():
    """
    `go_grpc_gateway_deps` loads all the relevant Go dependencies for grpc-gateway
    """
    ## Begin grpc-gateway dependencies - 05/13 ##

    YAML_SUM = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk="
    YAML_VERSION = "v1.0.0"

    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "%s" % YAML_SUM,
        version = "%s" % YAML_VERSION,
    )

    YAML_V2_SUM = "h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI="
    YAML_V2_VERSION = "v2.3.0"

    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "%s" % YAML_V2_SUM,
        version = "%s" % YAML_V2_VERSION,
    )

    ## End grpc-gateway dependencies ##
