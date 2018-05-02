<img height="100" src="../images/dummy.svg"/>

### Dummy-Logsink is a Log Sink for Fluent-bit

The Dummy-Logsink service receives a batch of raw log messages from fluent-bit and metadata via HTTP Post and responds with an status of 200 OK.
Optionally, it will delay responding by the number of seconds specified in the environment variable `DELAY`.

Intended for testing purposes.

### Contract

A single log message that comes in will minimally look something like this:
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

### Run in Docker

In one terminal window:
```
docker run -it -p 3000:3000 stevenacoffman/dummy-logsink
```

In another terminal window:
```
./test.sh
```

### Image Attribution
Crash Dummy by Petr Papasov from the Noun Project

Log by Arthur Schmitt from the Noun Project

Sink by Arthur Shlain from the Noun Project

Go Gopher created by Takuya Ueda