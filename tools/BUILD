load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "buildozer",
    srcs = ["buildozer.go"],
    deps = ["@com_github_bazelbuild_buildtools//edit:edit"],
    goos = "wasip1",
    goarch = "wasm",
    # build as a reactor
    linkmode = "c-shared"
)
