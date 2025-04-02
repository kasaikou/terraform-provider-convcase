# terraform-provider-convcase

`terraform-provider-convcase` is a custom Terraform provider that allows you to convert strings into various case styles, such as camel case, snake case, kebab case, and more. This provider is useful for managing string transformations in your Terraform configurations.

## Features

- Convert strings to different case styles:
    - Camel Case
    - Pascal Case
    - Snake Case
    - Constant Case
    - Kebab Case
    - Train Case
    - Path Style
    - Dot Style
- Supports custom Terraform functions and data sources.

## Installation

To use this provider, add it to your Terraform configuration:

```hcl
terraform {
    required_providers {
        convcase = {
            source  = "kasaikou/convcase"
            version = "0.1.0"
        }
    }
}
```

## Usage

### Example: Convert a String to Different Case Styles

```hcl
data "convcase" "example" {
    input = "example_string"
}

output "camel_case" {
    value = data.convcase.example.camel
}

output "snake_case" {
    value = data.convcase.example.snake
}

output "kebab_case" {
    value = data.convcase.example.kebab
}
```

### Example: Using Custom Functions

```hcl
output "camel_case" {
    value = conv_camel("example_string")
}

output "snake_case" {
    value = conv_snake("example-string")
}
```

## Development

### Prerequisites

- Go 1.23.4 or later
- Terraform Plugin Framework

### Build

To build the provider locally:

```bash
go build -o terraform-provider-convcase
```

### Test

Run the tests:

```bash
go test ./...
```

### Release

This project uses [GoReleaser](https://goreleaser.com) for releases. To create a release:

```bash
goreleaser release --clean
```

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Acknowledgments

This provider is built using the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework).
