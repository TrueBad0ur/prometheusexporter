# Custom simple template for prometheus exporter and grafana dashboard



# Run
```bash
go run main.go

cd prometheus_grafana
docker compose up
```

# Connect
grafana - ```0.0.0.0:3000 admin:admin```

prometheus - ```0.0.0.0:9090```

exporter - ```0.0.0.0:9100/metrics```
