# CLI Color Library (gookit/color)

A Go library for command-line color output with 16/256/RGB color support, cross-platform compatibility, and HTML-like tag functionality.

**ALWAYS reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Bootstrap and Dependencies
- `go mod tidy` - Download and organize dependencies. Takes ~5 seconds on first run.
- `go mod download` - Pre-download dependencies if needed.

### Build Operations
- `go build ./...` - Build all packages. EXTREMELY FAST: ~0.04 seconds. NEVER needs timeout adjustment.
- `go build .` - Build main package only.
- No Makefile or special build scripts needed - standard Go tooling only.

### Testing and Validation
- `go test ./...` - Run all tests. Takes ~15 seconds verbose mode. Set timeout to 30 seconds minimum.
- `go test -coverprofile=profile.cov ./...` - Generate coverage report. Takes ~1-2 seconds. NEVER CANCEL.
- `go test -v ./...` - Verbose test output showing all test details.
- `go test -race ./...` - Run tests with race detection (longer execution time).

### Code Quality
- `go fmt ./...` - Format code (always run before committing).
- `go vet ./...` - Run Go static analysis.
- NO golangci-lint configuration exists - do not attempt to run golangci-lint.
- NO custom linting setup - use standard Go tools only.

### Running Examples
- `go run _examples/demo.go` - Basic color demonstration.
- `go run _examples/color_256.go` - 256-color palette demonstration.
- `go run _examples/color_rgb.go` - RGB/true color demonstration.
- `go run _examples/color_tag.go` - HTML-like tag functionality.
- `COLOR_DEBUG_MODE=on go run _examples/envcheck.go` - Debug color environment detection.

## Validation Scenarios

**ALWAYS run these validation steps after making changes:**

### Core Functionality Validation
1. **Basic Color Output**: `go run _examples/demo.go` - Verify basic color functions work.
2. **256-Color Support**: `go run _examples/color_256.go` - Verify extended color palette.
3. **RGB Color Support**: `go run _examples/color_rgb.go` - Verify true color functionality.
4. **Tag Processing**: `go run _examples/color_tag.go` - Verify HTML-like tag parsing.

### Environment Detection Testing
1. **Debug Mode**: `COLOR_DEBUG_MODE=on go run _examples/envcheck.go` - Verify color detection logic.
2. **Color Capability Detection**: Run examples in different terminal environments to test detection.

### Test Coverage Validation
1. **Run Full Test Suite**: `go test -v ./...` - Must pass all tests.
2. **Coverage Check**: `go test -coverprofile=profile.cov ./...` - Maintain >95% coverage.
3. **Race Detection**: `go test -race ./...` - Verify thread safety.

## Key Project Structure

### Core Library Files
- `color.go` - Main color API and basic colors
- `color_16.go` - 16-color (4-bit) support
- `color_256.go` - 256-color (8-bit) support  
- `color_rgb.go` - RGB/true color (24-bit) support
- `color_tag.go` - HTML-like tag parsing and rendering
- `style.go` - Color styles and style combinations
- `utils.go` - Color detection and platform utilities

### Platform-Specific Code
- `detect_env.go` - Environment-based color detection
- `detect_nonwin.go` - Unix/Linux/macOS color detection
- `detect_windows.go` - Windows color detection

### Testing and Examples
- `*_test.go` - Comprehensive test files (97.7% coverage)
- `_examples/` - Rich set of demonstration programs
- `testdata/` - Test data files
- `colorp/` - Alternative printer interface sub-package

### Example Programs Location
```
_examples/
├── demo.go              # Basic color demonstration
├── color_16.go          # 16-color examples
├── color_256.go         # 256-color palette
├── color_rgb.go         # RGB/true color examples
├── color_tag.go         # HTML-like tag examples
├── envcheck.go          # Environment detection testing
└── theme_*.go           # Various themed color examples
```

## Common Development Tasks

### Adding New Color Functions
1. Add function to appropriate `color_*.go` file
2. Add corresponding test in `*_test.go` file
3. **ALWAYS test**: `go test ./... && go run _examples/demo.go`
4. Update documentation if adding public API

### Modifying Color Detection
1. Edit `detect_*.go` files for platform-specific logic
2. **ALWAYS test with debug mode**: `COLOR_DEBUG_MODE=on go run _examples/envcheck.go`
3. Test on different platforms if possible
4. Verify tests pass: `go test ./...`

### Adding HTML Tag Support
1. Modify `color_tag.go` for tag parsing logic
2. Add tests to `color_tag_test.go`
3. **ALWAYS validate**: `go run _examples/color_tag.go`
4. Test edge cases with malformed tags

### Performance Optimization
1. Run benchmarks: `go test -bench=. -benchmem`
2. Profile if needed: `go test -cpuprofile=cpu.prof -memprofile=mem.prof`
3. **ALWAYS verify functionality still works** after optimizations

## Environment and Dependencies

### Go Version Requirements
- **Minimum**: Go 1.18
- **Tested**: Go 1.18 through 1.24
- Check version: `go version`

### Dependencies (go.mod)
- `github.com/stretchr/testify` - Testing framework
- `github.com/xo/terminfo` - Terminal info library
- `golang.org/x/sys` - System interfaces
- **All dependencies are lightweight** - no complex setup required

### Platform Support
- **Linux**: Full support (16/256/RGB colors)
- **macOS**: Full support (16/256/RGB colors)  
- **Windows**: Full support including CMD and PowerShell
- **WSL**: Supported with automatic detection

## Debugging and Troubleshooting

### Color Detection Issues
1. **Enable debug mode**: `COLOR_DEBUG_MODE=on`
2. **Check environment**: `go run _examples/envcheck.go`
3. **Verify TERM variable**: `echo $TERM`
4. **Test in different terminals**: cmd, PowerShell, iTerm, etc.

### Test Failures
1. **Run verbose tests**: `go test -v ./...`
2. **Check race conditions**: `go test -race ./...`
3. **Verify environment**: Some tests are environment-dependent
4. **Check platform-specific code**: Windows vs Unix behavior

### Build Issues
1. **Update dependencies**: `go mod tidy`
2. **Clear module cache**: `go clean -modcache` (rarely needed)
3. **Check Go version**: Must be 1.18+

## Timing Expectations

### Build Operations
- **Build time**: ~0.04 seconds (extremely fast)
- **Module tidy**: ~5 seconds on first run, <1 second thereafter
- **NEVER need extended timeouts for builds**

### Test Operations
- **Basic test run**: ~1-2 seconds
- **Verbose tests**: ~15 seconds  
- **Coverage generation**: ~1-2 seconds
- **Race detection**: ~30 seconds
- **Set minimum 30-second timeout for test commands**

### Example Execution
- **All examples**: <1 second each
- **Debug mode examples**: 1-2 seconds
- **No long-running examples or demos**

## CI/CD Integration

### GitHub Actions Workflow
- **File**: `.github/workflows/go.yml`
- **Platforms**: Ubuntu, Windows, macOS
- **Go versions**: 1.18 through 1.24
- **Coverage**: Automatically uploaded to Coveralls

### Pre-commit Checklist
1. `go fmt ./...` - Format code
2. `go vet ./...` - Static analysis  
3. `go test ./...` - Run all tests
4. `go run _examples/demo.go` - Basic functionality check
5. Verify no new build warnings

Remember: This is a mature, stable library with excellent test coverage. Most changes should be small and focused. Always validate color output works correctly across different terminal environments when making changes to color detection or rendering logic.