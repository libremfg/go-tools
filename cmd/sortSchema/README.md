# Sort Schema

> http proxy a graphql schema request and returns a sorted schema

Useful for committing schema files under version control. SchemaSort can also proxy the _service sdl request, but is not sorted.

```
Options: 
  --endpoint <host>  the endpoint to proxy the request to. Path is forwarded
	                   onto the host.

	--help:            shows this message

	--log:             logs the payload to the current directory before and
	                   after sorting.
```

## Usage

```
$ ./sort-schema-v1.0.0.-windows-amd64.exe
Listening on 0.0.0.0:8081

```

## Maintainer 

Tom Hollingworth <tom.hollingworth@spruiktec.com>