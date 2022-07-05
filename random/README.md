# Random Go files I have used in training.

## Notes

In Go, concurrent tasks are called **goroutines**. 

Goroutines allow for concurrency: pausing one task to work others. 
And in some situations they allow parallelism: working on mulitple tasks simultaneosly.

Sitenote: In Linux, you need to set up credential manager for your Git to stop it for asking your credential time and again. 

Go channel pauses the main goroutine for a set amount of time so the other goroutines can run. 

We cannot use function return values in a **go** statement. So we can use **channel** to return values. 

```go
size = go responseSize("https://example.com") // You're saying, "go run this; I'm not going to wait"
fmt.Println(size) // But then what is the return value? 

```

**Channels** is a way to communicate between goroutines.
Not only do channels allow you to send values from one goroutine to another, they ensure the sending 
goroutine has sent the value before the receiving goroutine attempts to use it.

```go
var myChannel chan float64 // Declare a variable to hold a channel.
myChannel = make(chan float64) // Actually create the channel.

or


myChannel := make(chan float64)
```
Channels make sure that the sender sent its value and the receiver received the value: synchronize.

Composite types: slice, maps, structs.







