# Learn Go: Speedrun to Deep Knowledge
This guide is a concentrated roadmap to learn Go (Golang) fast while building a deep, practical understanding of the language, ecosystem, and internals.
Goals
- Acquire idiomatic Go knowledge.
- Master concurrency, tooling, testing, and performance.
- Understand runtime, memory model, compiler and linking basics.
- Build realistic projects and profiling skills.

Suggested timeline (12 weeks, flexible)
- Week 1: Fundamentals and syntax
- Weeks 2–3: Standard library, tooling, modules
- Weeks 4–5: Concurrency, channels, patterns
- Week 6: Testing, benchmarking, profiling
- Weeks 7–8: Networking, HTTP, gRPC
- Weeks 9–10: Internals (runtime, memory, GC), compiler basics
- Weeks 11–12: Projects, review, advanced topics

Core topics (with minimal required activities)
1) Basics
	- Setup: go toolchain, GOPATH vs modules, go env
	- Syntax: packages, imports, functions, methods, structs, interfaces
	- Error handling: errors pkg, sentinel vs wrapped errors (fmt.Errorf, errors.Is/As)
	- Exercises: implement CLI calculator, basic file I/O

2) Types and idioms
	- Slices, arrays, maps, pointers
	- Zero values, composite literals, tags
	- Idiomatic Go: effective Go reading, naming, small functions, error-first APIs
	- Exercises: implement a small in-memory key-value store

3) Concurrency
	- Goroutines, channels, select, buffered vs unbuffered
	- Patterns: worker pools, fan-in/fan-out, pipelines, cancellation with context
	- Synchronization: sync.Mutex, RWMutex, atomic operations
	- Memory model basics: happens-before, synchronization primitives
	- Exercises: build concurrent crawler, parallel map-reduce

4) Standard library & tooling
	- fmt, io, bufio, net/http, context, encoding/json, time, os/exec
	- go fmt, go vet, golangci-lint, go list, go test, go build, go run
	- Modules: go.mod, semantic import versions, replace directives
	- Exercises: implement REST API server, client with retries and backoff

5) Testing, benchmarking, profiling
	- Table-driven tests, testify vs stdlib, subtests, coverage
	- Benchmarks: testing.B, b.N, profiling with pprof (cpu, heap), trace
	- Race detector: go test -race
	- Exercises: write full test suite for server, profile and optimize hotspots

6) Networking & distributed systems
	- net/http idioms, middleware, HTTP/2 basics
	- gRPC: proto, streaming, interceptors, deadlines
	- Serialization: JSON vs protobuf, msgpack, etc.
	- Exercises: build gRPC service with client streaming and auth

7) Internals & performance
	- Go runtime overview: scheduler, goroutine stacks, GC basics
	- Memory allocation: stack vs heap, escape analysis
	- Compiler toolchain: go vet findings, SSA backend basics, build tags
	- Profiling: flamegraphs, pprof interpretation, reducing allocations
	- Exercises: find and fix allocations in a hot path, reduce GC pressure

8) Advanced topics
	- Generics (type parameters): patterns, constraints, common pitfalls
	- Unsafe package, cgo, syscalls (use sparingly, understand cost)
	- High-performance networking (netpoll, fasthttp), zero-copy techniques
	- Plugin systems, embedding, code generation (go:generate)

Hands-on projects (increasing difficulty)
- CLI tool with subcommands, config, and tests
- HTTP REST API with DB (Postgres), migrations, and auth
- Concurrent worker pool processing jobs from a queue (Redis/NATS)
- gRPC microservice with streaming and TLS
- Profiling challenge: improve throughput/latency by 2x

Learning resources
- Official docs: golang.org/doc, pkg.go.dev
- Effective Go and Go Proverbs
- Go blog (blog.golang.org) for internals and patterns
- Books: "The Go Programming Language" (Alan Donovan & Brian Kernighan), "Concurrency in Go" (Katherine Cox-Buday)
- Tools: delve (debugger), pprof, go vet, staticcheck, golangci-lint

Practical study routine
- Daily: 1–2 hours of focused coding + 30m reading of blog/docs
- Weekly: one small project or feature, write tests, profile and optimize
- Code review: read open-source Go code (docker, etcd, kubernetes subpackages)

Checklist before claiming mastery
- Comfortable writing idiomatic, tested, benchmarked Go code
- Can reason about goroutine scheduling, memory allocation, and GC behavior
- Able to profile and fix performance issues reliably
- Understand module versioning and common build/deploy patterns

Notes and tips
- Favor simple, readable solutions. Optimize after profiling.
- Learn to read compiler/tool output; add -tags and build constraints only when necessary.
- Use contexts for cancellation/timeouts; never ignore returned errors.

Start: set up Go 1.20+ (or latest stable), run go mod init, and pick the first small project above.

```
