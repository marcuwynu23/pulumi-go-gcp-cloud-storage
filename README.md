# Pulumi GCP Go: Configurable Storage Bucket Template

This template provisions a Google Cloud Storage bucket using Pulumi and Go. It demonstrates how to:
- Use the Pulumi GCP provider in a Go program
- Create a simple GCP resource (a Storage Bucket) with configurable settings
- Export resource outputs for use in your stacks

It's a great starting point for learning Pulumi with Go on GCP or bootstrapping a project that needs object storage.

## Providers

- Google Cloud Platform via the Pulumi GCP SDK for Go (`github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp`)

## Resources

- **Storage Bucket** (`gcp.storage.Bucket`)
  - Logical name: `my-bucket`
  - Configurable name, location, storage class, versioning, and uniform bucket level access

## Outputs

- **bucketName**: The URL of the newly created bucket (e.g., `https://storage.googleapis.com/my-bucket`)
- **bucketSelfLink**: The self link of the bucket resource

## When to Use This Template

- You want a minimal Pulumi program in Go targeting GCP
- You need a simple object storage bucket for assets or data
- You're exploring Pulumi's Go SDK for cloud provisioning

## Prerequisites

- Go 1.21 or later installed
- A Google Cloud account with billing enabled
- GCP credentials configured for Pulumi (for example, via `gcloud auth application-default login`)

## Usage

1. Scaffold a new project from this template:
   ```bash
   pulumi new gcp-go
   ```
2. When prompted, fill in:
   - **Project name**: your desired project identifier
   - **Description**: a short description of your stack
   - **gcp:project**: your target GCP project ID
3. Change into your project directory:
   ```bash
   cd <your-project-name>
   ```
4. Configure your bucket settings:
   ```bash
   cp Pulumi.dev.yaml.example Pulumi.dev.yaml
   # Edit Pulumi.dev.yaml with your settings
   pulumi config set gcp:project YOUR_PROJECT_ID
   pulumi config set storage:bucketName my-unique-bucket-name
   pulumi config set storage:location US
   pulumi config set storage:storageClass STANDARD
   pulumi config set storage:versioning true
   pulumi config set storage:uniformBucketLevelAccess true
   ```
5. Preview and deploy your stack:
   ```bash
   pulumi preview
   pulumi up
   ```

## Project Layout

```
├── Pulumi.yaml                  Pulumi project definition and template settings
├── Pulumi.dev.yaml.example      Template for local dev configuration
├── go.mod                       Go module declaration and dependencies
├── main.go                      Pulumi program defining the Storage Bucket
├── .gitignore                   Git ignore rules
└── LICENSE                      MIT License
```

## Configuration

**Important**: `Pulumi.dev.yaml` is gitignored to protect project credentials. Start with the template:

```bash
cp Pulumi.dev.yaml.example Pulumi.dev.yaml
```

Then fill in your `gcp:project` value and set it via:

```bash
pulumi config set gcp:project YOUR_PROJECT_ID
```

The following Pulumi configuration values are available:

| Name | Description | Default |
|------|-------------|---------|
| `gcp:project` | The Google Cloud project to deploy into | _required_ |
| `gcp:region` | The GCP region for the bucket | `US` |
| `storage:bucketName` | The name of the storage bucket | `my-pulumi-bucket` |
| `storage:location` | The location of the bucket | `US` |
| `storage:storageClass` | The storage class of the bucket | `STANDARD` |
| `storage:versioning` | Whether to enable versioning | `true` |
| `storage:uniformBucketLevelAccess` | Whether to enable uniform bucket level access | `true` |

## Next Steps

- Add more GCP resources (e.g., Compute Engine, Pub/Sub, Cloud Functions)
- Parameterize bucket settings such as lifecycle rules, CORS, and IAM bindings
- Integrate IAM bindings for fine-grained permission management
- Connect this bucket to other services or CI/CD pipelines

## Getting Help

- Pulumi Documentation: https://www.pulumi.com/docs/
- GCP Provider Reference: https://www.pulumi.com/registry/packages/gcp/
- Community Slack: https://slack.pulumi.com/
- GitHub Issues: https://github.com/pulumi/pulumi/issues