/**
 *  Go programs express error state with error values.
 *  The error type is a built-in interface similar to fmt.Stringer:
 *  type error interface {
 *  Error() string
 *  }
 *  (As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
 *  Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
 */

 package main

 import (
   "fmt"
   "time"
 )

 type EngineError struct {
   When time.Time
   What string
 }

 func (e * EngineError) Error() string {
   return fmt.Sprintf("At %v error occured: %v", e.When, e.What)
 }

 func run() error {
   return &EngineError{
     time.Now(),
     "Fuel leakage in value 314",
   }
 }

 func main() {
   if err := run(); err != nil {
     fmt.Println(err)
   }
 }
