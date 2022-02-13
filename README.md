# Belgic

A webserver for [ambiorix](https://ambiorix.john-coene.com) applications.

## Install

```bash
go get github.com/devOpifex/belgic
```

or

``` bash
go install github.com/devOpifex/belgic@latest
```

or download one of the available binaries.

## Use

Belgic requires a very simple configuration file.

```json
{
 "applications": "/belgic/apps",
 "port": "8080"
}
```

To create it you can use the `config` command and pass it the _full
path_ to the _directory_ where you want the configuration file
to be created.

```bash
./belgic config -p "path/to/directory"
```

- `applications`: the directory containing the ambiorix applications
you want to serve.
- `port`: port on which the apps should be served.

Add the `BELGIC_CONFIG` environment variable to point to the configuration
file you just created.

Voilà, all set, just launch the server.

```bash
./belgic
```

## How it works

It's very similar to the way shiny-server works.
Point Belgic to a directory containing the applications you want
to serve (in the config file).

belgic then looks at all these apps and serves them individually.
e.g.:

```
/apps
  | /app1
  |   app.R
  | /app2
      app.R
```

In the above, point the config file to `/apps`, run belgic,
then the applications will be served at `mysite.com/app1`
and `mysite.com/app2`.
