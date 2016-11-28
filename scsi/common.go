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
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func listScsiDevices(dirName string) ([]string, error) {

	var scsiList []string
	sysPathStats, err := ioutil.ReadDir(dirName)
	if err != nil {
		return scsiList, err
	}
	for _, dir := range sysPathStats {

		dvre := regexp.MustCompile(`^[0-9]:.?:.?:.?`)
		dirName := dir.Name()
		if dvre.MatchString(dirName) {
			scsiList = append(scsiList, dirName)
		}

	}
	return scsiList, nil

}

func getCounter(counterName string, scsiList []string, ns plugin.Namespace) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	for _, dev := range scsiList {
		newNs := make([]plugin.NamespaceElement, len(ns))
		copy(newNs, ns)
		newNs[2].Value = dev

		filePath := filepath.Join(scsiPath, dev, counterName)
		cnt, err := readHex(filePath)
		if err != nil {
			return metrics, nil
		}
		metric := plugin.Metric{
			Namespace: ns,
			Data:      cnt,
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func readHex(filename string) (int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// The int files that this is concerned with should only be one liners.
	line, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	i := strings.TrimSpace(line)

	number, _ := strconv.ParseInt(i, 10, 64)

	return number, nil

}
