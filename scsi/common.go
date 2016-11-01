package scsi

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func listScsiDevices(dirname string) ([]string, error) {

	var scsiList []string
	sysPathStats, err := ioutil.ReadDir(dirname)
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

func getCounter(counterName string) (int64, error) {
	filePath := filepath.Join(sysPath, counterName)
	cnt, err := readHex(filePath)
	if err != nil {
		return 0, err
	}
	return cnt, nil
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

	//if err != nil {
	//	return 0, err
	//	}
	return number, nil

}
