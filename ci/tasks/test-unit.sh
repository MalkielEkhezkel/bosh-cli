#!/usr/bin/env bash

set -e -x

source ~/.bashrc

export PATH=/usr/local/ruby/bin:/usr/local/go/bin:$PATH
export GOPATH=$(pwd)/gopath

cd gopath/src/github.com/cloudfoundry/bosh-cli
bin/clean
bin/install-ginkgo
bin/test-prepare
bin/test-unit
