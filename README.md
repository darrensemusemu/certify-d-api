# Certify-d API

This project supports `Certify-d`'s efforts as an online document certification platform. [What this project tries to solve.](https://notarycapetown.co.za/blog/why-you-need-to-certify-documents/)

## Architecture

A high-level [architectural overview](./docs/Architecture.md) of the system details how different person/actors, the systems itself, and other (external) systems interact.
## Project Structure

All services are for the API project  are within this single mono-repo. 
This project is structured is follows:

- [.prow](./.prow): k8s [Prow](https://prow.darrensemusemu.dev/) CI/CD jobs
- [api](./api): OpenAPI & gPRC related files
- [common](./common): shared pkg's usd by all micro-services
- [conf](./conf): project wide configuration files
- [docs](./docs): documentation related files
- [scripts](./scripts): scripts to run for builds, deploys, etc.
- [service.upload](./service.upload): micro-service, handling of all file storage operations
- [service.user](./service.user): micro-service, handling of all user operations
- [third_party](./third_party): external 3rd party tools/services
    - [kratos](./kratos): user identity & management system, see [Ory Kratos](https://www.ory.sh/docs/kratos/quickstart)
    - [oathkeeper](./kratos): identity & access auth proxy, [see ory.sh](https://www.ory.sh/docs/kratos/quickstart)

## Project versioning

Each project/micro-service is version independently. For each version, a git tag is created and used. The conversions for all tags are as follows: ***my_app/v0.0.0***

## See Also

[Swagger UI](https://certify-d.darrensemusemu.dev/api/swagger/index.html)


## TODO

- Certification service, handling of transforming files to stamped *'certified'* docs
- Notification service,  handling of all notification operations
- Order service,  handling of all order operations
- Payment service, handling of all payments operations
- <s>Bidding service: handling of outsourcing services to 3rd parties</s>