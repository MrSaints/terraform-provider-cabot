# terraform-provider-cabot

_A work in progress._

[Terraform][terraform] [provider][terraform-provider] for [Arachnys'][arachnys] [Cabot][cabot]. Create, manage, and manipulate status checks, and alerts for services.

_Cabot is a self-hosted, easily-deployable monitoring, and alerts service - like a lightweight PagerDuty._


## Install

```go
go get -u github.com/mrsaints/terraform-provider-cabot
go build -o terraform-provider-cabot
```


## Usage

### Configuration

- `base_url` _(string)_
- `username` _(string)_
- `password` _(string)_

### Data Sources

#### cabot_plugin

### Resources

#### cabot\_check\_graphite

#### cabot\_check\_http

#### cabot\_check\_icmp

#### cabot\_check\_jenkins

#### cabot\_instance

#### cabot\_service


[terraform]: https://www.terraform.io/
[terraform-provider]: https://www.terraform.io/docs/plugins/provider.html
[arachnys]: https://www.arachnys.com/
[cabot]: https://github.com/arachnys/cabot