[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-scsi.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-scsi )
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-scsi)](http://goreportcard.com/report/intelsdi-x/snap-plugin-collector-scsi)

This plugin collects metrics from scsi.  

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#operating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)
7. [Thank you](#thank-you)

## Getting Started

 Plugin collects specified metrics from linux scsi

### System Requirements

* golang 1.7+ - needed only for building
* This Plugin compatible with kernel > 2.6
* Linux/x86_64

### Operating systems

All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download use plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [Github Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-scsi

Clone repo into `$GOPATH/src/github/intelsdi-x/`:
```
$ git clone https://github.com/<yourGithubID>/intelsdi-x/snap-plugin-collector-scsi
```
Build the plugin by running make in repo:
```
$ make
```
This builds the plugin in `/build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Load the plugin and create a task, see example in [Examples](#examples).

## Documentation
### Examples

Example of running snap scsi collector and writing data to file.

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapteld -l 1 -t 0 &`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-scsi/latest/linux/x86_64/snap-plugin-collector-scsi
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-scsi
$ snaptel plugin load snap-plugin-publisher-file

Create a task manifest file  (exemplary files in [examples/tasks/] (examples/tasks/)):
```yaml
---
  version: 1
  schedule:
    type: "simple"
    interval: "1s"
  max-failures: 10
  workflow:
    collect:
      metrics:
         /intel/scsi/*/iodone_cnt: {}
         /intel/scsi/*/ioerr_cnt: {}
         /intel/scsi/*/iorequest_cnt: {}
      publish:
        - plugin_name: "file"
          config:
            file: "/tmp/scsi_metrics"
```
Download an [example task file](https://github.com/intelsdi-x/snap-plugin-collector-scsi/blob/master/examples/tasks/) and load it:

```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-scsi/master/examples/tasks/scsi-file.yaml
$ snaptel task create -t scsi-file.yaml
Using task manifest to create task
Task created
ID: 250323af-12b0-4bf8-a526-eb2ca7d8ae32
Name: Task-250323af-12b0-4bf8-a526-eb2ca7d8ae32
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
$ snaptel task watch 250323af-12b0-4bf8-a526-eb2ca7d8ae32
```

This data is published to a file `/tmp/scsi_metrics` per task specification

Stop task:
```
$ snaptel task stop 250323af-12b0-4bf8-a526-eb2ca7d8ae32
Task stopped:
ID: 250323af-12b0-4bf8-a526-eb2ca7d8ae32
```

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-scsi/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-scsi/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-scsi/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Authors: [Marcin Spoczynski](https://github.com/sandlbn/) and [Ramesh Raju](https://github.com/rraju2/)

## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
