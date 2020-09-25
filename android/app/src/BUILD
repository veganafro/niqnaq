load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_android_library")
load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_test")

android_binary(
    name = "app_binary",
    custom_package = "com.veganafro.app",
    manifest = "main/AndroidManifest.xml",
    deps = [
        "//app/src:app_lib",
    ],
)

kt_android_library(
    name = "app_lib",
    custom_package = "com.veganafro.app",
    manifest = "main/AndroidManifest.xml",
    resource_files = glob(["main/res/**"]),
    srcs = glob(["main/java/com/veganafro/app/*.kt"]),
    visibility = ["//:__pkg__"],
    deps = [
        "@maven//:androidx_appcompat_appcompat",
        "@maven//:androidx_constraintlayout_constraintlayout",
        "@maven//:androidx_core_core_ktx",
        "@maven//:androidx_fragment_fragment",
        "@maven//:androidx_lifecycle_lifecycle_common",
        "@maven//:androidx_lifecycle_lifecycle_viewmodel",
    ],
)

kt_jvm_test(
    name = "app_unit_test",
    srcs = glob(["test/java/com/veganafro/app/*.kt"]),
    test_class = "com.veganafro.app.ExampleUnitTest",
    deps = [
        "@maven//:junit_junit",
    ],
)

android_local_test(
    name = "app_android_test",
    # A manifest file is needed here because `androidx_test_monitor` package sets a `minSdkVersion`
    manifest = "androidTest/AndroidManifest.xml",
    test_class = "com.veganafro.app.ExampleInstrumentedTest",
    deps = [
        # Bazel doesn't currently support kt_android_local_test
        # The recommended best practice is to wrap a `kt_android_library` in an `android_local_test`
        "//app/src/androidTest/java/com/veganafro/app:app_android_test_lib",
        "@maven//:org_robolectric_robolectric",
        "@robolectric//bazel:android-all",
    ],
)
