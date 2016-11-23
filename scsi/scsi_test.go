// +build linux

/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2016 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scsi

import (
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestScsiCollectorPlugin(t *testing.T) {
	config := plugin.Config{}
	Convey("Create Scsi Collector", t, func() {
		scsiCol := ScsiCollector{}
		Convey("So Scsi should not be nil", func() {
			So(scsiCol, ShouldNotBeNil)
		})
		Convey("So Scsi should be of scsi type", func() {
			So(scsiCol, ShouldHaveSameTypeAs, ScsiCollector{})
		})
		Convey("scsiCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := scsiCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})

	Convey("Get Metric Scsi Types", t, func() {
		scsiCol := ScsiCollector{}
		var cfg = plugin.Config{}
		metrics, err := scsiCol.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 3)
	})

	Convey("Collect cnt Metrics", t, func() {
		scsiCol := ScsiCollector{}
		mts := []plugin.Metric{}
		for _, m := range scsiMetricsTypes {

			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(nsVendor, nsClass).AddStaticElements(m), Config: config})

		}
		metrics, err := scsiCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 6)
		So(metrics[0].Data, ShouldNotBeNil)

	})
}
