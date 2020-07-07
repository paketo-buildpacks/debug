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
	"fmt"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/libpak/sherpa"

	_ "github.com/paketo-buildpacks/debug/debug/statik"
)

type Debug struct {
	LayerContributor libpak.LayerContributor
	Logger           bard.Logger
}

func NewDebug(info libcnb.BuildpackInfo) Debug {
	return Debug{LayerContributor: libpak.NewLayerContributor("Debug", info)}
}

//go:generate statik -src . -include *.sh

func (d Debug) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	d.LayerContributor.Logger = d.Logger

	return d.LayerContributor.Contribute(layer, func() (libcnb.Layer, error) {
		s, err := sherpa.StaticFile("/debug.sh")
		if err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to load debug.sh\n%w", err)
		}

		layer.Profile.Add("debug.sh", s)

		layer.Launch = true
		return layer, nil
	})
}

func (Debug) Name() string {
	return "debug"
}
