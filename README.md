# `gcr.io/paketo-buildpacks/debug`
The Paketo Debug Buildpack is a Cloud Native Buildpack that configures debugging for JVM applications.

## Behavior
This buildpack will participate if all the following conditions are met

* `$BP_DEBUG_ENABLED` is set

The buildpack will do the following:

* Contribute debug configuration to `$JAVA_TOOL_OPTIONS`

## Configuration
| Environment Variable | Description
| -------------------- | -----------
| `$BP_DEBUG_ENABLED` | Whether to contribute debug support
| `$BPL_DEBUG_PORT` | What port the debug agent will listen on. Defaults to `8000`.
| `$BPL_DEBUG_SUSPEND` | Whether the JVM will suspend execution until a debugger has attached.  Defaults to `n`.


## Publishing the Port
When starting an application with debugging enabled, a port must be published.  To publish the port in Docker, use the following command:

```bash
$ docker run --publish <LOCAL_PORT>:<REMOTE_PORT> ...
```

The `REMOTE_PORT` should match the `port` configuration for the application (`8000` by default).  The `LOCAL_PORT` can be any open port on your computer, but typically matches the `REMOTE_PORT` where possible.

Once the port has been published, your IDE should connect to `localhost:<LOCAL_PORT>` for debugging.

![Eclipse Configuration](eclipse.png)

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
