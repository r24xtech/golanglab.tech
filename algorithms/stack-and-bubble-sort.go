package main

import (
  "fmt"
  "strconv"
)

type Stack struct {
  slice []int
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

func (s *Stack) Push() {
  x := get_int("Enter value: ")
  s.slice = append(s.slice, x)
}

func (s *Stack) Pop() {
  if len(s.slice) == 0 {
    fmt.Println("Empty stack!")
    return
  }
  s.slice = s.slice[0 : len(s.slice)-1]
}

func bubbleSort(numbers []int) {
  size := len(numbers)

  for i := 0; i < size; i++ {
    if !sweep(numbers, i) {
      return
    }
  }
}

func sweep(numbers []int, prevPasses int) bool{
  size := len(numbers)
  firstIndex := 0
  secondIndex := 1
  didSwap := false

  for secondIndex < (size - prevPasses) {
    firstNumber := numbers[firstIndex]
    secondNumber := numbers[secondIndex]
    if firstNumber > secondNumber {
      numbers[firstIndex] = secondNumber
      numbers[secondIndex] = firstNumber
      didSwap = true
    }
    firstIndex++
    secondIndex++
  }
  return didSwap
}

func main() {
  s := new(Stack)

  fmt.Println("|| Stack & BubbleSort ||")
  fmt.Println("1- Add; 2-Delete; 3-Print; 4-BubbleSort; 5-End\n\n")

  myloop:for {
    option := get_int("Choose option: ")
    switch option {
    case 1: s.Push()
    case 2: s.Pop()
    case 3: fmt.Println("bottom|", s.slice, "|top")
    case 4: bubbleSort(s.slice)
    case 5: break myloop
    default: fmt.Println("Invalid option.")
    }
  }
}
