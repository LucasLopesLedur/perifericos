# Peripherals API

## Features

- **GET /perifericos**: Retrieves all peripherals registered.
- **POST /perifericos**: Adds new peripherals.

### Project Structure

- `main.go`: Main server configuration.
- `README.md`: Project documentation.
- `go.mod`: Dependency management.

### Requirements

- Go (version 1.16 or higher)
- Gin (HTTP web framework for Golang)

## Usage

1. Clone the repository.
2. Run `go run main.go`.
3. Access `http://localhost:8080` to use the API.

## Example

### GET /perifericos

```json
[
    {
        "id": "789",
        "name": "Hyperx Quadcast",
        "price": 999.88
    }
]
```

## POST /perifericos
```
{
    "id": "123",
    "name": "Razer DeathAdder",
    "price": 299.99
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for more details.
