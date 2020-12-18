# Altran Golang ClientAPI



Following the requirements asked, this service is responsible to fetch all the file ports (ports.json) objects by the request of the endpoint "/ports", which is provided with the mandatories query parameters: skip and limit. Every time when each request is made on this endpoint, the resulted records are compiled and sent to service [PortDomainService](https://github.com/eliasbosco/altran-golang-PortDomainService) via [gRPC](https://github.com/grpc/grpc-go/blob/d79063fdde284ef7722591e56c72143eea59c256/examples/features/debugging/client/main.go) protocol with protobuf to envelope the data, which insert or update every record on SQLite database attached.

There's another endpoint (/ports-db) which is responsible to fetch all the database records and provide it as response json. There are some query parameters: skip and limit, in case to slice part of the result and port_id, to return only the port related to it's ID.

## Containers

All the service is dockerized and provided with docker compose script to facilitate the deploy step, all the commands are described below:

To bring the things on:

```
$ docker-compose up -d --build
```

After it's up, access:

(to response with port.json file results)

- [http://localhost:10000/ports?skip=0&limit=100](http://localhost:10000/ports?skip=0&limit=100)

(to response with sqlite results)

- http://localhost:10000/ports-db?port_id=ABCD
- http://localhost:10000/ports-db?skip=0&limit=100

To shut the service off:

```
$ docker-compose down
```

[^Notice: every environments variables are set on &quot;.env&quot; file, at the project root.]: 

## Resources

All the computational resource are limited to 50M memory, even lessly as mentioned on the pre-requisits.