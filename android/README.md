# Summary

Resources related to the Mono Android app can be found here.

## Getting started

It's assumed that the Android environment (`$ANDROID_HOME` and `ANDROID_NDK_HOME`) have been correctly configured.

This Android project is built using the [`bazel`](https://github.com/bazelbuild/bazel) build system. Install it with the following command:

```
$ brew install bazel
```

After installing `bazel`, confirm that everything is working by navigating to the `android` directory and executing:

```
$ bazel info workspace
```

If you see the current directory path printed, everything is working correctly.

Now, follow these steps to correctly configure IntelliJ:

   1. Click on the IntelliJ app
   2. Select `Configure` at the bottom right of the landing page
   3. Select `Plugins` from the dropdown
   4. Install the `Bazel` plugin
   5. Restart the IntelliJ app
   6. Select the new `Import Bazel Project` option on the landing page
   7. Select the `android` directory and navigate through the rest of the setup

Once that step is complete, it should be easy enough to build and install the app to a connected device or emulator in two steps:

```
$ bazel build //:mono
$ bazel mobile-install //:mono --start_app
```

Incremental builds can be sped up by executing the second command with a replaced target (the target that has been modified) and an additional `--icremental` flag.

## Testing

Each module should be well tested, and tests can be executed with this `bazel` command:

```
$ bazel test //some/target:name
```

