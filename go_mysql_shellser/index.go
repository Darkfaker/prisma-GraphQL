package main

import (
    "context"
    "fmt"
    prisma "test_go_prisma_ser/generated/prisma-client"
)

func main() {
    client := prisma.New(nil)
    ctx := context.TODO()

    // Create a new user
    name := "test_b"
    newUser, err := client.CreateUser(prisma.UserCreateInput{
        Name: name,
    }).Exec(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Created new user: %+v\n", newUser)

    users, err := client.Users(nil).Exec(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", users)
}
