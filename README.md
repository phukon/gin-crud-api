# Simple Invtry mgmt. API

This simple Go program implements a basic RESTful API for managing a book inventory. It uses the [Gin](https://github.com/gin-gonic/gin) framework for handling HTTP requests and responses.

## Features

- Get a list of all books
- Get details of a specific book by ID
- Create a new book
- Checkout a book (decrement quantity)
- Return a book (increment quantity)

## Getting Started

Make sure you have Go installed on your machine.

1.  Clone the repository:

```bash
 git clone <repository-url>
 cd <repository-directory>
```

2.  Initialize the Go module:

    `go mod init <module-name>`

    Replace `<module-name>` with your desired module name.

3.  Run the program:

    `go run main.go`

    The server will start on `localhost:8080`.

## API Endpoints

- **GET /books**: Get a list of all books
- **GET /books/:id**: Get details of a specific book by ID
- **POST /books**: Create a new book
- **PATCH /checkout**: Checkout a book (decrement quantity)
- **PATCH /return**: Return a book (increment quantity)

## Example Usage

### Get All Books

```bash
curl -X GET http://localhost:8080/books
```

### Get Book by ID

bashCopy code

`curl -X GET http://localhost:8080/books/1`

### Create a New Book

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": "5", "title": "New Book", "author": "Author Name", "quantity": 10}' http://localhost:8080/books
```

### Checkout a Book

```bash
curl -X PATCH -d 'id=1' http://localhost:8080/checkout
```

### Return a Book

```bash
curl -X PATCH -d 'id=1' http://localhost:8080/return
```

## Notes

- The program uses an in-memory slice (`books`) to store book data.
- The `gin` framework is used for handling HTTP requests and responses.
- The JSON field tags are used to customize the serialization and deserialization of struct fields.
