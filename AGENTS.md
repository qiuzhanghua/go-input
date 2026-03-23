# Go Input - Agent Guidelines
# Go Input - Agent 指南

This is a Go library for reading user input from the console. It supports Unix/Windows, masking, validation, and looping prompts.
这是一个 Go 库，用于从控制台读取用户输入。支持 Unix/Windows、掩码输入、验证和循环提示。

## Build, Lint & Test Commands
## 构建、Lint 和测试命令

```bash
# Run all tests
# 运行所有测试
go test ./...

# Run tests with verbose output
# 运行测试并显示详细输出
go test -v ./...

# Run a single test by name
# 按名称运行单个测试
go test -v -run "TestAsk" ./...
go test -v -run "TestSelect" ./...

# Run benchmarks
# 运行基准测试
go test -bench=. ./...

# Run go vet (static analysis)
# 运行 go vet（静态分析）
go vet ./...

# Build the library
# 构建库
go build ./...

# Tidy dependencies
# 整理依赖
go mod tidy
```

## Project Structure
## 项目结构

- `input.go` - Main UI struct, Options, and helpers
- `input.go` - 主 UI 结构体、选项和辅助函数
- `ask.go` - Ask() function for free-form input
- `ask.go` - 用于自由格式输入的 Ask() 函数
- `select.go` - Select() function for numbered selection
- `select.go` - 用于编号选择的 Select() 函数
- `read.go` - Core reading logic with masking support
- `read.go` - 支持掩码的核心读取逻辑
- `read_unix.go` - Unix terminal raw mode (linux/darwin/freebsd)
- `read_unix.go` - Unix 终端原始模式 (linux/darwin/freebsd)
- `read_windows.go` - Windows terminal raw mode
- `read_windows.go` - Windows 终端原始模式
- `translate.go` - Internationalization messages
- `translate.go` - 国际化消息
- `*_test.go` - Test files with Example functions for documentation
- `*_test.go` - 带 Example 函数的测试文件（用于文档）

## Code Style Guidelines
## 代码风格指南

### Go Version
### Go 版本
- Target: **Go 1.26** (use Go 1.26 features freely)
- 目标版本: **Go 1.26** (可自由使用 Go 1.26 特性)
- Minimum version is defined in `go.mod`
- 最低版本定义在 `go.mod` 中

### Build Tags
### 构建标签
- Use modern `//go:build` format (not `// +build`)
- 使用现代的 `//go:build` 格式（不是 `// +build`）
- Example: `//go:build linux || darwin || freebsd`
- 示例: `//go:build linux || darwin || freebsd`

### Imports
### 导入
- Standard library first, then third-party
- 优先标准库，然后是第三方库
- Group: stdlib | external
- 分组: stdlib | external
- No unused imports (enforced by compiler)
- 不使用未导入的包（编译器强制检查）

### Error Handling
### 错误处理
- Use `errors.New()` for static messages
- 静态消息使用 `errors.New()`
- Use `fmt.Errorf()` with `%w` for error wrapping: `fmt.Errorf("failed to x: %w", err)`
- 使用 `fmt.Errorf()` 和 `%w` 进行错误包装: `fmt.Errorf("failed to x: %w", err)`
- Never use `fmt.Errorf(T("..."))` - use `errors.New(T("..."))` instead
- 永远不要使用 `fmt.Errorf(T("..."))` - 请改用 `errors.New(T("..."))`
- Return errors early, not late
- 尽早返回错误，而不是晚返回

### Formatting
### 格式化
- 4-space indentation (tabs not allowed in this repo)
- 4 空格缩进（本仓库不允许使用制表符）
- No blank lines inside functions unnecessarily
- 函数内不必要的空行
- Blank line between import groups and between type declarations
- 导入组之间和类型声明之间需要空行

### Naming Conventions
### 命名约定
- `UI` - exported struct, PascalCase
- `UI` - 导出的结构体，PascalCase
- `ValidateFunc`, `readOptions` - exported/unexported types, PascalCase
- `ValidateFunc`, `readOptions` - 导出/未导出的类型，PascalCase
- `defaultWriter`, `maskString` - unexported vars/functions, camelCase
- `defaultWriter`, `maskString` - 未导出的变量/函数，camelCase
- Acronyms: use all caps for 2-letter (URL, HTTP, ID), PascalCase for longer (Repository, XML)
- 缩写词: 2 个字母用全大写 (URL, HTTP, ID)，更长的用 PascalCase (Repository, XML)
- Error messages: lowercase, no period at end (e.g., `"input is empty"`)
- 错误消息: 小写，末尾无句点（例如 `"input is empty"`）

### Types
### 类型
- Use `io.Reader` and `io.Writer` for I/O abstractions (not `*os.File` directly)
- 使用 `io.Reader` 和 `io.Writer` 进行 I/O 抽象（不直接使用 `*os.File`）
- Use `sync.Once` for one-time initialization
- 使用 `sync.Once` 进行一次性初始化
- Pointer receiver for methods on `UI` struct
- 对 `UI` 结构体的方法使用指针接收器

### Concurrency
### 并发
- Use `sync.OnceFunc` (Go 1.21+) instead of `sync.Once` + wrapper
- 使用 `sync.OnceFunc` (Go 1.21+) 而不是 `sync.Once` + 包装器
- Channel buffer size of 1 for signal notification: `make(chan os.Signal, 1)`
- 信号通知的通道缓冲区大小为 1: `make(chan os.Signal, 1)`
- Always call `defer signal.Stop(ch)` after `signal.Notify`
- 在 `signal.Notify` 之后始终调用 `defer signal.Stop(ch)`

### Testing
### 测试
- Table-driven tests with struct slices
- 使用结构体切片的表驱动测试
- Use `t.Fatalf` for fatal errors in tests
- 测试中的致命错误使用 `t.Fatalf`
- Example functions (Example*) serve as documentation
- Example 函数（Example*）作为文档
- Use `io.Discard` instead of `ioutil.Discard`
- 使用 `io.Discard` 而不是 `ioutil.Discard`
- Test names: `TestFunctionName` or `TestFunctionName_Scenario`
- 测试名称: `TestFunctionName` 或 `TestFunctionName_Scenario`

### Performance
### 性能
- Use `bufio.Reader` for line-based reading
- 使用 `bufio.Reader` 进行逐行读取
- Minimize allocations in hot paths
- 最小化热路径中的内存分配
- Use `strings.Clone` (Go 1.20+) when string copy is needed
- 需要字符串拷贝时使用 `strings.Clone` (Go 1.20+)

## Internationalization (i18n)
## 国际化 (i18n)

The `T()` function provides i18n support. All user-facing strings use `T("key")`:
`T()` 函数提供国际化支持。所有面向用户的字符串都使用 `T("key")`:
```go
fmt.Fprint(i.Writer, T("go-input.ask.enter-value"))
```

Message keys follow pattern: `go-input.<module>.<message>`
消息键遵循模式: `go-input.<module>.<message>`
