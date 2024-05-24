# Tech Challenge 1

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Architecture](#architecture)
- [Docker build and run](#docker-build-and-run)
- [How to use](#how-to-use)
- [Section 1 - Restaurant owner](#section-1-restaurant-owner)
  - [1 Create product](#1-create-product)
- [Section 2 Customer order](#section-2-customer-order)
  - [1 User identification](#1-user-identification)
  - [2 List all the categories](#2-list-all-the-categories)
  - [3 List products by the chosen category](#3-list-products-by-the-chosen-category)
  - [4 Pay the products amount](#4-pay-the-products-amount)
  - [5 Create an order](#5-create-an-order)
  - [6 List orders to follow](#6-list-orders-to-follow)
  - [7 List orders to prepare](#7-list-orders-to-prepare)
  - [8 Update order to preparing](#8-update-order-to-preparing)
  - [9 Update order to done](#9-update-order-to-done)
  - [10 Update order to delivered](#10-update-order-to-delivered)
  - [11 Update order to not delivered](#11-update-order-to-not-delivered)
- [Documentation](#documentation)
  - [Event storming](#event-storming)
  - [Postman collection](#postman-collection)
  - [Swagger](#swagger)
  - [Redoc](#redoc)
- [Opportunities](#opportunities)

## Description

The Tech Challenge 1 aims to do a solution for a Fast Food restaurant. With this software, the rastaurant can do a all the process of for a place
that makes all steps of a fast food dish order, as:

- Products creation/manipulation by the restaurant owner
- Customer identification
- Order creation with given products
- Payment process
- Order tracking by the chef
- Order tracking by the waiter and the customer

This projects only fits the Backend side, which means that customer needs to **choose** the products or combo by a interface previously. This Backend will only receive the *entire order with all chosen products or combos*. This Backend will not do a *step by step product selecion*.

All the Endpoints can be called by accessing `http://localhost:3210/api` API url.

To build and run this project. Follow the Docker section

## Architecture

To design this application was chosen `DDD (Domain Drive Design)` architecture to follow the principle of **protecting the model**.
The folder project was created to follow DDD principle:

- cmd/api: The application entrypoint
- docs: OpenAPI documentation
- internal: All the application data like services, repositories, entities, etc
- - internal/core: The Domain of this project
- - - internal/core/domain: All the domain structures representing the classes and aggregates
- - - internal/core/ports: All the domain interfaces. With this, we can protect the domain from external classes
- - - internal/core/services: All the business logic (Services or Use Cases).
- - internal/adapters: The folder to represent all the adapters for the DDD principle
- - - internal/adapters/driven: The driven adapter to comunicate with external dependencies, like databases and REST apis
- - - - internal/adapters/driven/entities: All the structures to represent Database Entities
- - - - internal/adapters/driven/external: REST Api consumers
- - - - internal/adapters/driven/repositores: The clue between Core Ports and External Dependencies. Here we have the implementations of the Core Ports interfaces
- - - internal/adapters/driver/handler: This is the driver adapter of the DDD principle. Here resides the REST handlers to be consumed by the client


## Docker build and run

This project was built using Docker and Docker Compose. So, to build and run the API, we need to run in the root of the project:

```
$ docker compose build
```

After the image build finish, run:

```
$ docker compose up -d
```

The command above may take a while.

After all the containers shows these below status:

```
 ✔ Container fastfood-database  Started
 ✔ Container fastfood-app       Started 
```

we can access `http://localhost:3210/api` endpoints.


## How to use

To use all the endpoints in this API, we can follow these sequence to simulate a customer making an order in a restaurant.
We can separate in three moments.

- Restaurant products manipulation. This is used by the `restaurant owner` to create all the product portfolio with its images and prices
- Customer self service. This is used by the `customer` to choose the products, pay for it and create an order 
- Order preparing and deliver. This is used by the `chef` and `waiter` to check the order status

We will divide in 2 sections: **Restaurant owner** and **Customer order**


## Section 1 Restaurant owner

This section will be used by the restaurant owner to manage the restaurant products

### 1 Create product
***(Owner view)***

- Cal the POST `http://localhost:3210/api/products` to create a Product
- Cal the PUT `http://localhost:3210/api/products/{id}` to update a Product
- Cal the DELETE `http://localhost:3210/api/products/{id}` to delete a Product

With those endpoints we can follow to *Section 2* to start the ***Order flow***


## Section 2 Customer order

This section will use all the Endpoints to make a entire order flow.

### 1 User identification
***(Customer view)***

- Call the GET `http://localhost:3210/api/customers/{cpf}` to login and get this `[Customer ID]`
or
- Cal the POST `http://localhost:3210/api/customers` to create a Customer and retrieve the `[Customer ID]`

> [!IMPORTANT]
> These endpoints have a CPF validation. So be aware that it is needed to pass a correct CPF number.

> [!NOTE]  
> The CPF does not need to be formatted.

We can use this site [CPF generator](https://www.4devs.com.br/gerador_de_cpf) to easly generate a new CPF whenever we need.

### 2 List all the categories
***(Customer view)***

- Call the GET `http://localhost:3210/api/products/categories` to get a string array with all created categories

### 3 List products by the chosen category
***(Customer view)***

- Call the GET `http://localhost:3210/api/products/categories/{category}` to get all products by a category

With this endpoints we can simulate a screen producst selection by chosing all products IDs we want to deal and create a Order

### 4 Pay the products amount
***(Customer view)***

- Call the GET `http://localhost:3210/api/payments/types` to show to customer which payment type to choose
- Call the POST `http://localhost:3210/api/payments` to pay for the amount and receive the `[Payment ID]`

### 5 Create an order
***(Customer view)***

- Call the POST `http://localhost:3210/api/orders` with:
- - All the `[Products IDs]` chosen [*required]
- - The `[Payment ID]` [*required*]
- - The `[Customer ID]` [*optional*]
- - Total price for the all products sum

### 6 List orders to follow
***(Customer and Waiter)***

- Call the GET `http://localhost:3210/api/orders/follow` to show a list of Orders to be followed by Customer and Waiter

The order can be followed by its ID:
- Call the GET `http://localhost:3210/api/orders/{id}` to show a an Orders to be followed by Customer and Waiter

### 7 List orders to prepare
***(Chef view)***

- Call the GET `http://localhost:3210/api/orders/to-prepare` to list the Orders with its [Order ID]

### 8 Update order to preparing
***(Chef view)***

- Call the PUT `http://localhost:3210/api/orders/{id}/preparing` to set Preparing status

### 9 Update order to done
***(Chef view)***

- Call the PUT `http://localhost:3210/api/orders/{id}/done` to set Done status

### 10 Update order to delivered
***(Waiter view)***

- Call the PUT `http://localhost:3210/api/orders/{id}/delivered` to set Delivered status to indicate that customer receive the meal. 
This is used to 'finish' the order and can be used to track some convertion rate

### 11 Update order to not delivered
***(Waiter view)***

- Call the PUT `http://localhost:3210/api/orders/{id}/not-delivered` to set Not Delivered status to indicate that customer doesn not receive the meal.
This is used to 'finish' the order and can be used to track some convertion rate


## Documentation

This project uses Swagger to show an site with all Endpoints used by this project to make an order in a Fast Food place. 
To create/update all Endpoints documentation just run `swag init -g cmd/api/main.go`. By doing this, we can see the documentation in
two different ways:

### Event Storming

This project was guided by the DDD Event Storming. This document was made in Miro. The file `event_storming.pdf` is in the root of this project.
We can algo see the Miro project by accessing the link:[Event Storming](https://miro.com/welcomeonboard/NU9HT081SExIOEkyYUNZOGxVNGQyQ2h3NjFxRjNTUlB2MUZVV2xaZmF4Q0RYV1lIajlvOERFbGo2Mmx6clBvVXwzNDU4NzY0NTg4MDA1OTg4NTQwfDI=?share_link_id=625787819956)

### Postman collection

In the root of this project we can find the file `postman_collection.json`. With this we can easly test all the Endpoints

### Swagger

http://localhost:3210/swagger/index.html

### Redoc

http://localhost:3211/docs


## Opportunities

Even though this project was made by following some Design Patterns like `Use Case` and `Repository` it does not separate the Data Source from the Repository. In the future it will be good to use `Data Source` Pattern to separate **Local** and **Remote** from the Repository to make a better separation of concern principle.

It will also a good opportunity to increase the `Unit Test` coverage.