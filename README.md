# simple-metrics
Simple tool to add metrics for your go program.
- Import and use
- Export to json
- Use ":" as event seperator
- Can use EventValue to store a float64 to slice as event(could be used to calculate average or total sum)


Single event:
```sh
metrics.Event("Hello")
```

Event with depth:
```sh
metrics.Event("Deep:Down")
metrics.Event("Deep:Down")
metrics.Event("Deep:Below:TheOcean")
```

Event with value:
```sh
metrics.EventValue("life", 42)
metrics.EventValue("life", 0.76)
metrics.EventValue("life", -238)
```

Eventvalue with depth:
```sh
metrics.EventValue("Use:Simple:Metrics", 808)
metrics.EventValue("Use:Simple:Metrics", -23.027)
```

Cannot overwrite value if already set:
```sh
metrics.Event("Hello")
metrics.Event("Hello:Error")
"unable to set map, already int or slice - event: Hello"
```

Clear all events:
```sh
metrics.ClearEvents()
```

Exports to Json:
```sh
[]bytes, error := metrics.ExportJson()
Export:
{
  "Deep": {
    "Below": {
      "TheOcean": 1
    },
    "Down": 2
  },
  "Hello": 1,
  "Use": {
    "Simple": {
      "Metrics": [
        808,
        -23.027
      ]
    }
  },
  "life": [
    42,
    0.76,
    -238
  ]
}
```
