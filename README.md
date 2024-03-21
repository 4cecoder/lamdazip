# Lambda Function Packager

A Go-based tool to package AWS Lambda functions along with their dependencies into zip files.

## Features

- Automatically packages Lambda functions and their dependencies into zip files
- Supports multiple Lambda functions
- Uses `viper` for configuration management
- Provides a command-line interface using `cobra`

## Installation

### Option 1: Go Install

If you have Go installed on your system, you can use `go install` to install the Lambda Function Packager.

1. Run the following command: `go install github.com/4cecoder/lamdazip@latest`

2. The `lamdazip` executable will be installed in your `$GOPATH/bin` directory.

### Option 2: Manual Installation
1. Clone the repository: `git clone https://github.com/4cecoder/lamdazip.git`
 
2. Change into the project directory: `cd lamdazip`
 
3. Build the executable: `go build`

## Configuration

The Lambda Function Packager uses a configuration file to specify the Lambda function names, site-packages directory, and destination directory. By default, it looks for a configuration file named `.lambda-packager.yaml` in the user's home directory.

Here's an example configuration file:

```yaml
function_names:
- cold_rail
- new_customer
- updated_customer
site_packages_dir: venv/Lib/site-packages
dest_dir: terraform/dev/lambda
```


    function_names: An array of Lambda function names to be packaged.
    site_packages_dir: The directory path of the virtual environment's site-packages.
    dest_dir: The destination directory where the packaged Lambda functions will be moved.

You can create a configuration file with a different name or location and specify its path using the --config flag when running the tool.
Usage

To package Lambda functions, run the following command:

```bash
./lamdazip
```


If you want to use a different configuration file, specify its path using the --config flag:
```bash
./lamdazip --config path/to/config.yaml
```

The tool will package each Lambda function specified in the configuration file along with its dependencies into separate zip files. The zip files will be moved to the corresponding directories under the specified destination directory.
Project Structure

    main.go: The main Go file containing the Lambda Function Packager logic.
    .lambda-packager.yaml: The default configuration file.
    funx/: Directory containing the Lambda function files.
    terraform/dev/lambda/: The destination directory for the packaged Lambda functions.

Dependencies

The Lambda Function Packager uses the following dependencies:

    github.com/spf13/cobra: A library for creating powerful modern CLI applications.
    github.com/spf13/viper: A complete configuration solution for Go applications.

Make sure to have these dependencies installed before running the tool.
Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.
License

This project is licensed under the MIT License.
