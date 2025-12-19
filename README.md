# Go Hiring Challenge

This repository contains a Go application for managing a product catalog, including products, variants, and categories.  
The application exposes REST APIs for listing products, retrieving product details, managing categories, and supports pagination, filtering, and standardized API responses.

## Project Structure

1. **cmd/**: Contains the application entry points.

    - `server/main.go`: Main application entry point, serves the REST API.
    - `seed/main.go`: Command to seed the database with initial product and category data.

2. **app/**: Contains application-level logic.
    - `catalog/`: HTTP handlers for catalog-related endpoints.
    - `categories/`: HTTP handlers for category-related endpoints.
    - `api/`: Shared API utilities (standardized response helpers).
    - `database/`: Database connection setup.

3. **models/**: Domain models and repository interfaces/implementations.
    - Product, Variant, and Category models
    - Repository interfaces and GORM-based implementations

4. **sql/**: Database migration scripts.

5. `.env`: Environment variables file for configuration.

## Features

- Product catalog with variants
- Product categories (Clothing, Shoes, Accessories)
- Product details endpoint with variant price inheritance
- Pagination with `offset` and `limit`
- Filtering by category and max price
- Standardized JSON API responses
- OpenAPI specification generated from code
- Unit tests for handlers and API utilities

## API Endpoints

### Catalog

- `GET /catalog`  
  List products with support for:
    - Pagination (`offset`, `limit`)
    - Filtering by category
    - Filtering by max price

- `GET /catalog/{code}`  
  Retrieve product details by product code, including:
    - Category information
    - Variants (variants without price inherit product price)

### Categories

- `GET /categories`  
  List all product categories.

- `POST /categories`  
  Create a new product category.

## API Response Format

All API responses follow a standardized format:

```json
{
  "status": true,
  "description": "Operation completed successfully",
  "payload": {},
  "errors": {}
}
```

## Setup Code Repository

1. Create a github/bitbucket/gitlab repository and push all this code as-is.
2. Create a new branch, and provide a pull-request against the main branch with your changes. Instructions to follow.

## Application Setup

- Ensure you have Go installed on your machine.
- Ensure you have Docker installed on your machine.
- Important makefile targets:
  - `make tidy`: will install all dependencies.
  - `make docker-up`: will start the required infrastructure services via docker containers.
  - `make seed`: ⚠️ Will destroy and re-create the database tables.
  - `make test`: Will run the tests.
  - `make run`: Will start the application.
  - `make docker-down`: Will stop the docker containers.

Follow up for the assignemnt here: [ASSIGNMENT.md](ASSIGNMENT.md)
