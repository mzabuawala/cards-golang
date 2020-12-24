/*
Pass by value types
-------------------
int
float
string
bool
struct

Pass by reference types -> Gotcha :)
-----------------------
Slices
Maps
Channels
Pointers
Functions

RAM VIEW FOR SLICE TYPE
=======================

----------------------
ADDR| VALUE
----|-----------------
1111| Len-Cap-PTR2HEAD
----|-----------------
1112| MURTUZA-ZABUA       ----> Becomes "Aziz Zabuawala"
----|-----------------
1113| Len-Cap-PTR2HEAD    <---- This gets copies when we pass so actual value change persists because both points to same data
----------------------
*/

package main

import "fmt"

func main() {
	slice := []string{"Murtuza", "Zabua"}
	updateSlice(slice)
	fmt.Println(slice) // WTF -> [Aziz Zabua]  :-o
	// Ealier we learnt that Go is pass by value langauge
}

func updateSlice(s []string) {
	s[0] = "Aziz"
}
