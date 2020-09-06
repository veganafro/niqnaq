load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_test")

kt_jvm_test(
    name = "app_unit_test",
    srcs = glob(["java/com/veganafro/mono/*.kt"]),
    test_class = "com.veganafro.mono.ExampleUnitTest",
    deps = [
        "@maven//:junit_junit",
    ],
)
