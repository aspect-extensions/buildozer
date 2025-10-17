# buildozer tasks

Buildozer is a machine-editor for BUILD files, along with some support for MODULE.bazel.

See https://github.com/bazelbuild/buildtools/blob/main/buildozer/README.md

This repository provides a WASM binding for the Buildozer Go library,
so it can be invoked from AXL's starlark context.

## print task

Run `aspect print` to get a syntactic view of a BUILD file.
This takes place PRIOR to the loading phase, so macros are not expanded.
This is useful to quickly search a repository without triggering a slow fetch and loading phase.

## install_ruleset task

Installs a specific version of rules_oci by updating the MODULE.bazel file.
