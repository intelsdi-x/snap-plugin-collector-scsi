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

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (

	// PLUGIN scsi collector namespace part
	pluginName = "scsi"
	// VERSION of scsi info plugin
	pluginVersion = 1
	// Type of plugin
	pluginType = plugin.CollectorPluginType
	// VENDOR namespace part
	nsVendor = "intel"
	// FS namespace part
	nsClass = "scsi"
	nsType  = "filesystem"
	SysPath = "sys_path"
	nVendor = "vendor"
	nModel  = "model"
)

var (
	//sysPath source of data for metrics
	sysPath = "/sys/bus/scsi/devices/"
	// prefix in metric namespace
	// prefix in metric namespac
	namespacePrefix = []string{nsVendor, nsClass, nsType}
)

// name of available metrics
var scsiMetricsTypes = []string{"iodone_cnt", "ioerr_cnt", "iorequest_cnt"}

//Scsi holds scsi statistics
type Scsi struct {
	data    int64
	sysPath string
}

// NewScsi creates new instance of plugin and returns pointer to initialized object.
func New() *Scsi {
	s := &Scsi{}
	return s
}

// Meta returns plugin meta data
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		pluginName,
		pluginVersion,
		plugin.CollectorPluginType,
		[]string{},
		[]string{plugin.SnapGOBContentType},
		plugin.ConcurrencyCount(1),
	)
}

// GetConfigPolicy returns a ConfigPolicy
func (sc *Scsi) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	rule, _ := cpolicy.NewStringRule("sys_path", false, "/sys/")
	node := cpolicy.NewPolicyNode()
	node.Add(rule)
	cp.Add([]string{nsVendor, pluginName}, node)
	return cp, nil
}

// GetMetricTypes returns list of exposed disk stats metrics
func (sc *Scsi) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}
	for _, metric := range scsiMetricsTypes {
		metric := plugin.MetricType{Namespace_: core.NewNamespace(nsVendor, nsClass, metric)}
		mts = append(mts, metric)
	}
	return mts, nil
}

// CollectMetrics returns metrics
func (sc *Scsi) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	metrics := []plugin.MetricType{}
	/*
		scsiList, err := listScsiDevices(sysPath)
		if err != nil {
			return nil, err
		}*/

	for _, m := range mts {
		ns := m.Namespace()
		//	cnt := m.Namespace().Strings[-1]
		cnt := m.Namespace().String()
		cntr, err := getCounter(cnt)
		if err != nil {
			return nil, err
		}
		metric := plugin.MetricType{
			Namespace_: ns,
			Data_:      cntr,
		}

		metrics = append(metrics, metric)
	}
	fmt.Println(metrics)
	return metrics, nil

}
