# Implement AWS Services using Pulumi Go

## Install pulumi CLI
`brew install pulumi/tap/pulumi`

## Prepare your stack

`pulumi stack init dev`

## Get needed go packages

```
go get "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
go get "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
```

## Provision your infrastructure

`pulumi up`
