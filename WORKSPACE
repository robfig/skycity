load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b7a62250a3a73277ade0ce306d22f122365b513f5402222403e507f2f997d421",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.16.3/rules_go-0.16.3.tar.gz",
)

git_repository(
    name = "bazel_gazelle",
    commit = "3c4f16ae4a2117f0908f58107d1b55e533c9a431",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# GENERATED CODE - do not edit
#
# To update: gazelle update-repos -from_file=go.mod

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    commit = "d65d576e9348",
    importpath = "github.com/kylelemons/godebug",
)

go_repository(
    name = "com_github_stripe_skycfg",
    commit = "02c893cb1266",
    importpath = "github.com/stripe/skycfg",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "20d25e280405",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.1",
)

go_repository(
    name = "net_starlark_go",
    commit = "f4938bde4080",
    importpath = "go.starlark.net",
)

go_repository(
    name = "org_golang_x_net",
    commit = "26e67e76b6c3",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "1d60e4601c6f",
    importpath = "golang.org/x/sync",
)
