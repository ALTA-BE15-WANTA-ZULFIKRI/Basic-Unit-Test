package main 

import (
   "fmt"
   "testing"
)

func TestArrayUnique(t *testing.T)  {
   arrayA := []int{1, 2, 3, 4}
   arrayB := []int{1, 3, 5, 10, 16}
   expected := []int{2, 4}

   result := ArrayUnique(arrayA, arrayB)

   if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
      t.Errorf("Result: %v, Expected: %v", result, expected)
   }
}

func TestArrayUnique2(t *testing.T)  {
   arrayA := []int{10, 20, 30, 40}
   arrayB := []int{5, 10, 15, 59}
   expected := []int{20, 30, 40}

   result := ArrayUnique(arrayA, arrayB)

   if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
      t.Errorf("Result: %v, Expected: %v", result, expected)
   }
}

func TestArrayUnique3(t *testing.T)  {
   arrayA := []int{1, 3, 7}
   arrayB := []int{1,3,5}
   expected := []int{7}

   result := ArrayUnique(arrayA, arrayB)

   if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
      t.Errorf("Result: %v, Expected: %v", result, expected)
   }
}

func TestArrayUnique4(t *testing.T)  {
   arrayA := []int{3, 8}
   arrayB := []int{2, 8}
   expected := []int{3}

   result := ArrayUnique(arrayA, arrayB)

   if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
      t.Errorf("Result: %v, Expected: %v", result, expected)
   }
}

func TestArrayUnique5(t *testing.T)  {
   arrayA := []int{1, 2, 3}
   arrayB := []int{3, 2, 1}
   expected := []int{}

   result := ArrayUnique(arrayA, arrayB)

   if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
      t.Errorf("Result: %v, Expected: %v", result, expected)
   }
}