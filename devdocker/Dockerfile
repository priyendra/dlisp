FROM ubuntu:18.04

# Pre-install curl to setup bazel apt-get repo.
RUN apt-get update && apt-get install -y curl software-properties-common
RUN echo "deb [arch=amd64] " \
      "http://storage.googleapis.com/bazel-apt stable jdk1.8" | \
      tee /etc/apt/sources.list.d/bazel.list
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -

# Basic dev environment (c++, bazel)
# NOTE: We install boost as a system library. This is okay since we do not
# anticipate having to recompile boost. Also since boost is mostly header-only
# anyways, this will not be a problem anyway. Most other libraries will be
# installed as part of the 3rdparty toolchain framework.
RUN apt-get update && apt-get install -y \
      bazel \
      git \
      golang-go \
      man-db \
      sudo \
      vim
