# CloudQuery Hetzner Robot Source Plugin

[![test](https://github.com/serverness/cq-source-hrobot/actions/workflows/test.yaml/badge.svg)](https://github.com/serverness/cq-source-hrobot/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/serverness/cq-source-hrobot)](https://goreportcard.com/report/github.com/serverness/cq-source-hrobot)

A [Hetzner Robot](https://robot.hetzner.com/) source plugin for CloudQuery that loads data from Hetzner Robot WebService API to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Authentication

Credentials are used from the `HETZNER_TOKEN` environment variable with format `"{robot-generated-username}:{password}"`.

For more information on how to obtain an account, refer to the [official Hetzner Robot documentation](https://robot.hetzner.com/doc/webservice/en.html#preface).

Additionally, to use the Server ordering endpoints such as standard server products and the market market ones, follow [the official instructions](https://robot.hetzner.com/doc/webservice/en.html#activation).

## Configuration

The following source configuration file will sync all data from Scaleway to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "serverness"
  path: "serverness/hrobot"
  version: "${VERSION}"
  tables:
    - "*"
  skip_tables:
    - "hrobot_order_market_servers"
  destinations: 
    - "postgresql"
  backend: local
  spec:
    # plugin spec section
```

### Plugin Spec

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
