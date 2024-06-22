<h1 align="center">go-url-shortener</h1>

<div align="center">

This demonstration project implements a URL shortener API built with Go 1.22 using Chi. It allows users to shorten long URLs and retrieve the original URL from the shortened version.

<h2>Features</h2>

- Shorten URLs
- Redirect to original URL from shortened version

<h2>Getting Started</h2>
<h3>Prerequisites:</h3>

- Go installed
- Make(optional)

<h2>Running the Application:</h2>

1. Clone this repository.
2. Config the project by following this steps
3. Run the application using:

```bash
$ go run cmd/app/main.go

or

$ make run
```

<h1>API Endpoints: </h1>

- <b>Shorten a URL:</b>

  - <b>Method</b>: POST
  - <b>URL</b>: /shorten
  - <b>Request Body</b>:

  ```json
  {
    "url": "<URL to be shortened>"
  }
  ```

  - <b>Response</b>:
    - <b>On success</b>: JSON object containing the alias of url
    - <b>On failure</b>: Error message

- <b>Redirect to original URL:</b>

  - <b>Method</b>: GET
  - <b>URL</b>: /{alias} (Replace {alias} with the alias retrieved from /shorten)
  - <b>Response</b>:
    - Redirects to the original URL

<h2>Configuration:</h2>

1. Create `.env` file in project root folder and see the `.example.env` file.

<h2>Technologies and libs used: </h2>

- Go 1.22
- Chi v10
- PostgreSQL to persist urls
- Chi Render
- https://github.com/ilyakaznacheev/cleanenv
- https://github.com/go-playground/validator
