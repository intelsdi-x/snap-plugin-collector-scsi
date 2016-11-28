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
	"os"
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	invalidValue = iota
	invalidEntry = iota
)

var (
	mockMts = []plugin.Metric{

		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "scsi", "*", "iodone_cnt"),
		},

		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "scsi", "*", "ioerr_cnt"),
		},
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "scsi", "*", "iorequest_cnt"),
		},
	}
	srcMockFile    = "/tmp/scsi_mock"
	srcMockFileInv = "/tmp/scsi_invalid_mock"
)

func TestScsiCollectorPlugin(t *testing.T) {

	//config := plugin.Config{}
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
		createMockFiles()
		//	defaultSrcFile = srcMockFile
		scsiCol := ScsiCollector{}
		cfg := plugin.Config{
			"sysPath": "/sys",
		}
		mts := []plugin.Metric{}

		for _, metric := range scsiMetricsTypes {

			metric := plugin.Metric{Namespace: plugin.NewNamespace(nsVendor, nsClass).AddDynamicElement("", "").AddStaticElement(metric)}
			mts = append(mts, metric)
			fmt.Println(mts)
		}

		So(len(mts), ShouldEqual, 3)
		metrics, err := scsiCol.GetMetricTypes(cfg)

		So(err, ShouldBeNil)
		So(metrics, ShouldNotBeEmpty)
		So(len(metrics), ShouldResemble, 3)
	})

	Convey("Collect cnt Metrics", t, func() {
		scsiCol := ScsiCollector{}
		mts := []plugin.Metric{}
		cfg := plugin.Config{
			"sysPath": "/sys",
		}
		for _, m := range scsiMetricsTypes {

			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(nsVendor, nsClass).AddStaticElements(m), Config: cfg})
			fmt.Println(mts)
		}
		metrics, err := scsiCol.CollectMetrics(mts)
		So(err, ShouldBeNil)

		So(len(metrics), ShouldResemble, 3)
		So(metrics[0].Data, ShouldNotBeNil)

	})
}

func createMockFiles() {
	deleteMockFiles()
	// 	mocked content of srcMockFile (kernel 2.6+)
	srcMockFileCont := []byte(`  8    0 test_scsi  0x234 `)
	f, _ := os.Create(srcMockFile)
	f.Write(srcMockFileCont)
}

func createInvalidMockFile(kind int) {
	os.Remove(srcMockFileInv)

	var srcMockFileContInv []byte

	switch kind {
	case invalidValue:
		srcMockFileContInv = []byte(`    8       0 test_scsi abc`)
		break

	case invalidEntry:
		srcMockFileContInv = []byte(`    1       2 unknown entry`)
		break

	default:
		srcMockFileContInv = []byte(``)
		break

	}

	f, _ := os.Create(srcMockFileInv)
	f.Write(srcMockFileContInv)

}

func deleteMockFiles() {
	os.Remove(srcMockFile)
	os.Remove(srcMockFileInv)

}
