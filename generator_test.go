package entropyforge

import (
    "strings"
    "testing"
)

func BenchmarkGenerateDiceware(b *testing.B) {
    opts := DicewareOptions{
        WordCount: 6,
        Separator: "-",
        Capitalize: true,
        AddNumbers: true,
    }
    
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        _, err := Diceware(opts)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkGenerateDicewareLowAlloc(b *testing.B) {
    opts := DicewareOptions{
        WordCount: 6,
        Separator: "-",
        Capitalize: true,
        AddNumbers: true,
    }
    
    var builder strings.Builder
    
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        err := DicewareLowAlloc(opts, &builder)
        if err != nil {
            b.Fatal(err)
        }
        builder.Reset()
    }
}

func BenchmarkGenerateCsprng(b *testing.B) {
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        _, err := Csprng(32)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkGenerateCsprngLowAlloc(b *testing.B) {
    var builder strings.Builder
    
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        err := CsprngLowAlloc(32, &builder)
        if err != nil {
            b.Fatal(err)
        }
        builder.Reset()
    }
}


// Parallel benchmarks
func BenchmarkGenerateDicewareParallel(b *testing.B) {
    opts := DicewareOptions{WordCount: 6, Separator: "-"}
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, err := Diceware(opts)
            if err != nil {
                b.Fatal(err)
            }
        }
    })
}

func BenchmarkGenerateDicewareLowAllocParallel(b *testing.B) {
    opts := DicewareOptions{WordCount: 6, Separator: "-"}
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        var builder strings.Builder
        for pb.Next() {
            err := DicewareLowAlloc(opts, &builder)
            if err != nil {
                b.Fatal(err)
            }
            builder.Reset()
        }
    })
}