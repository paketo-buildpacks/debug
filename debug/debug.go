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

package debug

import (
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type Debug struct {
	LayerContributor libpak.LayerContributor
	Logger           bard.Logger
}

func NewDebug(info libcnb.BuildpackInfo) Debug {
	return Debug{LayerContributor: libpak.NewLayerContributor("Debug", info)}
}

func (d Debug) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	d.Logger.Body(bard.FormatUserConfig("BPL_DEBUG_PORT", "the port the JVM will listen on", "8000"))
	d.Logger.Body(bard.FormatUserConfig("BPL_DEBUG_SUSPEND", "whether the JVM will suspend on startup", "n"))

	d.LayerContributor.Logger = d.Logger

	return d.LayerContributor.Contribute(layer, func() (libcnb.Layer, error) {
		layer.Profile.Add("debug", `PORT=${BPL_DEBUG_PORT:=8000}
SUSPEND=${BPL_DEBUG_SUSPEND:=n}

printf "Debugging enabled on port %%s" "${PORT}"
if [[ "${SUSPEND}" = "y" ]]; then
  printf ", suspended on start"
fi
printf "\n"

export JAVA_OPTS="${JAVA_OPTS} -agentlib:jdwp=transport=dt_socket,server=y,address=${PORT},suspend=${SUSPEND}"`)

		layer.Launch = true
		return layer, nil
	})
}

func (Debug) Name() string {
	return "debug"
}
