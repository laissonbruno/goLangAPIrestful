# Simple RESTful API with Go (Golang)

This is a simple example of a RESTful API built with Go (Golang). It demonstrates basic CRUD (Create, Read, Update, Delete) operations for managing items.

## Requirements

- Go (Golang) installed on your machine
- The `gorilla/mux` package for routing (`go get -u github.com/gorilla/mux`)

## Usage

1. Clone this repository:

    ```bash
    git clone https://github.com/laissonbruno/goLangAPIrestful.git
    ```

2. Navigate to the project directory:

    ```bash
    cd goLangAPIrestful
    ```

3. Run the Go program:

    ```bash
    go run main.go
    ```

4. Once the server is running, you can interact with the API endpoints using tools like cURL or Postman:

    - `GET /items`: Retrieve all items.
    - `GET /items/{id}`: Retrieve a specific item by ID.
    - `POST /items`: Create a new item.

5. To stop the server, press `Ctrl + C` in the terminal.

## API Endpoints

- `GET /items`: Retrieve all items.
- `GET /items/{id}`: Retrieve a specific item by ID.
- `POST /items`: Create a new item.

## Example

1. Get all items:

    ```bash
    curl http://localhost:8000/items
    ```

2. Get an item by ID (replace `{id}` with the actual ID):

    ```bash
    curl http://localhost:8000/items/{id}
    ```

3. Create a new item:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"id":"3","name":"New Item","description":"Description of New Item"}' http://localhost:8000/items
    ```

