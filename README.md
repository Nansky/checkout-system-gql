# Simple Checkout tasks
#### _Kuncie Take Home Test_

[![GraphQL|GQL](https://github.com/graph-gophers/graphql-go/raw/master/docs/img/logo.png)](https://github.com/graph-gophers/graphql-go)

This Repository is purposed as a Kuncie's technical test as Backend Engineer. This service using graphql as protocol from Fromtend (Assuming graphiQL playground as a Frontend). You can test this service using tool like Postman. 

## Dependency Required
- Golang built-in `net/http` package
- `Redis`
- `go-graphql` as package for graphql

## Features
There are 3 features that have been made :
- [Mutation] Add Items to Cart
- [Query] Retrieve the item added from cart
- [Mutation] Checkout all Items from cart

## Manual Installation
- clone the source code from this repo
- Install Redis if you dont set it up, Run it
- Run The following commands

 ```sh
./checkout-system-gql
or 
go run checkout-system-gql
 ``` 
And then you will see this in your CLI
`2021/11/22 12:59:10 Server Running at port 6969`
- And then you can try to open GraphiQL playground to test the service from your browser
  `http://localhost:6969/add_checkout`

## Schema Design
I'm using go-graphql , so the schema is defined inside the code, you can check the schemas under `schemas/` subfolder

## Service Specification
Use this query to test in Playground
1. Add Items to Cart
```
mutation {
  addItems(buyerId:2, itemType:1, quantity:1){
    items
    price
    total
  }
}
```
Where the params are :
| Plugin | README |
| ------ | ------ |
| buyerId | The ID of the buyer (Unique) for each buyer |
| itemType | The type of items you want to Buy (1. Macbook, 2. Google Home, 3. Alexa Speaker, 4. Raspberry Pi)|
| quantity | Total items you want to buy |

2. Query items in cart
```
query {
	buyer(buyerId: <Your Buyer ID in integer> ){
    name
    sku
    price
    quantity
  }  
}
```

example query result 
```
{
  "data": {
    "buyer": [
      {
        "name": "MacBook Pro",
        "price": "5399.99",
        "quantity": "1",
        "sku": "43N23P"
      },
      {
        "name": "Google Home",
        "price": "49.99",
        "quantity": "3",
        "sku": "120P90"
      },
      {
        "name": "Alexa Speaker",
        "price": "109.50",
        "quantity": "2",
        "sku": "A304SD"
      }
    ]
  }
}
```

3. Checkout
```
mutation {
  checkout(buyerId:2) {
    desc
    total
  }
}
```
based on this below cases, 
![testcases](https://i.postimg.cc/3WQwXZQH/Screen-Shot-2021-11-22-at-16-41-54.png)

if we apply it on this mutation, the result will be like this 
```
{
  "data": {
    "checkout": [
      {
        "desc": "Scanned Items: MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B",
        "total": "$26999.95"
      },
      {
        "desc": "Scanned Items: Google Home, Google Home, Google Home, Google Home, Google Home, Google Home, Google Home, Google Home, Google Home, Google Home",
        "total": "$349.93"
      },
      {
        "desc": "Scanned Items: Raspberry Pi B, Raspberry Pi B",
        "total": "$60.00"
      },
      {
        "desc": "Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker",
        "total": "$985.50"
      }
    ]
  }
}
```

