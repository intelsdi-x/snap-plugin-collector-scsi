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
	"path/filepath"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (

	// PLUGIN scsi collector namespace part
	Name = "scsi"
	// VERSION of scsi info plugin
	Version = 1
	// VENDOR namespace part
	nsVendor = "intel"
	// FS namespace part
	nsClass = "scsi"
)

var (
	//sysPath source of data for metrics
	scsiPath = "bus/scsi/devices/"
	// name of available metrics
	scsiMetricsTypes = []string{"iodone_cnt", "ioerr_cnt", "iorequest_cnt"}
)

type ScsiCollector struct {
}

//GetConfigPolicy returns a ConfigPolicy
func (ScsiCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {

	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"intel", "scsi"},
		"sysPath", false, plugin.SetDefaultString("/sys"))
	return *policy, nil
}

// GetMetricTypes returns the metric types
func (ScsiCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {

	mts := []plugin.Metric{}
	for _, metric := range scsiMetricsTypes {
		metric := plugin.Metric{Namespace: plugin.NewNamespace(nsVendor, nsClass).AddDynamicElement("device_id", "Id of the scsi device").AddStaticElement(metric)}
		mts = append(mts, metric)
	}
	return mts, nil
}

// CollectMetrics returns metrics
func (ScsiCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	sysPathConf, err := mts[0].Config.GetString("sysPath")
	if err != nil {
		return nil, err
	}

	sysPath := filepath.Join(sysPathConf, scsiPath)

	for _, m := range mts {
		lastElement := len(m.Namespace.Strings()) - 1
		cnt := m.Namespace.Strings()[lastElement]
		scsiList, err := listScsiDevices(sysPath)
		if err != nil {
			return nil, err
		}

		metric, err := getCounter(cnt, scsiList, m.Namespace)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric...)

	}
	return metrics, nil

}
