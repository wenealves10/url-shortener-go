# url-shortener-go

## Description

This is a URL shortener project developed in Go. It allows you to shorten long URLs to make them easier to share and remember.

## Dependencies

Before running the project, make sure you have MongoDB installed. You can run MongoDB using the following Docker command:

```bash
docker run -it -p 27017:27017 mongo
```

## Installation

To install the URL shortener, follow these steps:

1. Clone the repository to your local directory:

```bash
git clone https://github.com/wenealves10/url-shortener-go.git
```

2. Navigate to the project directory:

```bash
cd url-shortener-go
```

3. Run the following command to install the dependencies:

```bash
go mod download
```

4. Now, you can run the project.

## Usage

To run the URL shortener, follow these steps:

1. Navigate to the project directory:

```bash
cd url-shortener-go
```

2. Run the following command to start the server:

```bash
go run main.go
```

3. The server will start and listen on the specified port (by default, port 8080).

4. You can use a tool like cURL or an HTTP client to make requests to the server.

## Contributing

If you want to contribute to this project, follow these steps:

1. Fork this repository.

2. Create a branch for your new feature or bug fix:

```bash
git checkout -b my-contribution
```

3. Make the desired changes and add the modified files:

```bash
git add .
```

4. Commit your changes:

```bash
git commit -m "My contribution"
```

5. Push your changes to your fork:

```bash
git push origin my-contribution
```

6. Open a pull request in this repository, describing your changes.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
