/**
 * When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value
 * on the right hand side.When the right hand side of the declaration is typed,
 * the new variable is of that same type:
 * However, when the right hand side contains an untyped numeric constant, the new variable
 * may be an int, float64, or complex128 depending on the precision of the constant:
 * i := 42           // int
 * f := 3.142        // float64
 * g := 0.867 + 0.5i // complex128
 */

package main
import "fmt"

func main() {
    var j int = 4
    i := j            // i is an int
    k := 3.142        // k would become float 64
    l := 0.867 + 0.5i // l would be complex128

    const f = "%T( %v )\n"  // %T --> type,  %v --> value
    fmt.Printf(f, i, i)
    fmt.Printf(f, k, k)
    fmt.Printf(f, l, l)
}
