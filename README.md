# health-sidecar

A simple healthcheck sidecar container with Pager Duty integration. It is intended to run alongside a service to monitor its status and trigger a Pager Duty event when it becomes unhealthy.

# Configuration

The behaviour is configurable through environment variables.
| Variable | Description | Default |
|---|---|---|
| HTTP_ENDPOINT | Endpoint to run the query on. | http://127.0.0.1 |
| DELAY | Interval between healtchecks, in seconds. | 20 |
| PAGERDUTY_SERVICEKEY | Pager Duty integration key for the service. | N/A |
