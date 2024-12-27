# 30-Day Golang Challenge

Welcome to the 30-day Golang challenge! This challenge is designed for intermediate learners who want to deepen their understanding of Go. Over the next 30 days, we will explore essential Go concepts and dive into advanced topics while building a solid foundation for real-world Go applications.

## Learning Plan

### **Days 1-6 (Foundational Review & Intermediate Start)**

#### **Day 1: Functions**
- Function syntax (`func` keyword)
- Return types and multiple return values
- Named return values
- Function expressions (anonymous functions)
- Basic recursion

#### **Day 2: Methods & Structs**
- Structs: defining and using them
- Methods on structs (value receiver vs pointer receiver)
- Method receivers (when to use pointer vs value)
- Struct tags
- Struct embedding (composition)

#### **Day 3: Arrays, Slices, Maps**
- Arrays: declaration, initialization, and access
- Slices: defining, slicing, and appending
- Maps: creating, reading, updating, and deleting keys
- Difference between arrays and slices
- Passing slices to functions

#### **Day 4: Defer, Panic, Recover**
- `defer` keyword
- Panics and handling errors
- Recovering from panics
- Best practices for using panic/recover

#### **Day 5: Go Routines & Channels**
- Go routines: creating and managing them
- Channels: defining, sending, and receiving data
- Buffered vs unbuffered channels
- Synchronizing Go routines with channels

#### **Day 6: WaitGroups & Select**
- WaitGroup: creating and waiting for Go routines
- `sync` package and `WaitGroup`
- Using `select` with multiple channels

---

### **Assessment 1 (Days 1-6)**

**Project: Simple Task Manager**

Create a simple command-line Task Manager application that allows you to:
1. **Define a Task** - Each task has a name, description, and a deadline.
2. **Add Tasks** - Use slices to store tasks and dynamically add new tasks.
3. **View Tasks** - Display all tasks in a readable format.
4. **Mark Tasks as Completed** - Use maps to track the task completion status.
5. **Concurrency** - Use Go routines to simulate task completion by printing "Task completed" for each task after a random delay.
6. **Error Handling** - If a task has no description or name, panic and handle it with a recovery mechanism.

Use:
- Structs for task representation.
- Channels for task completion notifications.
- WaitGroups to synchronize the completion of all tasks.

---

### **Days 7-12 (Advanced Data Structures & Functions)**

#### **Day 7: Interfaces**
- Defining interfaces
- Implicit vs explicit implementation
- Empty interfaces
- Type assertion

#### **Day 8: Anonymous Functions & Closures**
- Closures and their use cases
- Creating anonymous functions
- Function expressions as arguments and returns

#### **Day 9: Error Handling & Custom Errors**
- Error interface and handling errors
- Creating custom error types
- The `errors` package

#### **Day 10: Generics (Intro)**
- Go 1.18+ generics syntax
- Type constraints
- Type parameters and functions

#### **Day 11: Generics (Advanced)**
- Generic types and functions with constraints
- Working with slices, maps, and structs using generics
- Generic interfaces

#### **Day 12: Package Organization & Testing**
- Creating packages and organizing code
- Unit testing in Go with `testing` package
- Writing basic test functions
- Running tests and benchmarks

---

### **Assessment 2 (Days 7-12)**

**Project: Generic Data Structures**

Create a generic stack data structure that can store any type of data. Implement the following features:
1. **Push** - Add an element to the stack.
2. **Pop** - Remove and return the top element.
3. **Peek** - Return the top element without removing it.
4. **Size** - Return the current size of the stack.
5. **Testing** - Write unit tests to validate the stack operations for different data types (int, string, custom structs).
6. **Error Handling** - Handle the case where you try to pop from an empty stack by returning an error.

Use:
- Go generics for implementing a stack that works with any data type.
- Custom error handling for empty stack situations.

---

### **Days 13-18 (Concurrency & Advanced Concepts)**

#### **Day 13: Mutexes & Sync Package**
- Synchronization in Go
- Using `sync.Mutex` to prevent race conditions
- Read/Write locks with `sync.RWMutex`

#### **Day 14: Context Package**
- Understanding the `context` package
- Using `context` for managing cancellation and timeouts
- Context propagation across Go routines

#### **Day 15: Go Routines Deep Dive**
- Advanced Go routine patterns
- Creating pools of workers using Go routines
- Use cases for Go routines in real-world scenarios

#### **Day 16: Channel Patterns**
- Channel Direction (Send only, Receive only)
- Fan-out and Fan-in patterns with Go routines
- Using `select` with multiple channels efficiently

#### **Day 17: Profiling and Optimization**
- Basic profiling tools in Go
- Identifying bottlenecks using the `pprof` package
- Optimizing Go code for performance

#### **Day 18: Reflection**
- Using the `reflect` package
- Reflection for types, methods, and structs
- Use cases for reflection (e.g., serialization)

---

### **Assessment 3 (Days 13-18)**

**Project: Worker Pool with Context and Mutex**

Create a worker pool that processes tasks concurrently. Implement the following:
1. **Worker Pool** - Create a pool of Go routines that can process tasks concurrently. Use a `sync.WaitGroup` to wait for all tasks to complete.
2. **Task Queue** - Implement a queue of tasks that need to be processed. Use channels to send tasks to the workers.
3. **Mutex Locking** - Ensure that shared resources (e.g., task completion status) are protected from race conditions by using `sync.Mutex`.
4. **Context Cancellation** - Use the `context` package to cancel all tasks if the operation takes longer than a specified timeout.
5. **Profiling** - Use `pprof` to profile the application and identify performance bottlenecks.

---

### **Days 19-24 (Advanced Techniques & Real-World Projects)**

#### **Day 19: Building REST APIs with Go**
- Introduction to building HTTP servers
- Handling requests with `net/http`
- Routing and middleware
- Creating simple RESTful APIs

#### **Day 20: Working with JSON**
- Encoding and decoding JSON
- Custom marshaling and unmarshaling
- Handling nested JSON

#### **Day 21: WebSockets & Real-time Communication**
- Introduction to WebSockets in Go
- Setting up a WebSocket server
- Sending and receiving real-time data

#### **Day 22: File I/O & Handling Files**
- Reading and writing files in Go
- File operations (open, close, read, write)
- Working with directories and file paths

#### **Day 23: Go Modules & Dependency Management**
- Understanding Go modules
- Creating and using Go modules
- Managing dependencies with `go.mod` and `go.sum`

#### **Day 24: Testing Advanced Features**
- Writing integration tests
- Mocking dependencies in tests
- Testing concurrency

---

### **Assessment 4 (Days 19-24)**

**Project: Real-Time Chat Server with WebSockets**

Create a simple real-time chat application using WebSockets:
1. **WebSocket Server** - Set up a WebSocket server to handle incoming client connections.
2. **Client Communication** - Allow clients to send messages to each other in real time.
3. **Broadcasting** - Implement a broadcast feature to send messages to all connected clients.
4. **JSON Parsing** - Use JSON for sending and receiving messages.
5. **Testing** - Write tests to simulate multiple WebSocket connections and ensure the chat functionality works properly.
6. **File I/O** - Log chat history to a file, and implement functionality to load chat history when the server restarts.

---

### **Days 25-30 (Capstone Project & Final Review)**

#### **Day 25-27: Capstone Project (1st Phase - Design and Setup)**
- Plan and start a real-world project (e.g., CLI app, REST API, or utility tool)
- Focus on structuring the project, defining features, and setting up testing and documentation.

#### **Day 28-29: Capstone Project (2nd Phase - Coding & Optimization)**
- Implement core features of the project
- Focus on concurrency, data management, and error handling
- Optimize for performance, readability, and maintainability

#### **Day 30: Capstone Project (3rd Phase - Final Review & Refactoring)**
- Review your code
- Refactor for better structure and efficiency
- Write final documentation
- Submit or deploy your project (e.g., create a GitHub repo or share the executable)

---

## Additional Tips

- **Daily Practice**: Try coding for at least 2-3 hours every day. You can break it up into smaller sessions if needed.
- **Keep a Journal**: Track what you’ve learned, issues you’ve encountered, and solutions.
- **Explore Go’s Documentation**: Refer to the [Go Documentation](https://golang.org/doc/) for in-depth information about each topic.
- **Join a Community**: Consider joining forums like Go's Reddit or Discord communities to get feedback on your code.
