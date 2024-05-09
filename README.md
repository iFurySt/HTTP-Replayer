# Traffic Replayer
Traffic Replayer is a simple traffic replay tool.

[Project Demo Video](https://github.com/iFurySt/TrafficReplayer/assets/16201837/07752d5a-9c8f-4a80-b874-6222ea0099cf)

## Quick Start
```bash
tp -R 1/m -N lo0 -M POST -P 8888 -U /api/v1/task http://localhost:8899/api/v2/task http://localhost:8898/api/v3/task
```

## License
Traffic Replayer is licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.
