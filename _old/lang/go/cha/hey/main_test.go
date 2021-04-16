package main

import "os"

func Example_main() {
	os.Args = []string{"", "You"}
	main()
	os.Args = []string{"", ""}
	main()

	// Output:
	// Hey You
	// Hey

}

/*
	os.Args = []string{"", "You", "There"}
	main()
	os.Args = []string{"", "You There"}
	main()
	os.Args = []string{"", "something"}
	main()
*/
// Hey You
// Hey You
// Hey You There
// Hey something
