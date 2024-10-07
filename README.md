# DNS Check

DNS Check is a command-line tool that helps you verify if a list of websites can be accessed through specific DNS server. The tool reads a list of DNS servers from a file (dns.txt), takes one or more website addresses as command-line arguments, and checks if each website can be opened with the provided DNS servers.

## Features

- **Custom DNS Server List**: Specify DNS servers in `dns.txt` with the format `name=ip`.
- **Multi-site Checking**: Pass one or more website addresses as arguments to check their availability.
- **DNS Resolution Test**: For each DNS server, the program attempts to resolve and access the websites provided.

## Usage

1. **Prepare the DNS server list**:

    Create or modify the `dns.txt` file in the project directory. The file should contain the DNS server name and IP in the following format:

    ```txt
    cloudflare=1.1.1.1
    google=8.8.8.8
    ```

2. **Run the program**:

    Pass one or more website addresses as arguments (without `http://` or `https://`). For example, to check `example.com` and `github.com`:

    ```bash
    ./dns-check example1.com example2.com
    ```

3. The program will attempt to resolve each website using each DNS server from the `dns.txt` file and display whether the website can be accessed.
