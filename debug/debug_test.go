/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package debug_test

import (
	"io/ioutil"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/debug/debug"
	"github.com/sclevine/spec"
)

func testDebug(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx libcnb.BuildContext
		d   debug.Debug
	)

	it.Before(func() {
		var err error

		ctx.Buildpack.Info.Version = "test-version"

		ctx.Layers.Path, err = ioutil.TempDir("", "debug")
		Expect(err).NotTo(HaveOccurred())

		d = debug.NewDebug(ctx.Buildpack.Info)
	})

	it("contributes debug configuration", func() {
		layer, err := ctx.Layers.Layer("test-layer")
		Expect(err).NotTo(HaveOccurred())

		layer, err = d.Contribute(layer)
		Expect(err).NotTo(HaveOccurred())

		Expect(layer.Launch).To(BeTrue())
		Expect(layer.Profile["debug"]).To(Equal(`PORT=${BPL_DEBUG_PORT:=8000}
SUSPEND=${BPL_DEBUG_SUSPEND:=n}

printf "Debugging enabled on port %s" "${PORT}"
if [[ "${SUSPEND}" = "y" ]]; then
  printf ", suspended on start"
fi
printf "\n"

export JAVA_OPTS="${JAVA_OPTS} -agentlib:jdwp=transport=dt_socket,server=y,address=${PORT},suspend=${SUSPEND}"`))
	})
}
