![](belgic.png)

A reverse proxy and load balancer for 
[ambiorix](https://ambiorix.dev) applications
(and [shiny](https://shiny.rstudio.com/)).

:warning: This implements a round robin, requests are redirected to
whatever backend is next on the queue. This means you should not
store any session-related data in the environment, use databases,
cookies, parameters, etc. (as one should anyway).
This will NOT work with shiny applications.

## Install

```bash
go get github.com/devOpifex/belgic
```

or

``` bash
go install github.com/devOpifex/belgic@latest
```

or download one of the [available binaries](https://github.com/devOpifex/belgic/tree/master/bin).

## Use

Belgic requires a very simple configuration file.

```json
{
 "path": "/belgic",
 "port": "8080",
 "backends": "max",
 "attempts": 3
}
```

To create it you can use the `config` command and pass it the _full
path_ to the _directory_ where you want the configuration file
to be created.

```bash
./belgic config -p "path/to/directory"
```

- `path`: the path containing the ambiorix application
you want to serve. It assumes the application is in an `app.R` file.
- `port`: port on which the apps should be served.
- `backends`: number of background applications to run in the background.
Defaults to the maximum number of cores available on the machine.
- `attempts`: number of times to try and revive a backend if it dies.

Add the `BELGIC_CONFIG` environment variable to point to the configuration
file you just created.

Voilà, all set, just launch the server.

```bash
belgic start
# ./belgic start
```

## Backends

The server will launch multiple applications in the background.
The number of applications running in the background is determined
by the `backends` variable defined in the configuration file.
Either set this option to the number of applications you want to 
run in the background __as a string__, 
e.g.: set it to `"4"` not `4`.
If set to `"max"` if will run one application for each core
available on the machine.
