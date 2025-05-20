/*
🔹 What is an Interface in Go?
An interface in Go is a set of method signatures.
If a type (like a struct) implements all the methods of an interface, it is said to satisfy that interface — no keyword or special declaration needed!

Think of it as a contract — “if you can do X and Y, you’re part of this interface.”


Imagine you say:

“Hey, I don’t care what type of object you are, as long as you can do a certain action, I’ll work with you.”

That’s an interface — it's just a set of behaviors (method signatures).

💬 “If your type has this method, you're allowed in.”



BASIC SYNTAX

type InterfaceName interface {
    Method1()
    Method2(param int) string
}

*/


package main

import (
	"fmt"
	"math"
)


