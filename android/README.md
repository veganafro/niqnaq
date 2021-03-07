# Summary

Resources related to the Mono Android app can be found here.

## Structure

```
android
├── app
│   ├── proguard-rules.pro
│   ├── .gitignore
│   ├── BUILD
│   └── src
│       ├── androidTest
│       │   ├── AndroidManifest.xml
│       │   └── java
│       │       └── com.veganafro.app
│       │           ├── BUILD
│       │           └── ExampleInstrumentedTest.kt
│       ├── test
│       │   └── java
│       │       └── com.veganafro.app
│       │           └── ExampleUnitTest.kt
│       └── main
│           ├── AndroidManifest.xml
│           ├── java
│           │   └── com.veganafro.app
│           │       └── MainActivity.kt
│           └── res
│               ├── layout
│               │   └── activity_main.xml
│               └── values
│                   ├── colors.xml
│                   ├── strings.xml
│                   └── styles.xml 
├── .gitignore
└── WORKSPACE
```

Each feature should be a new module defined as a top level directory with a `BUILD` file. Each module should follow the
structure outlined above.

* androidTest
   
   `androidTest` contains tests that need the Android library and/or Robolectric to run, including `Activity`
   and `Fragment` tests that need to launch either component. Bazel doesn't currently support kt_android_local_test. So,
   the recommended best practice is to wrap a `kt_android_library` in an `android_local_test`.
* main
   
   `main` contains the module's source code.
* test
   
   `test` contains unit tests that can be run without the Android library and/or Robolectric, including `ViewModel`
   tests that focus on business logic.
 

## Getting started

### Setting up the dev environment

Start by installing [`brew`](https://brew.sh/)

Next, run the following commands that install the Bazel build tool and Java 8:

```
$ brew install bazel
$ brew tap adoptopenjdk/openjdk
$ brew install --cask adoptopenjdk8
$ brew tap homebrew/core
```

Once those prerequisites have been installed, add the following lines to `.bash_profile`:

```
export JAVA_HOME=$(/usr/libexec/java_home -v1.8)
source $(brew --prefix)/etc/bash_completion.d
```

Create a directory `~/android-sdk` and add the following lines to `.bash_profile`:

```
export ANDROID_HOME=~/android-sdk
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/21.3.6528147
export PATH=$JAVA_HOME/bin:$ANDROID_HOME/emulator:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools:$ANDROID_HOME/tools/bin:$PATH
```

Now, download the Android SDK command line tools from [this](https://developer.android.com/studio) link. Extract those
tools from the zip file and navigate to `cmdline-tools/bin/`.

Run the following commands to install the required Android development tools:

```
$ ./sdkmanager --install "platform-tools" "platforms;android-30" --sdk_root=~/android-sdk
$ ./sdkmanager --install "ndk;21.3.6528147" --sdk_root=~/android-sdk
$ ./sdkmanager --install "emulator" "patcher;v4" "sources;android-30" "build-tools;30.0.3" --sdk_root=~/android-sdk
```

Make sure that the tools were properly installed by running `sdkmanager --list`. Also verify that there were no redundant
installations.

### Running the app

After installing `bazel`, confirm that everything is working by navigating to the `android` directory and executing:

```
$ bazel sync
$ bazel info workspace
```

If you see the current directory path printed, everything is working correctly.

Now, follow these steps to correctly configure IntelliJ:

   1. Open the IntelliJ app
   2. Select `Configure` at the bottom right of the landing page
   3. Select `Plugins` from the dropdown
   4. Install the `Bazel` plugin
   5. Restart the IntelliJ app
   6. Select the new `Import Bazel Project` option on the landing page
   7. Select the `android` directory and navigate through the rest of the setup

Once that step is complete, it should be easy enough to build and install the app to a connected device or emulator in two steps:

```
$ bazel build //app:app_binary
$ bazel mobile-install //app:app_binary --start_app
```

Incremental builds can be sped up by executing the second command with a replaced target (the target that has been modified)
and an additional `--incremental` flag.

## Testing

Each module should be well tested, and tests can be executed with this `bazel` command:

```
$ bazel test //some/target:name
```

Each tested class corresponds to a `kt_jvm_test` (for unit tests) or an `android_local_test` (for integration tests).
