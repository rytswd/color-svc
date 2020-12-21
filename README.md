# color-svc

Simple HTTP server that gives you random color output.

## 🌅 Contents

- [Why and What is this for?](#-why-and-what-is-this-for)
- [Getting Started](#-getting-started)
- [Available Endpoints](#-available-endpoints)
- [How to Use](#-how-to-use)
- [Server Configurations](#-server-configurations)

## ❓ Why and What is this for?

This is originally created for testing Service Mesh. The server is meant to be really simple, but provides different output depending on how you start up the server.

## 🚀 Getting Started

The simplest way to get `color-svc` server started is the following:

```bash
$ docker run -it --rm \
    -p 8800:8800 \
    rytswd/color-svc:latest color-svc

# Output
2020/12/17 13:12:22 Server setup complete.

        Per Request Delay: Disabled

        Red    : Enabled
        Green  : Enabled
        Blue   : Enabled
        Yellow : Enabled

        CORS Enabled    : Disabled
        Logging Enabled : Disabled

        Total Available Colors: 15

2020/12/17 13:12:22 Server starting
...

# More log lines will appear as request is handled
```

With the server running, you can use the following:

```bash
$ curl 'http://localhost:8800/random'

# Output
  Generated Color
    "Green" - with HEX "#008000"
```

## 🚥 Available Endpoints

| Endpoint         | Description                                                                                       |
| ---------------- | ------------------------------------------------------------------------------------------------- |
| `/random`        | Generate random color, based on available color set                                               |
| `/random/red`    | Generate random redish color, returns error if disabled                                           |
| `/random/green`  | Generate random greenish color, returns error if disabled                                         |
| `/random/blue`   | Generate random bluish color, returns error if disabled                                           |
| `/random/yellow` | Generate random yellowish color, returns error if disabled                                        |
| `/SOME_COLOR`    | Get specific color - `SOME_COLOR` is case insensitive, and when not found, you will get "unknown" |

## 🧪 Request Parameters

| Parameter   | Description                |
| ----------- | -------------------------- |
| `?fmt=json` | Return JSON representation |

## 🏁 How to use

<details>

<summary>With Docker</summary>

The simplest way to get started is the following:

```bash
$ docker run -it --rm \
    -p 8800:8800 \
    rytswd/color-svc:latest color-svc
```

You can adjust the behaviour with providing environmental variables. For example, the below command will only provide bluish colors.

```bash
$ docker run -it --rm \
    -p 8800:8800 \
    -e DISABLE_RED=true \
    -e DISABLE_GREEN=true \
    -e DISABLE_YELLOW=true \
    rytswd/color-svc:latest color-svc
```

Also, the container is pushed to GitHub Container Registry at `ghcr.io/rytswd/color-svc`.

</details>

<details>

<summary>With Kubernetes</summary>

This repository contains example Kubernetes Service + Deployment YAMLs in [/k8s](k8s) directory.

You can use them as is, or adjust it as you like. There are some environmental variables that can adjust the server behaviours.

Example:

```bash
$ kubectl apply \
    -f https://raw.githubusercontent.com/rytswd/color-svc/main/k8s/account.yaml \
    -f https://raw.githubusercontent.com/rytswd/color-svc/main/k8s/color-svc-default.yaml
```

[k8s]: https://github.com/rytswd/color-svc/tree/main/k8s/

</details>

<details>

<summary>From code</summary>

Simply run with `go run cmd/server/main.go`. You can provide environmental variables to adjust some behaviours.

</details>

## ⚙️ Server Configurations

The following environmental variables are read when the server starts up.

| Env Name                     | Description                                                    | Default |
| ---------------------------- | -------------------------------------------------------------- | ------- |
| `ENABLE_DELAY`               | Enable delay for all requests                                  | false   |
| `DELAY_DURATION_MILLISECOND` | Delay duration in millisecond, only used when delay is enabled | 1000    |
| `ENABLE_CORS`                | Enable CORS                                                    | false   |
| `DISABLE_LOGGING`            | Disable per request logging                                    | false   |
| `DISABLE_RED`                | Disable all redish colors                                      | false   |
| `DISABLE_GREEN`              | Disable all greenish colors                                    | false   |
| `DISABLE_BLUE`               | Disable all bluish colors                                      | false   |
| `DISABLE_YELLOW`             | Disable all yellowish colors                                   | false   |
