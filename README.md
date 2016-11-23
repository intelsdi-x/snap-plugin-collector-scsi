# snap-plugin-collector-scsi
Collects Linux SCSI statistics

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#operating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

 Plugin collects specified metrics from linux scsi

### System Requirements


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
This builds the plugin in `/build/rootfs`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

### Examples
Example running scsi plugin, passthru processor, and writing data into an csv file.

Documentation for snap file publisher can be found [here](https://github.com/intelsdi-x/snap)

In one terminal window, open the snap daemon :
```
$ snapd -t 0 -l 1
```
The option "-l 1" it is for setting the debugging log level and "-t 0" is for disabling plugin signing.

In another terminal window:

Load collector and processor and Publisher plugins
```
$ snapctl plugin load $SNAP_SCSI_PLUGIN/build/rootfs/snap-plugin-collector-scsi
$ snapctl plugin load $SNAP_PATH/build/plugin/snap-plugin-publisher-file
$ snapctl plugin load $SNAP_PATH/build/plugin/snap-plugin-processor-passthru
```

See available metrics for your system
```
$ snapctl metric list
```

Create a task file. For example, sample-scsi-task.json:

```
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/scsi/ioerr_cnt": {},
                "/intel/scsi/iodone_cnt": {},
                "/intel/scsi/iorequest_cnt": {},


            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```

## Documentation

*************Need to be update this area******************************************


### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type  |Description
----------|-----------|-----------|-----------|-----------|
/intel/scsi//iodoneCnt | int64|
/intel/scsi//ioerrorCnt | int64|
/intel/scsi//iorequestCnt/ int64|



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
