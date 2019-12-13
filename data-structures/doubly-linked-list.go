package main

import (
  "fmt"
  "strconv"
)

type Node struct {
  value int
  next *Node
  previous *Node
}

type List struct {
  size int
  head *Node
  tail *Node
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

func (l *List) InsertFront() {
  n := new(Node)
  x := get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.head.previous = n
    n.next = l.head
    l.head = n
  }
  l.size++
}

func (l *List) InsertBack() {
  n := new(Node)
  x := get_int("Enter value: ")
  n.value = x
  if l.size == 0 {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.next = n
    n.previous = l.tail
    l.tail = n
  }
  l.size++
}

func (l *List) Print() {
  if l.size == 0 {
    fmt.Println("[...] the list is empty!")
  } else {
    keep := new(Node)
    keep = l.head
    fmt.Println()
    fmt.Print("front| ")
    for {
      fmt.Print(l.head.value, " ")
      if l.head.next == nil {
        break
      }
      l.head = l.head.next
    }
    fmt.Print(" |back")
    fmt.Println("\n")
    l.head = keep
  }
}

func (l *List) RemoveBack() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.tail = l.tail.previous
    l.size--
  } else {
    l.tail = l.tail.previous
    l.tail.next = nil
    l.size--
  }
}

func (l *List) RemoveFront() {
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    l.head = l.head.next
    l.size--
  } else {
    l.head = l.head.next
    l.head.previous = nil
    l.size--
  }
}

func (l *List) RemoveAt() {
  x := get_int("Remove by index: ")
  if l.size == 0 {
    fmt.Println("Can't remove...the list is empty!")
  } else if l.size == 1 {
    if x == 0 {
      l.RemoveFront()
    } else {
      fmt.Println("Invalid index.")
    }
  } else if l.size == 2 {
    if x == 0 {
      l.RemoveFront()
    } else if x == 1 {
      l.RemoveBack()
    } else {
      fmt.Println("Invalid index.")
    }
  } else {
    if x == 0 {
      l.RemoveFront()
    } else if x == l.size-1 {
      l.RemoveBack()
    } else {
      if x > 0 && x < l.size-1 {
        keep := new(Node)
        keep = l.head
        for i := 0; i < x; i++ {
          l.head = l.head.next
        }
        l.head.previous.next = l.head.next
        l.head.next.previous = l.head.previous
        l.size--
        l.head = keep
      } else {
        fmt.Println("Invalid index.")
      }
    }
  }
}

func (l *List) InsertAt() {
  x := get_int("Insert by index: ")
  if l.size == 0 {
    l.InsertFront()
  } else if l.size == 1 {
    if x == 0 {
      l.InsertFront()
    } else if x == 1 {
      l.InsertBack()
    } else {
      fmt.Println("Invalid index.")
    }
  } else {
    if x == 0 {
      l.InsertFront()
    } else if x == l.size {
      l.InsertBack()
    } else if x > 0 && x < l.size {
      n := new(Node)
      keep := new(Node)
      keep = l.head
      for i := 0; i < x; i++ {
        l.head = l.head.next
      }
      z := get_int("Enter value: ")
      n.value = z
      n.next = l.head
      n.previous = l.head.previous
      l.head.previous = n
      l.head.previous.previous.next = n
      l.size++
      l.head = keep
    } else {
      fmt.Println("Invalid index.")
    }
  }
}

func main() {
  l := new(List)

  fmt.Println("|| Welcome ||")
  fmt.Println("[1] InsertBack; [2] InsertFront; [3] Print; [4]InsertByIndex;")
  fmt.Println("[5] RemoveBack; [6] RemoveFront; [7]RemoveByIndex; [8] End;\n\n")

  myloop:for {
    option := get_int("Choose option: ")
    switch option {
    case 1: l.InsertBack()
    case 2: l.InsertFront()
    case 3: l.Print()
    case 4: l.InsertAt()
    case 5: l.RemoveBack()
    case 6: l.RemoveFront()
    case 7: l.RemoveAt()
    case 8: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
