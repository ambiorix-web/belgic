# Eburon

A webserver for [ambiorix](https://ambiorix.john-coene.com) applications.

## Install

```bash
go get github.com/devOpifex/eburon
```

or

``` bash
go install github.com/devOpifex/eburon@latest
```

or download one of the available binaries.


## Use

Eburon requires a very simple configuration file.

```json
{
 "applications": "/eburon/apps",
 "port": "8080"
}
```

To create it you can use the `config` command and pass it the _full
path_ to the _directory_ where you want the configuration file
to be created.

```bash
./eburon config -p "path/to/directory"
```

- `applications`: the directory containing the ambiorix applications
you want to serve.
- `port`: port on which the apps should be served.

Voilà, all set, just launch the server.

```bash
./eburon
```

## How it works

It's very similar to the way shiny-server works.
Point eburon to a directory containing the applications you want
to serve (in the config file).

Eburon then looks at all these apps and serves them individually.
e.g.:

```
/apps
  | /app1
  |   app.R
  | /app2
  |   app.R
```

In the above, point the config file to `/apps`, run eburon,
then the applications will be served at `mysite.com/app1`
and `mysite.com/app2`.
