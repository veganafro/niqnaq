# Start from the latest bazel image
FROM l.gcr.io/google/bazel:3.5.0

LABEL maintainer="Jeremy Muhia <jmuhia80@gmail.com>"

# Set the working directory in the container
WORKDIR /go/src/github.com/veganafro/mono/

# Copy local files into the container's working directory
COPY . .

# Read the build arguments from the yml file
ARG dummy_host
ARG port
ARG project

# Set the environment variables from the build arguments
ENV DUMMY_HOST $dummy_host
ENV PORT $port
ENV PROJECT $project

## Build the binary with bazel
RUN bazel build $project

# Run the service using bazel
ENTRYPOINT ["/bin/bash", "-c", "bazel run ${PROJECT}"]

