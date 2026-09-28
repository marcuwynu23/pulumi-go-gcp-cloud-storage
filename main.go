package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		// Get configuration values
		bucketName := cfg.Get("storage:bucketName")
		if bucketName == "" {
			bucketName = "my-pulumi-bucket"
		}

		location := cfg.Get("storage:location")
		if location == "" {
			location = "US"
		}

		storageClass := cfg.Get("storage:storageClass")
		if storageClass == "" {
			storageClass = "STANDARD"
		}

		versioning := cfg.GetBool("storage:versioning")
		uniformBucketLevelAccess := cfg.GetBool("storage:uniformBucketLevelAccess")

		// Create a GCP resource (Storage Bucket)
		bucket, err := storage.NewBucket(ctx, "my-bucket", &storage.BucketArgs{
			Name:                       pulumi.String(bucketName),
			Location:                   pulumi.String(location),
			StorageClass:               pulumi.String(storageClass),
			Versioning:                 &storage.BucketVersioningArgs{Enabled: pulumi.Bool(versioning)},
			UniformBucketLevelAccess:   pulumi.Bool(uniformBucketLevelAccess),
		})
		if err != nil {
			return err
		}

		// Export the DNS name of the bucket
		ctx.Export("bucketName", bucket.Url)
		ctx.Export("bucketSelfLink", bucket.SelfLink)
		return nil
	})
}