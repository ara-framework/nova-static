# Nova Static
Nova Static is an tool to Server Side Include (SSI) views in static sites using [Hypernova](https://github.com/airbnb/hypernova)

## Environment Variables

```
HYPERNOVA_BATCH=http://hypernova:3000/batch
STATIC_FOLDER=./public
```

## Install

Dowload package
```
go get github.com/ara-framework/nova-static/nova-static
```

Install package
```
go install github.com/ara-framework/nova-static/nova-static
```

Export Path

```
export PATH="$PATH:$GOPATH/bin"
```

## Run

```
$ nova-static
```