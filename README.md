# Eventgen

This is an event generator for generating logs using a template.

### Getting Started

1. Build Eventgen

```
go build -o eventgen main.go
```

2. Run Eventgen

Currently there are only 2 provided templates, `combined` and `common` log formats for web events:

```
./eventgen generate
```

### Configuration

The `config.yaml` file holds all the config for EventGen.

```
# This is the template to use to generate events
template: web/combined

# This is the start time of the events that are generated
# It has to follow the format specified in the example below
# If the start time is not configured, it defaults to the current time
startTime: 08/Apr/2018:15:04:05 +1000

# This is the time duration for the events we want to generate
# ie. If this is set to 15 minutes, the events generated will
# Have a randomized timestamp between the start time and the start time + duration
# Duration can be specified in the format of \d+d \d+h \d+m
duration: 15m

# The number of events we want to generate
noOfEvents: 1000
```

### Custom Templates

You can build your own template by placing the template file under the `templates` directory.

There are only a few processors that are available as of now:

```
%HostName% - A generated IP address
%HttpMethod% - One of GET, POST, PUT, DELETE
%Referer% - A random referer, or -
%ResourceSize% - A random 5 digit string
%StatusCode% - A HTTP status code
%Timestamp:<Format>% - A generated timestamp that will follow the provided format
%Url% - A generated URL
%UserAgent% - A generated user agent
%Lookup:<File>% - Random lookup of value from a list of newline-separated text file
```

# Using in Secoura

Add the following example configuration into the `plugins.yml` file to enable Eventgen in Secoura:

```yaml
eventgen:
  - eventgen-collector:
      enabled: true
      template: web/combined-csv
      duration: 15m
      number_of_events: 100
      delimiter: --SECOURA_END_OF_EVENT--
      overrides:
        interval: 1000
        line_breaker: --SECOURA_END_OF_EVENT--
```
