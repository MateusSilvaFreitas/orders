﻿# Order System

## Overview

This project is a simple REST API developed in Go to manage orders, clients, products, and their relationships. It follows the MVC (Model-View-Controller) pattern to separate concerns, organizing the application into three main layers:

- **REST Layer**: Handles API endpoints using the Gin framework.
- **Service Layer**: Implements business logic.
- **Repository Layer**: Manages database interactions using `database/sql`.

## Features

- CRUD operations for orders, clients, and products
- RESTful API architecture
- Structured code with MVC pattern
- Uses Gin framework for routing
- Uses `database/sql` for database interactions

## Technologies Used

- **Go** (Golang)
- **Gin** (HTTP web framework)
- **database/sql** (Standard Go package for SQL database interaction)
- **MySQL** (Database)

## Installation & Setup

### Prerequisites

- Go installed on your system
- Running MYSQL on your system
- Execute the migration.sql

### Steps

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up environment variables:
   ```sh
   go run main.go
   ```

## API Endpoints

### Clients

- `GET /clients` - Get all clients
- `POST /clients` - Create a new client
- `GET /clients/:id` - Get a specific client
- `PUT /clients/:id` - Update client details
- `DELETE /clients/:id` - Delete a client

### Products

- `GET /products` - Get all products
- `POST /products` - Create a new product
- `GET /products/:id` - Get a specific product
- `PUT /products/:id` - Update product details
- `DELETE /products/:id` - Delete a product

### Orders

- `GET /orders` - Get all orders
- `POST /orders` - Create a new order
- `GET /orders/:id` - Get a specific order
- `PUT /orders/:id` - Update order details
- `DELETE /orders/:id` - Delete an order

## Contributing

This is a study project, so be free to sugest something!
