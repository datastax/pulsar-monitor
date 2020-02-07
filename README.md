# Operation Monitoring for Pulsar

This is a ops monitoring tool for
- monitoring Pulsar admin tenants endpoint
- measure message latency from producing to consuming
- report and monitor heartbeat with OpsGenie
- alert on Slack

This is a data driven tool. The configuraion is a json file. Here is a [template](../config/runtime_template.json).
The configuration json file can be specified in the overwrite order of 
- an environment variable `PULSAR_OPS_MONITOR_CFG`
- an command line argument `./pulsar-monitor -config /path/to/pulsar_ops_monitor_config.json`
- A default path to `../config/runtime.json`

## Docker
The runtime.json file must be mounted as /app/runtime.json

This runs a multi stage build that produces a 18MB docker image.
```
$ sudo docker build -t pulsar-ops-monitor .
```

Run docker container with Pulsar CA certificate and expose Prometheus metrics for collection.

``` bash
$ sudo docker run -d -it -v /home/ming/go/src/gitlab.com/operation-monitor/config/runtime.yml:/config/runtime.yml -v /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem:/etc/ssl/certs/ca-bundle.crt -p 8080:8080 --name=pulsar-monitor pulsar-ops-monitor
```

## Prometheus
This program exposes a Prometheus `\metrics` endpoint to allow measured Pulsar latency to be collected by Prometheus.
