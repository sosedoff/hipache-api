# hipache-api

HTTP interface for [Hipache](https://github.com/hipache/hipache)

## Install

```
go get sosedoff/hipache-api
```

Or use one of the binaries from [Github Releases](https://github.com/sosedoff/hipache-api/releases).

## Usage

Start http server:

```
hipache-api
```

By default hipache-api will try to connect to redis server on `localhost:6379`. 
To change that you can specify extra environment variables:

```
REDIS_HOST=host REDIS_PORT=port hipache-api
```

Enable api key authentication:

```
API_KEY=key hipache-api
```

### Docker

There's a prebuilt Docker image available. First, pull the image:

```
docker pull sosedoff/hipache-api
```

Then start container:

```
docker run -d -p 5000 -e REDIS_HOST=host -e REDIS_PORT=port -e API_KEY=key sosedoff/hipache-api
```

## API

```
   GET /frontends
   GET /frontends/:name
  POST /frontends?host=site.com&backends=http://host1:port,http://host2:port
  POST /frontends/:name?backends=http://host1:port,http://host2:port
DELETE /frontends/:name
DELETE /frontends/:name/backend?backend=http://host1:port
  POST /flush
```

## Build

```
make build   # Make development build
make release # Build binaries for linux/osx
```

## License

The MIT License (MIT)

Copyright (c) 2014 Dan Sosedoff, <dan.sosedoff@gmail.com>