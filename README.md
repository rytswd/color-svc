# color-svc

Simple HTTP server that gives you random color output.

## 🌅 Contents

- [Getting Started](#-getting-started)
- [Why and What is this for?](#-why-and-what-is-this-for)
- [Available Endpoints](#-available-endpoints)
- [How to Use](#-how-to-use)
- [Server Configurations](#-server-configurations)

## 🚀 Getting Started

The simplest way to get `color-svc` server started is the following:

```bash
$ docker run -it --rm \
    -p 8800:8800 \
    rytswd/color-svc:latest color-svc

# Output
2020/12/17 13:12:22 Server setup complete.
2020/12/17 13:12:22     Red    Enabled: true
2020/12/17 13:12:22     Green  Enabled: true
2020/12/17 13:12:22     Blue   Enabled: true
2020/12/17 13:12:22     Yellow Enabled: true
2020/12/17 13:12:22
2020/12/17 13:12:22     Total Available Colors: 15
2020/12/17 13:12:22 Server starting
...

# More log lines will appear as request is handled
```

With the server running, you can use the following:

```bash
$ curl http://localhost:8800/random/

# Output
  Generated Color
    "Green" - with HEX "#008000"
```

## ❓ Why and What is this for?

This is originally created for testing Service Mesh. The server is meant to be really simple, but provides different output depending on how you start up the server.

## 🧪 Available Endpoints

| Endpoint          | Description                                                                                       |
| ----------------- | ------------------------------------------------------------------------------------------------- |
| `/random/`        | Generate random color, based on available color set                                               |
| `/get/SOME_COLOR` | Get specific color - `SOME_COLOR` is case insensitive, and when not found, you will get "unknown" |
| `/red/`           | Generate random redish color. Can be disabled by environmental variable.                          |
| `/green/`         | Generate random greenish color. Can be disabled by environmental variable.                        |
| `/blue/`          | Generate random bluish color. Can be disabled by environmental variable.                          |
| `/yellow/`        | Generate random yellowish color. Can be disabled by environmental variable.                       |

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

The following environmental variables are used when starting the server.

| Env Name                     | Description                                                    | Default |
| ---------------------------- | -------------------------------------------------------------- | ------- |
| `ENABLE_DELAY`               | Enable delay for all requests                                  | false   |
| `DELAY_DURATION_MILLISECOND` | Delay duration in millisecond, only used when delay is enabled | 1000    |
| `DISABLE_RED`                | Disable all redish colors                                      | false   |
| `DISABLE_GREEN`              | Disable all greenish colors                                    | false   |
| `DISABLE_BLUE`               | Disable all bluish colors                                      | false   |
| `DISABLE_YELLOW`             | Disable all yellowish colors                                   | false   |
