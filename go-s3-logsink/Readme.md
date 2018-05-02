<img height="100" src="../images/GoS3LogSink.svg"/>

### Go-S3-Logsink is a Log Sink for S3

The Go-S3-Logsink service processes a batch of raw log messages and metadata and sends them to the appropriate S3 bucket and prefix. See the k8s directory for a sample kubernetes deployment including environment variables.

### Contract

A single log message that comes in will minimally look like this:
```
[{"date":1521039841.000000, "container_id":"03ba6033fda747aa199082ebbcae6e6019af6f3f5a09326143cd4894ae9809e6", "container_name":"/talk-other-format", "source":"stdout", "log":"[#|Wed Mar 14 15:04:01 UTC 2018|INFO|count=1;flag=foo|It's a message - with some stuff|#]", "topic":"tomcat"}]
```

For readability:
```
[
  {
    "date": 1521039841,
    "container_id": "03ba6033fda747aa199082ebbcae6e6019af6f3f5a09326143cd4894ae9809e6",
    "container_name": "/talk-other-format",
    "source": "stdout",
    "log": "[#|Wed Mar 14 15:04:01 UTC 2018|INFO|count=1;flag=foo|It's a message - with some stuff|#]",
    "topic": "tomcat"
  }
]
```

### Image Attribution
Log by Arthur Schmitt from the Noun Project

Sink by Arthur Shlain from the Noun Project

Go Gopher created by Takuya Ueda

AWS S3 Logo by AWS
