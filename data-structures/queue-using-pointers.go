package main

import (
  "fmt"
  "strconv"
)

type Node struct {
  value int
  next *Node
}

type Queue struct {
  front *Node
  back *Node
  size int
}

func get_int(str string) int {
  for {
    var int_string string
    fmt.Print(str)
    fmt.Scanln(&int_string)
    int_number, error := strconv.Atoi(int_string)
    if error  == nil {
      return int_number
    }
    fmt.Println("Invalid input.")
  }
}

func (q *Queue) Push() {
  x := get_int("Enter value: ")
  n := new(Node)
  n.value = x
  if q.size == 0 {
    q.front = n
    q.back = q.front
  } else {
    q.back.next = n
    q.back = n
  }
  q.size++
}

func (q *Queue) Pop() {
  if q.size > 0 {
    q.front = q.front.next
    q.size--
  } else {
    fmt.Println("Can't pop...the queue is empty!")
  }
}

func (q *Queue) Print() {
  if q.size == 0 {
    fmt.Println("[...] The queue is empty!")
  } else {
    keep := new(Node)
    keep = q.front
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(q.front.value, " ")
      if q.front.next == nil {
        break
      }
      q.front = q.front.next
    }
    fmt.Print(" |back")
    q.front = keep
    fmt.Println("\n")
  }
}

func main() {
  q := new(Queue)

  fmt.Println("|| Welcome to the Queue ||")
  fmt.Println("1-Add; 2-Delete; 3-Print; 4-End\n\n")

  myloop:for {
    option := get_int("Choose option: ")
    switch option {
    case 1: q.Push()
    case 2: q.Pop()
    case 3: q.Print()
    case 4: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
