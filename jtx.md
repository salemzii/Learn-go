


# CHAPTER 4, ARRAYS, SLICES AND MAPS  
# Arrays 

## Declaring an array with go calculating the size 
	```
		array := [...]TYPE{ELEMENTS}
	```
## Declaring an array and initilizing some indexes

	```
		array := [LENGTH]TYPE{INDEX_0 : Value, INDEX_2 : Another_Value}
	```

Copying an array of pointers copies the pointer values and not the values that the pointers are pointing to.

An idiomatic way of creating a slice is to use a slice literal. It’s similar to creating an array, except you don’t specify a value inside of the [ ] operator.

Calculating the length and capacity for any new slice is performed using the following formula.

How length and capacity are calculated For slice[i:j] with an underlying array of capacity k

Length: j - i 
Capacity: k - i


The append operation is clever when growing the capacity of the underlying array. Capacity is always doubled when the existing capacity of the slice is under 1,000 elements. Once the number of elements goes over 1,000, the capacity is grown by a factor of 1.25, or 25%. This growth algorithm may change in the language over time.


GIA :===> pg 75; Listing 4.36

GIA :===> pg 75; SECTION 4.2.5 -> Passing slices between functions Passing a slice between two functions requires nothing more than passing the slice by value. Since the size of a slice is small, it’s cheap to copy and pass between functions.
On a 64-bit architecture, a slice requires 24 bytes of memory. The pointer field requires 8 bytes, and the length and capacity fields require 8 bytes respectively. Since the data associated with a slice is contained in the underlying array, there are no prob- lems passing a copy of a slice to any function. Only the slice is being copied, not the underlying array (see figure 4.22).


When you index a map in Go, it will always return a value, even when the key doesn’t exist. In this case, the zero value for the value’s type is returned. Iterating over a map is identical to iterating over an array or slice. You use the key- word range; but when it comes to maps, you don’t get back the index/value, you get back the key/value pairs.


4.3.4 Passing maps between functions :====> Passing a map between two functions doesn’t make a copy of the map. In fact, you can pass a map to a function and make changes to the map, and the changes will be reflected by all references to the map.


SUMMARY:
	 . Arrays are the building blocks for both slices and maps.  Slices are the idiomatic way in Go you work with collections of data. Maps are the way you work with key/value pairs of data. 
	 .The built-in function make allows you to create slices and maps with initial length and capacity. Slice and map literals can be used as well and support set- ting initial values for use. 
	  . Slices have a capacity restriction, but can be extended using the built-in func- tion append. 
	  . Maps don’t have a capacity or any restriction on growth.  The built-in function len can be used to retrieve the length of a slice or map.  The built-in function cap only works on slices. 
	  .Through the use of composition, you can create multidimensional arrays and slices. You can also create maps with values that are slices and other maps. A slice can’t be used as a map key. 
	  . Passing a slice or map to a function is cheap and doesn’t make a copy of the underlying data structure.





# CHAPTER 5: GO'S TYPE SYSTEM


There are two types of receivers in Go: value receivers and pointer receivers.


When you declare a method using a value receiver, the method will always be operating against a copy of the value used to make the method call.

When you call a method declared with a pointer receiver, the value used to make the call is shared with the method.

This is thanks to the pointer receiver. Value receivers operate on a copy of the value used to make the method call and pointer receivers operate on the actual value.

Strings, just like integers, floats, and Bool- eans, are primitive data values and should be copied when passed in and out of functions or methods.

Reference types in Go are the set of slice, map, channel, interface, and function types. When you declare a variable from one of these types, the value that’s created is called a header value. 

All the different header values from the different reference types contain a pointer to an underlying data structure. Each reference type also contains a set of unique fields that are used to manage the underlying data structure.

When the decision is made that a struct type value should not be mutated when something needs to be added or removed from the value, then it should follow the guidelines for the built-in and reference types.


## 5.4 INTERFACES

Polymorphism is the ability to write code that can take on different behavior through the implementation of types.

When a user-defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type. This assignment stores the value of the user-defined type into the interface value.

Method Sets define the rules around interface compliance.

To understand why values of type user don’t implement the interface, when an interface is implemented with a pointer receiver, you need to understand what method sets are. 

Method sets define the set of methods that are associated with values or pointers of a given type.


Methods Receivers Values ----------------------------------------------- 
(t T) == T and *T 
(t *T) == *T

If you implement an interface using a pointer receiver, then only pointers of that type implement the interface. If you implement an interface using a value receiver, then both values and pointers of that type implement the interface.

The question now is why the restriction? The answer comes from the fact that it’s not always possible to get the address of a value.

Example: 
	```go
		package main

		import (
			"fmt"
		)

		type Duration int
		func (d *Duration) pretty() string { return fmt.Sprintf("Duration: %d", *d) }	

		func main (){

			fmt.Println(Duration(15).pretty())
		}

		Error
		./main.go:12:27: cannot call pointer method pretty on Duration

	```

Because it’s not always possible to get the address of a value, the method set for a value only includes methods that are implemented with a value receiver.


## 5.5 Type embedding

Example: 

	```go 

		package main 

		import (
			"fmt"
		)


		type notifier interface {

			notify()
		}

		type user struct{
			name string
			email string
		}

		func (u *user)notify(){
			fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
		}

		type admin struct {
			user 
			level string
		}


		func sendNotification(n notifier) {
			n.notify()
		}

		func main(){

			ad := admin{
				user: user{
					name: "john smith",
					email: "john@yahoo.com",
				},
				level: "super",
			}

			sendNotification(&ad)
		}

	```
Thanks to inner type promotion, the implementation of the interface by the inner type has been promoted up to the outer type. That means the outer type now implements the interface, thanks to the inner type’s implementation.

What if the outer type wants to implement a similar implementation of the inner type?

	```go 
		type admin struct {
			user 
			level string
		}

		func (a *admin) notify() { 
			fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email) 
		}
	```

In cases like this, the (u *user).notify() implementation is not promoted, therefore the (a *admin).notify() method can be directly accessed when passed to an interface. In case where there's also a need to pass in the user.notify() implementation, this can be done by passing the user type to the interface e.g `sendNotification(&ad.user)` will send notification for user type.


## 5.6 Exporting and unexporting identifiers 
The ability to apply visibility rules to the identifiers you declare is critical for good API design. Go supports the exporting and unexporting of identifiers from a package to provide this functionality.


5.7 Summary 
	. User-defined types can be declared using the keyword struct or by specifying an existing type. 
	. Methods provide a way to add behavior to user-defined types. 
	. Think of types as having one of two natures, primitive or non-primitive. 
	. Interfaces are types that declare behavior and provide polymorphism. 
	. Type embedding provides the ability to extend types without the need for inheritance. 
	. Identifiers are either exported or unexported from packages.


# Concurrency

Concurrency in Go is the ability for functions to run independent of each other.
When a function is created as a goroutine, it’s treated as an independent unit of work that gets scheduled and then executed on an available logical processor.

Concurrency synchronization comes from a paradigm called communicating sequential processes or CSP. CSP is a message-passing model that works by communicating data between goroutines instead of locking data to synchronize access. The key data type for synchronizing and passing messages between goroutines is called a channel.

## 6.1 Concurrency vs Parallelism

A thread is a path of execution that’s scheduled by the operating system to run the code that you write in your functions.

The operating system schedules threads to run against physical processors and the Go runtime schedules goroutines to run against logical processors.

Concurrency is not parallelism. Parallelism can only be achieved when multiple pieces of code are executing simultaneously against different physical processors. Par- allelism is about doing a lot of things at once. Concurrency is about managing a lot of things at once


## Goroutines
A WaitGroup is a counting semaphore that can be used to maintain a record of run- ning goroutines. When the value of a WaitGroup is greater than zero, the Wait method will block.

Based on the internal algorithms of the scheduler, a running goroutine can be stopped and rescheduled to run again before it finishes its work.

Goroutines can only run in parallel if there’s more than one logical processor and there’s a physical processor available to run each goroutine simultaneously.

## Race Conditions 

When two or more goroutines have unsynchronized access to a shared resource and attempt to read and write to that resource at the same time, you have what’s called a race condition. Race conditions are the reason concurrent programming is compli- cated and has a greater potential for bugs. Read and write operations against a shared resource must always be atomic, or in other words, done by only one gorou- tine at a time.

Go has a special tool that can detect race conditions in your code. It’s extremely useful to find these types of bugs, especially when they’re not obvious. To use this built-in feature, run 

`go run -race <NAME_OF_FILE>`


One way we can fix and eliminate the race condition is by using the support Go has for synchronizing goroutines by locking down shared resources.

Another way to synchronize access to a shared resource is by using a mutex. A mutex is named after the concept of mutual exclusion. A mutex is used to create a critical section around code that ensures only one goroutine at a time can execute that code section.

When a resource needs to be shared between goroutines, channels act as a conduit between the goroutines and provide a mechanism that guarantees a synchronous exchange.

When declaring a channel, the type of data that will be shared needs to be specified. Values and pointers of built-in, named, struct, and reference types can be shared through a channel.

An unbuffered channel is a channel with no capacity to hold any value before it’s received.

A buffered channel is a channel with capacity to hold one or more values before they’re received.

In buffered channels there are certain conditions where a send or recieve operation will block.

A receive operation will block if there are no values in the channel to be retrieved. 

A send operation will block if there’s no available buffer to place the value being sent.

When a channel is closed, goroutines can still recieves from the channel but can no longer send to the channel.

6.6 Summary 
	. Concurrency is the independent execution of goroutines. 
	. Functions are created as goroutines with the keyword go. 
	. Goroutines are executed within the scope of a logical processor that owns a sin- gle operating system thread and run queue. 
	. A race condition is when two or more goroutines attempt to access the same resource. 
	. Atomic functions and mutexes provide a way to protect against race conditions. 
	. Channels provide an intrinsic way to safely share data between two goroutines. 
	. Unbuffered channels provide a guarantee between an exchange of data. Buff- ered channels do not.