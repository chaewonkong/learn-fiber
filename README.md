# ⚡ Blazing Fast Fiber App

| includes           |
| ------------------ |
| gorm configuration |
| logger setting     |
| Makefile           |
| unit test          |
| mockery            |
| Dockerfile         |

It's implemented with clean architecture, yet does not uses DI(Dependency Injection).

## Why not DI?

First of all, unlike Java/Springboot, DI is not mandatory in most Go webframeworks.
DI makes it easier to substitute one module by another, give better test experience, in the expense of code complexity and performance.

With prevous experiences of using `uber/fx` and `google/wire`, I personally find no overwhelming benefits of DI over it's fallbacks.

I personally believe simplicity is the essence of Go, so in this project I omitted implementation of DI intentionally.
