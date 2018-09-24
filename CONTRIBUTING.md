# Contributing

## Dev Build

Some tools are included to make contributing easier.

### make build/dev

```sh
make build/dev
```

The HTML and javascript are located in the /ui/ folder, and packaged up into byte code with go-bindata when the app is built.  If you're working on the UI, this target will bundle the UI assets in a linked mode so you can update them without having to rebuild / reload the server.

### make run/debug

```sh
make run/debug
```

This sets the DEBUG=true environment variable which outputs additional information.  You can use the debug logger in your own contirbutions to help troubleshoot issues:

```go
debug.Println("Checkpoint myFunc")
```

## Guidelines

* Please include all applicable unit tests.
* Please update the sequence diagram (puml/cloudtables.puml)
* Please make sure the code will pass the linter

### Errors

Please wrap all errors in a meaningful message before returning.  See examples in code.

### Logs

Please use one of the two loggers:

**Standard** logger for INFO level messages.  This should include information that is useful to the end user.

```go
log.Println("Request completed successfully.")
```

**Debug** logger for messages useful to the person updating the code.

```go
for _, element := range response {
    debug.Printf("Key: %s, Value: %s", element.Key, element.Value)
}
```

Thanks for your interest in this project.