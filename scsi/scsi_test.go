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
	"fmt"
	"regexp"
	"testing"

	. "github.com/goconvey/convey"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"
)

func TestScsi(t *testing.T) {
	Convey("Meta should return metadata for the plugin", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, pluginName)
		So(meta.Version, ShouldResemble, pluginVersion)
		So(meta.Type, ShouldResemble, plugin.CollectorPluginType)
	})

	Convey("Create Scsi", t, func() {
		sc := New()
		Convey("So sc should not be nil", func() {
			So(sc, ShouldNotBeNil)
		})
		Convey("So sc should be of scsi type", func() {
			So(sc, ShouldHaveSameTypeAs, &Scsi{})
		})
		Convey("sc.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := sc.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
	})

	Convey("Get Metrics ", t, func() {
		sc := New()
		var cfg = plugin.ConfigType{}

		Convey("So should return 3 types of metrics", func() {
			metrics, err := sc.GetMetricTypes(cfg)
			So(len(metrics), ShouldBeGreaterThan, 1)
			So(err, ShouldBeNil)
		})

		Convey("So should check namespace", func() {
			metrics, err := sc.GetMetricTypes(cfg)

			drvNamespace := metrics[0].Namespace().String()

			fmt.Println(drvNamespace)

			drv := regexp.MustCompile(`/intel/scsi/iodone_cnt`)
			So(true, ShouldResemble, drv.MatchString(drvNamespace))
			So(err, ShouldBeNil)

			drvNamespace1 := metrics[1].Namespace().String()
			drv1 := regexp.MustCompile(`/intel/scsi/ioerr_cnt`)
			So(true, ShouldResemble, drv1.MatchString(drvNamespace1))
			So(err, ShouldBeNil)

			drvNamespace2 := metrics[2].Namespace().String()
			drv2 := regexp.MustCompile(`/intel/scsi/iorequest_cnt`)
			So(true, ShouldResemble, drv2.MatchString(drvNamespace2))
			So(err, ShouldBeNil)
		})

	})

	Convey("Collect Metrics", t, func() {
		sc := &Scsi{}

		cfgNode := cdata.NewNode()

		//	pwd, err := os.Getwd()
		//	sysPath = filepath.Join(pwd, "/sys/")
		sd := "0:0:0:0"

		//So(err, ShouldBeNil)

		//	So(err, ShouldBeNil)

		Convey("So should get iodone_cnt metrics", func() {

			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace(sd, "iodone_cnt"),
				Config_:    cfgNode,
			}}
			collect, err := sc.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			fmt.Println(collect[0].Data_)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType int64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)

		})
		Convey("So should get ioerror_cnt metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace(sd, "ioerr_cnt"),
				Config_:    cfgNode,
			}}
			collect, err := sc.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType int64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get iorequest_cnt metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace(sd, "iorequest_cnt"),
				Config_:    cfgNode,
			}}
			collect, err := sc.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType int64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
	})
}
