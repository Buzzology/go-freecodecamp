# Learn Go Programming
FreeCodeCamp.org 6.5hr YouTube Go Tutorial - https://www.youtube.com/watch?v=YS4e4q9oBaU  

Up to: https://youtu.be/YS4e4q9oBaU?t=20042

## Environment setup
Compound gopath - is similar to a normal path var, seperate using semi-colon.  
Project requires a single `src` directory  
pkg is where intermediate binaries are stored e.g. temp for while generating binaries.  
Go install adds to the app's bin folder.

## Vars Notes
Shadowed variables is a term used to describe inner scope variables overriding higher scoped variables of the same name.  
There is no implicit type conversion (intentional)  

## Types Notes
### Boolean
Has a value of zero even if not initialised.  

### Numeric

#### Integer
8-64bit.
Because no implicit conversion remainders are dropped when dividing.

#### Float  

#### Complex
real, imag, hmmmm....

### Text
Can directly treat string as arrays of bytes (will need to convert to string to get the char)  
`+` can be used to concatenate  
Rune represents utf32, uses single quotes `'a'`,

## Constants
- Should not be all uppercase otherwise they'll be made public 
- iota: unique incrementing value
```
const (  // iota is scoped to each block
    c0 = iota  // c0 == 0
    c1 = iota  // c1 == 1
    c2 = iota  // c2 == 2
) 

const (  // reset to 0 due to scope
    a0 = iota  // a0 == 0
    a1 = iota  // a1 == 1
    a2 = iota  // a2 == 2
) 
```  
Can use `_` to discard initial value to ensure all are over zero:
```
const (
    _ = iota
    a1  // a1 == 1
    a2  // a2 == 2
    a3  // a3 == 3
) 
```

Can also apply an offset:
```
const (
    cat = iota + 5 // cat == 5
    dog  // dog == 6
    cow  // cow == 7
) 
```


## Arrays
Full expression: `grades := [3]int { 1, 5, 10 }`  
Infer the length: `grades := [...]int { 1, 5, 10 }`  
Empty array of 3: `grades := [3]int`  
Length of array: `len(grades)`  


## Slices
Slices are backed by arrays but may be part (or all) of the backing array. Think of as a "slice" of an array.  
Initialise: `grades := []int {1, 2, 3}`  
Initialise: `grades2 := grades[2:] // Slice from the 2nd element to the end` 
Initialise: `grades3 := grades[:2] // Slice from the first element up to but not including the second index`  

Note that unlike arrays slices all point to the same underlying data:  
```
a := []int {1,2,3}
b := a
b[1] = 5
fmt.Println(a) // [1 5 3]
fmt.Println(b) // [1 5 3]
```  

Make can be used to initialise the slice's backing array to certain length, e.g. 100: `a := make([]int, 3, 100)`

This is useful as a slice isn't constrained by it's initialised length. If you have a slice that exceeds it's backing array's length the entire array is copied to a new one to cater for the increased length. This can be expensive so creating it initially with the increased length helps to circumvent this.  

Can use the spread operator `...` in a similar manner to javascript to concatenate slices:
```
a = append(a, []int {2, 3, 4, 5}...) // appends 2,3,4,5 to a
```  

Use `cap` to determine the length of underlying array, `len` to determine length of the slice.  

### Delete an entry from a slice  
https://stackoverflow.com/a/20545912/522859  
```
m = append(m[:i], m[i+1:]...)
```


## Maps  
Initialise: `states := make(map[string]int)`  
Initialise:
```
states := map[string]int {
    "Qld": 9001,
    "Nsw": 0,
    "Tas": 5,
}
```  

Order of entries in a map is *not* guaranteed. 

Checking if a value exists:
```
state, ok := states["NZ"] // ok will be false if the key doesn't exist  
```  
Maps are passed by reference.


## Structs
```
type Docter struct {
    number int
    actorName string
    companions []string
}
...
aDoctor := Docter {
    number: 3,
    actorName: "Jon Pertwee",
    companions: []string {
        "Liz shaw",
        "Jo",
        "Sarah Smith",
    }
}
```

Can declare them inline but not common:
```
aDoctor := struct {name string } { name: "John" }
```
By default structs copy by value, not by reference. Can use `&` to point to the reference.  


Uses composition instead of inheritance.  
```
type Animal struct {
    Name string
    Origin string
}

type Bird struct {
    Animal // Note this
    SpeedKPH float32
    CanFly bool
}

func main() {
    b := Bird{}
    b.Name = "Emu" // Means we can assign a name etc
    b.Origin = "Australia"
    b.SpeedKPH = 48
    b.CanFly = false
    fmt.Println(b)
}
```

If using literal syntax the animal struct must be declared:
```
b := Bird {
    Animal: Animal { Name: "Emu", Origin: "Australia" },
    SpeedKPH: 48,
    CanFly: false,
}
```

Note that normally you would use interfaces instead of structs for this. One exeption might be to "embed" behaviour e.g. giving a controller functionality from the base controller.

### Tags
Might be used in conjunction with a validation library.  
```
type Animal struct {
    Name string `required max:"100"` <-- This is the tag
    Origin string
}

func main() {
    t := reflect.TypeOf(Animal{})
    field, _ := t.FieldByName("Name")
    fmt.Println(field.Tag)
}
```

## If statements  
Cannot have single line if statements, must have curly braces  


## Case statements
Doesn't support falling through by default, must use the `fallthrough` statement. However, using fallthrough will ignore the subsequent case check, it will *always* fallthrough. This allows for the normal `break` statement to be omitted. We can also use multiple cases in a single case:
```
switch 5 {
    case 1, 5, 10:
        ...
    case 2, 4, 6:
        ...
    default:
        ...
}
```  
Note that the values to check in each still cannot overlap. Can also use if statements etc:
```
i := 10
switch {
    case i <= 10:
        ...
    case i <= 20:
        ...
    default:
        ...
}
```
Note that when done in this fashion the values *can* overlap.

## Looping
All loops are done using for, but can get pretty much all normal functionality out of this. Still uses `continue` and `break`.

```
for initializer; test; increments {}
for test {}
for {}
```

Can use labels in loops to exit nested loops completely:
```
Loop:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Println(i * j)
            if i * j >= 3 {
                break Loop // Exits both loops
            }
        }
    }
```

Looping over an entire slice:
```
s := []int {1,2,3}
for k, v := range s {
    fmt.Println(k,v)
}
```

## Defer
defer executes statement at the end of a function but before it's returned. They are executed in *last in first out* fashion.  

Often used for closing resources, similar to dispose:  
```
res, err := http.Get("http://www.google.com/robots.txt")
if err != nil {
    log.Fatal(err)
}

robots, err := ioutil.ReadAll(res.Body)
defer res.Body.Close()

if err != nil {
    log.Fatal(err)
}

fmt.Printf("%s", robots)
```

One important thing to note with defer is that it takes the value of any variables at the point the defer statement is made:  
```
a := "state"
defer fmt.Println(a)
a = "end"
...
Output: "start"
```

## Panic and recover
Similar to a typical exception.
Defer statements are still executed if a panic is unhandled.

Can use `recover` to handle panics in a similar manner to try/catch but it is only useful in deferred functions. Current function won't continue but higher up the stack will be able to so long as the defer recovers it and doesn't "re-panic":
```
func main() {
    fmt.Println("start")
    panicker()
    fmt.Println("end")
}

func panicker() {
    fmt.Println("about to panic")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error: ", err)
            panic(err) // Re-throw if not able to handle
        }
    }
}
```  


## Pointers
Pointer types use an `*` as a prefix to the type being pointed to e.g. `*int - a pointer to an integer`  
Use the `addressof` operator `&` to get address of variable  
Dereference a pointer by preceding with an asterisk (*)  
Complex types e.g. structs are automatically dereferenced  
Can use the addressof operator (&) to get the address of a value type if it already exists:
```
ms := myStruct{foo: 42}
p := &ms  
```

By default all assignment operators are `copy` operations  
Slices and maps are exceptions because they contain pointers to underlying data (e.g. arrays)  

```
var a int = 42
var b *int = &a // b is declaring a pointer to type int. Is assigned a pointer to a
fmt.Println(a, b) // 42, 0x1040a124
fmt.Println(&a, b)  // 0x1040a124, 0x1040a124 (& is the "address of" operator)
fmt.Println(a, *b)  // 0x1040a124, 42 (* is the "dereferencing" operator)

*b = 14 // Dereference b and assign a value
fmt.Println(a, *b)  // 14 14
```

## Functions  
Can specify type at the end:
```
func main() {
    sayGreeting("Hello", "Batman")
}

func sayGreeting(greeting, name string) { // NOTE: string only specified for last var because they are the same type
    fmt.Println(greeting, name)
}
```

Create a slice of input (called a variadic parameter). Can only have one per function and must be at the end:  
```
func main() {
    s := sum(1, 2, 3, 4, 5)
    fmt.Println("The sum is ", *s)
}

func sum(values ... int) *int {  // <-- Here
    fmt.Println(values)
    result := 0
    for _, v := range values {
        result += v
    }
    return &result
}
```

Can return multiple results from a function. Common way of handling errors:
```
func main() {
    d := divide(5.0, 0.0) 
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot divide by zero")
    }

    return a / b, nil
}
```

Anonymous functions can be used:
```
func main() {
    func() {
        fmt.Println("Hello GO!")
    }() 
}
```

Can declare and pass functions as vars:  
```
func main() {
    var divide func(float64, float64) (float64, error)
    divide = func(a,b float64) (float64, error) {
        if b == 0.0 {
            return 0.0, fmt.Errorf("Cannot divide by zero")
        } else {
            return a / b, nil
        }
    }

    d, err := divide(5.0, 3.0)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(d)
}
```

## Methods
Functions that execute in a known context. 
```
func main() {
    g := greeter {
        greeting: "Hello",
        name: "Go",
    }
    g.greet()
}

type greeter struct {
    greeting string
    name string
}

func (g greeter) greet() {
    fmt.Println(g.greeting, g.name)
}
```

Note in above that we're not using a reference above so it's being passed by value. Any changes in the method won't affect the caller's instance. If we wanted to make it by reference we can do the following:
```
func (g *greeter) greet() {
    fmt.Println(g.greeting, g.name)
}
```

## Interfaces  
Interfaces are implemented implicitly (you don't have to declare their implementation - no implements keyword):  
```
func main() {
    var w Wrtier = ConsoleWriter{}
    w.Write([]byte("Hello Go!))
}

type Writer interface {
    Write([]byte) (int, error)
}

type ConsoleWriter struct {

}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
    n, err := fmt.Println(string(data))
    return n, err
}
```

By convention interfaces should end with `er`, especially if it's a single method interface e.g. Writer, Reader, Opener, Incrementer, ...  

Interfaces can be composed of multiple other interfaces. This is done the same way as embedding structs:
```
type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

type WriterCloser interface {
    Writer
    Closer
}
```

To safely convert types use something like the following:
```
var wc WriterCloser = NewBufferedWriterCloser()
wc.Write([]byte("Hello Youtube"))
wc.Close()

r, ok := wc.(io.Reader)
if ok {
    fmt.Println(r)
} else {
    fmt.Println("Conversion failed")
}
```

Empty interfaces can be used if casting to an unknown type. You can then test the type etc:
```
var myObj interface{} = NewBufferedWriterCloser()
if wc, ok := myObj.(WriterCloser); ok {
    ...
}
```

Type switches can also be used for unknown types:
```
var i interface{} = 0
switch i.(type) {
    case int:
        fmt.Println("i is an integer")
    case string:
        fmt.Println("i is a string")
    default:
        fmt.Println("I don't know what i is")
}
```

Implementing with values vs pointers:  
- Method set of `value` is all methods with value receivers  
- Method sets of `pointer` is all methods, they have access to both already  

Best practices:  
- Interfaces should be as small as possible
  - Larger interfaces should all be composed of multiple smaller interfaces  
- You don't need to export interfaces for other people to use like you would in c#. People can define their own interfaces.  
- Define functions and methods to receive interfaces where possible (unless you need access to the values)  


## Go Routines
Use `go` keyword to execute new *green* thread. Green threads aren't *real* threads but a lightweight version that allows for a significantly higher number of *threads* to be run.  

Can use anonymous functions for this:  
```
func main() {
    var msg = "Hello"
    go func() {
        fmt.Println(msg)
    }()
    time.Sleep(100 * time.Millisecond)
}
```

Note that accessing variables outside the scope can create race conditions:
```
func main() {
    var msg = "Hello"
    go func() {
        fmt.Println(msg) // Will often display goodbye instead of hello
    }()
    msg = "goodbye" 
    time.Sleep(100 * time.Millisecond)
}
```

Can avoid this by passing in the variable as an argument:
```
func main() {
    var msg = "Hello"
    go func(msg string) {
        fmt.Println(msg) // Will often display goodbye instead of hello
    }(msg)
    msg = "goodbye" 
    time.Sleep(100 * time.Millisecond)
}
```

Can remove the `sleep` by using wait groups:
```
var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)

    go func(msg string) {
        fmt.Println(msg)
		wg.Done()
    }(msg)
    msg = "goodbye"
	wg.Wait() 
}

func sayHello() {
	fmt.Println("hello")
}
```

`runtime.GOMAXPROCS` is used to set/limit the number of CPU threads. By default this is number of cores.  

### Best practices
- External facing libraries should try to avoid using goroutines if possible. Leave it to the consumer to decide when they want them.   
- Identity race conditions early, use `go run -race src/main.go`  


## Channels  
Created using make, note that they are strongly typed - can only use the data type they are initiated with:
```
ch := make(chan int)
```

Using a channel to send the value of 42 from one goroutine to another:  
```
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()
}
```

Multiple go routines can make use of a single channel, can also have more senders than receivers (or vice versa).  

You will reach a deadlock if a goroutine tries to send a message on a channel but there are no receivers.  

While you *can* have a single goroutine reading and writing from a channel you will normally want to split them up (is this true???). Can explicitly mark a channel as send or receive using the following notations:  
```
var wg = sync.WaitGroup{}

func main() {
    ch := make(chan int)
    wg.Add(2)

    // Read only
    go func(ch <-chan int) {
        i := <- ch // Read value from channel
        fmt.Println(i)
        wg.Done()
    }(ch)

    // Send only
    go func (ch chan<- int) {
        ch <- 42 // Send value of 42 across channel
        wg.Done()
    }(ch)

    wg.Wait()
}

```

Note that receive is `ch <-chan` and send is `ch chan<-`. A way to remember this is that its an arrow pointing to which way the data is going. `<-chan` is data flowing out of the channel, `ch chan<-` is data flowing into the channel.  

Receiving multiple values from a channel:
```
func main() {
    ch := make(chan int, 50) // Second arg is buffer length
    wg.Add(2)
    go func(ch <-chan int) {
        // Loop over all messages in channel
        for i := range ch {
            fmt.Println(i)
        }
        wg.Done()
    }

    go func(ch chan<- int) {
        ch <- 42
        ch <- 27
        close(ch) // Must close the channel when done or the receiver will break
        wg.Done()
    }
}
```

Can use the following syntax in scenarios where unsure if the channel will be closed:  
```
func main() {
    ch := make(chan int, 50) // Second arg is buffer length
    wg.Add(2)
    go func(ch <-chan int) {
        for {
            if i, ok := <- ch; ok {
                fmt.Println(i)
            } else {
                break
            }
        }
        wg.Done()
    }(ch)

    go func(ch chan<- int) {
        ch <- 42
        ch <- 27
        close(ch) // Must close the channel when done or the receiver will break
        wg.Done()
    }
}
```

### Select statements
Can use select statement to handle messages from multiple channels. The following example uses a secondary channel to signal when a logger goroutine should terminate.   
```
const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Using empty struct is similar to using a bool channel but requires no memory. A performance convention.

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	doneCh <- struct{}{} // Pass an empty struct to signal end
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for {
		select { // Blocks until a message is received on one of the channels
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %s\n", entry.time.Format("2020-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
		// NOTE: If you add a default statement it won't block
	}
}
```
