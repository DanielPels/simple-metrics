# simple-metrics
Simple tool to add metrics for your go program.
- Export to json
- Use ":" as event seperator


Single event:
```sh
metrics.Event("Hello")
```

Event with 1 depth:
```sh
metrics.Event("Deep:Down")
```

Exports to Json:
```sh
{"Deep":{"Down":1},"Hello":1}
```

Clear all events:
```sh
metrics.ClearEvents()
```
