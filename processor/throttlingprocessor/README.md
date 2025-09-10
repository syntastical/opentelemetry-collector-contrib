# Log Throttling Processor

Left mostly blank while in review.

## Example Config

The following example sets up a log throttling processor that limits logs to 1 log every 6 seconds per log file, but only for logs that have an attribute `service.name` with the value `service1`.  The count is bucketed by the expression `attributes["log.file.name"]`, so the throttling will occur on a per-file basis. 

```yaml
receivers:
  filelog:
    include: [ ./example/*.log ]
processors:
  log_throttle:
    interval: 6
    threshold: 1
    key_expression: 'attributes["log.file.name"]'
    conditions:
      - 'attributes["service.name"] == "service1"'
exporters:
  debug:
    verbosity: detailed
service:
  pipelines:
    logs:
      receivers: [ filelog ]
      processors: [ log_throttle ]
      exporters: [ debug ]
```
