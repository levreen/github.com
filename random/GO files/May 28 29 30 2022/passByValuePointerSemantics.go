// sample program to show the basic concept of using a pointer
// to share data.

package main

func main() {

	// declare variable of type in with a value of 10.
	count := 10

	// display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\t\tAddr Of[", &count, "]")

	// pass the "address of" count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\t\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
// go:noinline
func increment(inc *int) { // *int newToMe

	// increment the "value of" count that the "pointer points to".
	*inc++ // newToMe

	println("inc:\tvalue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points to[", *inc, "]")
}
