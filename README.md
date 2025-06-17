# Ping Connection Database

## Getting started
use `go get`:

```sh
go get -u github.com/varnion-rnd/ping-connection
```

### Running in Elastic
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/elastic"
)

func Ping(db *elasticsearch.Client) error {
	return pingconnection.Ping(db)
}
```

### Running in Influx
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/influx"
)

func Ping(ctx context.Context, db db influxdb2.Client) error {
	return pingconnection.Ping(db)
}
```

### Running in Mongo
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/mongo"
)

func Ping(ctx context.Context, db *mongo.Client) error {
	return pingconnection.Ping(db)
}
```

### Running in MySQL
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/mysql"
)

func Ping(dsn string) error {
	return pingconnection.Ping(db)
}
```

### Running in Postgres
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/postgres"
)

func Ping(dsn string) error {
	return pingconnection.Ping(db)
}
```

### Running in Redis
```go
import (
	pingconnection "github.com/varnion-rnd/ping-connection/redis"
)

func Ping(ctx context.Context, db *redis.Client) error {
	return pingconnection.Ping(db)
}
```