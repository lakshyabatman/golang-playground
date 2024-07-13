package main

import ("fmt"
        "math/rand"
  )


func repeatFunc[T any, K any](done <- chan K, fn func() T) <- chan T {
  stream := make(chan T)
  go func() {
    defer close(stream)
    for {
      select {
      case <- done:
        return
      case stream <- fn():

    }
    }
  }()
  return stream
}



func main()  {
  done := make(chan int)

  randNumFetcher := func() int {return rand.Intn(5000000)}
  defer close(done)
  stream := repeatFunc(done, randNumFetcher)
  for x := range stream {
    fmt.Println(x)
  }
  
}


