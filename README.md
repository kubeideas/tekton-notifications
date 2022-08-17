# Tekton Notifications

This document describes details about Notifications and how to build this.

## Supported types of notification

Notifications can be used in Tekton pipelines to send notifications messages. It was designed to use environment variables to get required information to send messages.

Supported technologies:

- SMTP
  - Environment variables SMTP_SERVER, SMTP_PORT, SMTP_USERNAME, SMTP_PASSWORD must be stored in an appropriate kubernetes secret and associated with notifications task.
- Slack
  - Use this link <https://api.slack.com/messaging/webhooks> to configure slack webhook and store URL in an appropriate Kubernetes secret and associated with notifications task.

## Build binary

``` bash
make build
```

## Build image and push

``` bash
make docker-build IMG=kubeideas/tekton-notifications:v0.1.0
make docker-push IMG=kubeideas/tekton-notifications:v0.1.0
```

## Notifications usage

To obtain all environment variables required for every type of technology apply the command below:

``` bash
notifications --usage
```
