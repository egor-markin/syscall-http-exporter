# System Call HTTP Exporter

## Description

System Call HTTP Exporter is a lightweight Go application that allows you to expose system calls as HTTP endpoints. This project enables you to execute commands like `ls -l /` and get the results through a simple HTTP request. It supports multiple endpoints, making it ideal for monitoring, debugging, or automating system tasks remotely. This tool provides a flexible and secure way to interact with your system via RESTful APIs.

One particular use case for this project is to export certain system calls from the host machine to a Telegraf instance running in a Docker container. For example, you can use it to export Phusion Passenger's status data to Telegraf running in a Docker container on the same machine. The command that I was exporting was: `passenger-status -v --show=xml`.

## Features

- **Expose System Commands:** Configure and expose any system command as an HTTP endpoint.
- **Multiple Endpoints:** Easily define and manage multiple endpoints, each with its own command and response settings.
- **Customizable Endpoints:** Specify the content type for each endpoint, ensuring proper formatting of responses.
- **Configurable via JSON:** Simple configuration file (JSON) to manage endpoints and server settings.
- **Secure Execution:** Safeguards to prevent unauthorized access and ensure secure command execution.

## Getting Started

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/system-call-http-exporter.git
   cd system-call-http-exporter
   ```

2. **Build the application:**
   ```sh
   go build -o system-call-exporter main.go
   ```

3. **Create a configuration file (`config.json`):**
   ```json
   {
     "address": "localhost:8080",
     "endpoints": [
       {
         "command": "ls -l /",
         "endpoint": "/list-root",
         "content_type": "text/plain"
       },
       {
         "command": "date",
         "endpoint": "/current-date",
         "content_type": "text/plain"
       },
       {
         "command": "passenger-status -v --show=xml",
         "endpoint": "/passenger-status",
         "content_type": "application/xml"
       }
     ]
   }
   ```

4. **Run the application:**
   ```sh
   ./system-call-exporter -config=config.json
   ```

5. **Access the endpoints:**
    - List root directory: `http://localhost:8080/list-root`
    - Get current date: `http://localhost:8080/current-date`
    - Get Passenger status: `http://localhost:8080/passenger-status`

## Configuration

The application uses a JSON configuration file to define endpoints and server settings. Here is an example configuration:

```json
{
  "address": "localhost:8080",
  "endpoints": [
    {
      "command": "ls -l /",
      "endpoint": "/list-root",
      "content_type": "text/plain"
    },
    {
      "command": "date",
      "endpoint": "/current-date",
      "content_type": "text/plain"
    },
    {
      "command": "passenger-status -v --show=xml",
      "endpoint": "/passenger-status",
      "content_type": "application/xml"
    }
  ]
}
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.