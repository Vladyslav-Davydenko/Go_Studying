package main

import "fmt"

// zeroval receives a COPY of the integer value.
// Any changes made inside this function do NOT affect the original variable.
func zeroval(ival int) {
	ival = 0 // This only changes the local copy
}

// zeroptr receives a POINTER (memory address) to an integer.
// The * in "*int" means "pointer to int" - this is a TYPE declaration.
func zeroptr(iprt *int) {
	// The * here is the DEREFERENCE operator.
	// It means: "go to the memory address stored in iprt and change the value there"
	// We're NOT changing the pointer itself, we're changing what it POINTS TO.
	*iprt = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Pass by value - zeroval gets a copy, original i is unchanged
	zeroval(i)
	fmt.Println("zeroval:", i) // Still 1

	// Pass by reference - zeroptr gets the ADDRESS of i
	// &i means "give me the memory address of i"
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // Now 0

	// Print the memory address of i
	fmt.Println("address of i:", &i)

	// ============================================
	// POINTER OPERATIONS EXPLAINED
	// ============================================

	fmt.Println("\n--- Pointer Operations ---")

	value := 42
	ptr := &value // ptr now holds the ADDRESS of value

	fmt.Println("value:", value)   // 42 - the actual value
	fmt.Println("&value:", &value) // 0x... - address of value
	fmt.Println("ptr:", ptr)       // 0x... - same address (ptr holds it)
	fmt.Println("*ptr:", *ptr)     // 42 - value AT the address ptr holds

	// ============================================
	// KEY INSIGHT: & and * are INVERSE operations
	// ============================================
	//
	//   &  = "address of" (value → address)
	//   *  = "value at"   (address → value)
	//
	//   &value  gives you the ADDRESS of value
	//   *ptr    gives you the VALUE at the address
	//
	//   They cancel each other out:
	//   *(&value) == value  ✓
	//   &(*ptr) == ptr      ✓

	fmt.Println("\n--- Inverse Operations ---")
	fmt.Println("*(&value) == value:", *(&value) == value) // true
	fmt.Println("&(*ptr) == ptr:", &(*ptr) == ptr)         // true

	// ============================================
	// CHANGING POINTER vs CHANGING VALUE
	// ============================================

	fmt.Println("\n--- Changing Pointer vs Value ---")

	a := 10
	b := 20
	vptr := &a // vptr points to a

	fmt.Println("a:", a, "| b:", b, "| *vptr:", *vptr)

	// Change the VALUE at the address (modifies 'a')
	*vptr = 100
	fmt.Println("After *vptr = 100:")
	fmt.Println("a:", a, "| b:", b, "| *vptr:", *vptr) // a is now 100

	// Change the POINTER itself (now points to 'b' instead of 'a')
	vptr = &b
	fmt.Println("After vptr = &b:")
	fmt.Println("a:", a, "| b:", b, "| *vptr:", *vptr) // *vptr now shows b's value

	// Now changing *vptr will modify 'b', not 'a'
	*vptr = 200
	fmt.Println("After *vptr = 200:")
	fmt.Println("a:", a, "| b:", b, "| *vptr:", *vptr) // b is now 200
}
