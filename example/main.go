package main

import (
    "fmt"
    "log"
    "strings"
    "github.com/muzzii255/entropy-forge"
)

func main() {
    opts := entropyforge.DicewareOptions{
        WordCount:  6,
        Separator:  "-",
        Capitalize: true,
        AddNumbers: true,
        AddSymbols: true,
        Uppercase: true,
    }
    
    password, err := entropyforge.Diceware(opts)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Diceware Password: %s\n", password)
    
    // Analysis example
    analysis := entropyforge.AnalyzePassword(password)
    fmt.Printf("Strength: %s (%d/100)\n", analysis.Strength, analysis.Score)
    fmt.Printf("Entropy: %.1f bits\n", analysis.Entropy)
    
    // CSPRNG example
    csprng, err := entropyforge.Csprng(24)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CSPRNG Password: %s\n", csprng)

    // Low Allocation example
    result := &strings.Builder{}
    entropyforge.CsprngLowAlloc(20, result)
    fmt.Printf("CSPRNG Low Allocation : %s\n",result.String())
    
    entropyforge.DicewareLowAlloc(opts, result)
    fmt.Printf("CSPRNG Low Allocation : %s\n",result.String())

}