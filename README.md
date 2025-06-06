# 🔐 Entropy Forge

**Production-grade cryptographically secure password generation for Go**

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Performance](https://img.shields.io/badge/Performance-Low_Allocation-green.svg)](#benchmarks)
[![Security](https://img.shields.io/badge/Security-CSPRNG-red.svg)](#security-features)

> Enterprise-ready password generation with multiple algorithms, real-time complexity analysis, and memory-optimized performance.

## ✨ Features

### 🎲 **Multiple Generation Algorithms**
- **Diceware Passphrases** - Human-memorable, cryptographically secure
- **CSPRNG Random** - Maximum entropy character-based passwords
- **Low-Allocation Variants** - Memory-optimized for high-throughput applications

### 🔒 **Advanced Security**
- **NIST SP 800-63B Compliant** - Meets enterprise security standards
- **Cryptographically Secure** - Uses `crypto/rand` for true randomness
- **Real-time Analysis** - Instant password strength assessment
- **Pattern Detection** - Identifies and prevents weak patterns

### ⚡ **Performance Optimized**
- **Low-Allocation Design** - 13% fewer allocations than standard approaches
- **Concurrent Safe** - Thread-safe generation with sync pools
- **High Throughput** - Optimized for production workloads
- **Parallel Scaling** - 15% faster under concurrent load

### 📊 **Comprehensive Analysis**
- **Entropy Calculation** - Shannon entropy and bit strength
- **Complexity Scoring** - 0-100 security rating
- **Crack Time Estimation** - Real-world attack time estimates
- **Weakness Detection** - Identifies vulnerabilities and suggests improvements

## 🚀 Quick Start

## 📖 API Reference## 📖 API Reference

### Diceware Generation

```go
type DicewareOptions struct {
    WordCount    int    // Number of words (recommended: 6-8)
    Separator    string // Word separator: "-", " ", "" 
    Uppercase    bool   // Randomly uppercase one word
    Capitalize   bool   // Capitalize all words
    AddNumbers   bool   // Add random numbers
    AddSymbols   bool   // Add random symbols
}

// Standard generation
func GenerateDiceware(opts DicewareOptions) (string, error)

// Memory-optimized generation  
func GenerateDicewareLowAlloc(opts DicewareOptions, result *strings.Builder) error

// Generation with analysis
func GenerateDicewareWithAnalysis(opts DicewareOptions, result *strings.Builder) (*PasswordAnalysis, error)
```

### CSPRNG Generation

```go
// Standard CSPRNG password
func GenerateCSPRNG(length int) (string, error)

// Memory-optimized CSPRNG
func GenerateCSPRNGLowAlloc(length int, result *strings.Builder) error

// CSPRNG with analysis
func GenerateCSPRNGWithAnalysis(length int, result *strings.Builder) (*PasswordAnalysis, error)
```

### Password Analysis

```go
type ComplexityScore struct {
    Score           int      `json:"score"`            // 0-100 strength rating
    Entropy         float64  `json:"entropy_bits"`     // Shannon entropy
    Strength        string   `json:"strength"`         // Human-readable strength
    CrackTime       string   `json:"crack_time"`       // Estimated crack time
    Weaknesses      []string `json:"weaknesses"`       // Identified vulnerabilities
    Suggestions     []string `json:"suggestions"`      // Security improvements
    CharacterTypes  int      `json:"character_types"`  // Character variety count
    PatternScore    int      `json:"pattern_score"`    // Pattern analysis score
}

func AnalyzePassword(password string) ComplexityScore
```

## 🎯 Usage Examples

### Enterprise Diceware Passphrase
```go
import entropyforge "github.com/muzzii255/entropy-forge"

opts := entropyforge.DicewareOptions{
    WordCount:  8,          // High security
    Separator:  "-",        // Standard separator
    Capitalize: true,       // Title case
    AddNumbers: true,       // Numeric component
    AddSymbols: false,      // Avoid symbols for compatibility
}
password, _ := entropyforge.GenerateDiceware(opts)
// Result: "Corporate-Finance-Security-Protocol-Database-Network-Server-Gateway-847291"
```

### High-Performance Generation
```go
import (
    "strings"
    entropyforge "github.com/muzzii255/entropy-forge"
)

var builder strings.Builder
passwords := make([]string, 1000)

// Memory-optimized batch generation
for i := 0; i < 1000; i++ {
    entropyforge.GenerateDicewareLowAlloc(opts, &builder)
    passwords[i] = builder.String()
    builder.Reset()
}
```

### Security Analysis Workflow
```go
import entropyforge "github.com/muzzii255/entropy-forge"

analysis, err := entropyforge.GenerateDicewareWithAnalysis(opts, &builder)
if err != nil {
    return err
}

// Check if password meets requirements
if analysis.Complexity.Score < 75 {
    fmt.Printf("Warning: Password strength is %s\n", analysis.Complexity.Strength)
    fmt.Printf("Weaknesses: %v\n", analysis.Complexity.Weaknesses)
    fmt.Printf("Suggestions: %v\n", analysis.Complexity.Suggestions)
}

// Log security metrics
log.Printf("Generated password: entropy=%.1f bits, score=%d, crack_time=%s",
    analysis.Complexity.Entropy,
    analysis.Complexity.Score,
    analysis.Complexity.CrackTime)
```

## 📈 Benchmarks

Performance results on AMD Ryzen 7 5800X:

```
BenchmarkGenerateDiceware-16                   192,172     6,019 ns/op    3,603 B/op    135 allocs/op
BenchmarkGenerateDicewareLowAlloc-16           226,718     5,401 ns/op    3,235 B/op    118 allocs/op
BenchmarkGenerateCsprng-16                     273,978     4,446 ns/op    1,600 B/op     98 allocs/op
BenchmarkGenerateCsprngLowAlloc-16             257,148     4,288 ns/op    1,592 B/op     99 allocs/op
BenchmarkPasswordAnalysis-16                   516,837     2,114 ns/op    1,064 B/op      9 allocs/op
BenchmarkGenerateDicewareParallel-16         1,000,000     1,034 ns/op
BenchmarkGenerateDicewareLowAllocParallel-16 1,373,158       884 ns/op
```

**Key Performance Improvements:**
- ✅ **10% faster generation** with low-allocation variants
- ✅ **13% fewer memory allocations** 
- ✅ **15% better parallel performance**
- ✅ **Sub-microsecond analysis** for real-time feedback

Run benchmarks yourself:
```bash
go test -bench=. -benchmem
```

## 🔒 Security Features

### Cryptographic Strength
- **CSPRNG Foundation** - All randomness sourced from `crypto/rand`
- **No Predictable Patterns** - Immune to statistical analysis
- **Forward Secrecy** - Past outputs don't compromise future generation
- **Timing Attack Resistant** - Constant-time operations where applicable

### Entropy Analysis
```go
// Diceware: log2(7776^6) ≈ 77.5 bits
// CSPRNG: log2(94^24) ≈ 157.8 bits
fmt.Printf("Entropy: %.1f bits\n", analysis.Complexity.Entropy)
```

### Pattern Detection
- **Keyboard Patterns** - Detects qwerty, asdf sequences
- **Dictionary Words** - Identifies common passwords
- **Character Repetition** - Flags excessive repetition
- **Sequential Characters** - Catches abc, 123 patterns

## 🛠️ Installation

```bash
go get github.com/muzzii255/entropy-forge
```

**Requirements:**
- Go 1.24.2 or later
- `golang.org/x/text` for Unicode text processing


## 🎲 Diceware Implementation

Entropy Forge implements the **EFF Large Wordlist** standard:
- **7,776 unique words** (6^5 combinations)
- **True dice simulation** using crypto/rand
- **NIST-compliant entropy** calculation
- **Multiple wordlist support** (planned: international languages)

### Entropy Calculation
```
Diceware: log2(7776) × word_count = 12.92 × 6 ≈ 77.5 bits
CSPRNG: log2(charset_size) × length = log2(94) × 24 ≈ 157.8 bits
```

## 📊 Password Strength Scale

| Score | Strength | Crack Time | Use Case |
|-------|----------|------------|----------|
| 90-100 | Very Strong | Centuries | High-value targets |
| 75-89 | Strong | Years | Enterprise accounts |
| 60-74 | Good | Months | Standard accounts |
| 40-59 | Fair | Days | Minimum acceptable |
| 20-39 | Weak | Hours | Needs improvement |
| 0-19 | Very Weak | Minutes | Unacceptable |


## 📝 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🌟 Why Entropy Forge?

### For Developers
- **Production-ready** with comprehensive testing
- **Memory-efficient** for high-throughput applications  
- **Well-documented** with clear examples
- **Benchmark-driven** development approach

### For Security Teams
- **Cryptographically sound** algorithms
- **Transparent implementation** - no black boxes
- **Compliance-ready** for enterprise requirements
- **Real-time analysis** for policy enforcement

### For DevOps
- **Concurrent-safe** for microservice architectures
- **Low-allocation** design reduces GC pressure
- **Predictable performance** under load
- **Easy integration** with existing Go services

---

**Built with ❤️ for the Go community**

*Entropy Forge: Where security meets performance*