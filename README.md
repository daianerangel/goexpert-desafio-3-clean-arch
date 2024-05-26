# Go Expert Desafio Clean Arch

## Instructions

### Run project

Execute `make docker-up`

Create table with `1.sql` file 

Execute `make run-app` 

#### Test with web server

* Open the api.http file
* Run the first request to create a order
* Run the second request to get all orders

#### Test with grpc

* In other terminal tab execute `make run-evans`
* Write `call CreateOrder` to create a new order
* Write `call ListOrder` to get all orders

#### Test with graphql

* In a web browser, go to `http://localhost:8080/`
* Write this code to create a new order 

```graphql 
mutation createOrder {
  createOrder(input: {id: "cccffff", Price: 12.2, Tax: 2.0}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

* Write this code to list a new order
```graphql
query queryOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```





