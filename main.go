package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "hh-pulumi-go", nil)
		if err != nil {
			return err
		}

		// Create Secuirty Group
		group, err := ec2.NewSecurityGroup(ctx, "web-sg", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Enable http Access"),
			Ingress: ec2.SecurityGroupIngressArray{
				ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(80),
					ToPort:     pulumi.Int(80),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
		})
		if err != nil {
			return err
		}
		// Create EC2 insatnce
		server, err := ec2.NewInstance(ctx, "test", &ec2.InstanceArgs{
			Ami:                 pulumi.String("ami-0827b6c5b977c020e"),
			InstanceType:        pulumi.String("t2.micro"),
			VpcSecurityGroupIds: pulumi.StringArray{group.ID()},
		})
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("public IP", server.PublicIp)
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
