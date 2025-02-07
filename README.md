# DevEx CLI Tool

DevEx is a powerful command-line interface (CLI) tool designed to automate various DevOps tasks. It leverages AI to generate commands and code snippets, making it easier for developers and DevOps engineers to perform routine tasks efficiently. Additionally, it provides functionality to manage Pritunl VPN connections.

## Features

- **AI-Powered Command and Code Generation**: Interact with an AI agent to generate terminal commands and code snippets for various tasks.
- **Pritunl VPN Management**: Easily manage your Pritunl VPN connections with commands to connect, disconnect, add, and configure VPNs.

## Installation

To install the DevEx CLI tool, clone the repository and build the project:

```bash
git clone https://github.com/theswarnim/devex.git
cd devex
go build -o devex
```

## Usage

### AI-Powered Command and Code Generation

The `ai` command allows you to interact with an AI agent to generate commands or code snippets. Here are some examples:

1. **Generate a new Go module**:
    ```bash
    devex ai "generate a new go module"
    ```
    ![Generate Go Module](artifacts/generate_go_module.png)

2. **Generate an AWS CLI command using the prod profile to copy a file from local to an S3 bucket**:
    ```bash
    devex ai "generate a new aws cli command using prod profile to copy file from local to s3 bucket"
    ```
    ![Generate AWS CLI Command](artifacts/generate_aws_cli_command.png)

3. **Generate code to read a file in Go**:
    ```bash
    devex ai "generate code to read a file in go"
    ```
    ![Generate Go Code](artifacts/generate_go_code.png)

### Pritunl VPN Management

The `vpn` command allows you to manage your Pritunl VPN connections. Here are some examples:

1. **Connect to a VPN**:
    ```bash
    devex vpn connect <vpn_name>
    ```
    ![Connect to VPN](artifacts/connect_vpn.png)

2. **Disconnect from a VPN**:
    ```bash
    devex vpn disconnect <vpn_name>
    ```
    ![Disconnect from VPN](artifacts/disconnect_vpn.png)

3. **Add a new VPN configuration**:
    ```bash
    devex vpn add <vpn_name>
    ```
    ![Add VPN Configuration](artifacts/add_vpn.png)

4. **Set the Pritunl client CLI location**:
    ```bash
    devex vpn set-client-cli <path_to_pritunl_client_cli>
    ```
    ![Set Pritunl Client CLI Path](artifacts/set_client_cli_path.png)

## Examples

### AI-Powered Command and Code Generation

**Example 1**: Generate a Python script to read a file called `foo.com` in the `/tmp` directory.

```bash
devex ai "create a python code to read a file called foo.com in tmp directory"
```

**Response**:
```
- Task Description: Create a Python script that reads a file named "foo.com" in the tmp directory.
- Code Snippet:
```python
# read_file.py

file_path = "/tmp/foo.com"

try:
    with open(file_path, 'r') as file:
        content = file.read()
        print(content)
except FileNotFoundError:
    print("File not found.")
except Exception as e:
    print("An error occurred:", e)
```
- Explanation: This Python script attempts to open and read the file "foo.com" located in the tmp directory. If the file is found, its content is printed to the console. If the file is not found or an error occurs during reading, appropriate messages are displayed.

**Example 2**: Check the CPU utilization in macOS.

```bash
devex ai "how to check the cpu utilization in mac"
```

**Response**:
```
- Task Description: Check the CPU utilization in macOS.
- Generated Command(s):
  - macOS: "top -l 1 | grep 'CPU usage'"
- Explanation: This command uses the `top` command to display CPU utilization information for a single iteration and then filters the output to show the line containing "CPU usage". This will give you a snapshot of the current CPU utilization on your macOS system.
```

### Pritunl VPN Management

**Example 1**: Connect to a VPN.

```bash
devex vpn connect my_vpn
```

**Example 2**: Disconnect from a VPN.

```bash
devex vpn disconnect my_vpn
```

**Example 3**: Add a new VPN configuration.

```bash
devex vpn add my_vpn
```

**Example 4**: Set the Pritunl client CLI location.

```bash
devex vpn set-client-cli /path/to/pritunl-client
```

## Contributing

We welcome contributions from the community. If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

## Feedback

We are constantly working to improve the DevEx CLI tool. New features are coming soon! We would love to hear your feedback and suggestions. Please reach out to us on [GitHub](https://github.com/theswarnim/devex) or [contact@theswarnim.com](mailto:contact@theswarnim.com).
