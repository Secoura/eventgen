# EventGen

This is an event generator for S that can generate logs using a template.

### How Do I Use It?

1. Build Eventgen

```
go build -o eventgen main.go
```

2. Run Eventgen

Currently there are only 2 provided templates, `combined` and `common` log formats for web events:

```
./eventgen web -t <template>
```

### Configuration

The `config.yaml` file holds all the config for EventGen.

```
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

You can build your own template by placing the template file under the `templates/web` directory.

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
```