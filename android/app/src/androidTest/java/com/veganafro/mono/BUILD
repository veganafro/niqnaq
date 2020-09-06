load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_android_library")

kt_android_library(
    name = "app_android_test_lib",
    srcs = glob(["*.kt"]),
    custom_package = "com.veganafro.mono.test",
    visibility = ["//app/src/androidTest:__pkg__"],
    deps = [
        "@maven//:androidx_test_monitor",
        "@maven//:androidx_test_ext_junit",
    ],
)
