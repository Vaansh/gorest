<h1> GoRest </h1>

A simple RESTful API with Go and Gin.

<b> Table of Contents </b>

- [Usage](#usage)
  - [Installation](#installation)
  - [Run](#run)
  - [Requests](#requests)
- [Output](#output)
- [License](#license)

## Usage

### Installation

```sh
git clone https://github.com/Vaansh/gorest.git
```

### Run

```sh
go run .
```

### Requests

**Examples**:

```sh
curl http://localhost:8080/items \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4", "name": "Item4", "price": 14.55}'
```

```sh
curl http://localhost:8080/items \
    --header "Content-Type: application/json" \
    --request "GET"
```

```sh
curl http://localhost:8080/items/4 \
    --header "Content-Type: application/json" \
    --request "GET"
```

```sh
curl http://localhost:8080/items/5 \
    --header "Content-Type: application/json" \
    --request "GET"
```

## Output

![Output](https://github.com/Vaansh/gorest/blob/main/docs/resource/server.png)

## License

Distributed under the MIT License. See `LICENSE` for more information.
