package entropyforge


import (
	"crypto/rand"
	"time"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"

)


var (
    stringBuilderPool = sync.Pool{
        New: func() interface{} {
            return &strings.Builder{}
        },
    }
    
    byteSlicePool = sync.Pool{
        New: func() interface{} {
            slice := make([]byte, 64) 
            return &slice
        },
    }
)


type DicewareOptions struct {
    WordCount    int
    Separator    string // "-", " ", "_"
    Uppercase    bool   // Make one random word ALL CAPS
    Capitalize   bool   // Capitalize first letter of random word  
    AddNumbers   bool   // Add numbers at end
    AddSymbols   bool   // Add symbols
}





func CsprngLowAlloc(length int, result *strings.Builder) error {
	result.Reset()
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*|/"
	for range length {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return err
		}else {
			result.WriteByte(charset[randomIndex.Int64()])
		}
	}

	return nil
	
}



func DicewareLowAlloc(opts DicewareOptions, result *strings.Builder) error {
    result.Reset() 
    
    special := "!@#$%^&*/|"
    var uppercaseWordIndex int = -1
    if opts.Uppercase {
        if idx, err := rand.Int(rand.Reader, big.NewInt(int64(opts.WordCount))); err != nil {
            return err
        } else {
            uppercaseWordIndex = int(idx.Int64())
        }
    }
    
    var specialIndex int = -1
    var specialCh byte 
    if opts.AddSymbols {
   		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(special))))
        if err != nil {
            return err
        }
        specialCh =  special[idx.Int64()]
        if idx, err := rand.Int(rand.Reader, big.NewInt(int64(opts.WordCount))); err != nil {
            return err
        } else {
            specialIndex = int(idx.Int64())
        }
    }
    var numIndex int = -1
    var numVal string
    if opts.AddNumbers {
	    if idx, err := rand.Int(rand.Reader, big.NewInt(int64(opts.WordCount))); err != nil {
	        return err
	    } else {
	        numIndex = int(idx.Int64())
	    }
		num, err := rand.Int(rand.Reader,big.NewInt(int64(696969)))
		if err != nil {
			return err
		}else{
			numVal = strconv.Itoa(int(num.Int64()))
		}
    }
    
    
    
    for i := range opts.WordCount {
        diceRoll, err := rollDiceLowAlloc()
        if err != nil {
            return err
        }
        if word, exists := wordList[diceRoll]; exists {
         	if opts.Capitalize {
          		result.WriteString(caser.String(word))
          	}else{
           		result.WriteString(word)
           }
        	if i > 0 && opts.Separator != "" {
                result.WriteString(opts.Separator)
            }
            if i == uppercaseWordIndex {
            	result.WriteString(strings.ToUpper(word))
            }
            if i == specialIndex   {
            	result.WriteByte(specialCh)            	
            }
            if i == numIndex{
            	result.WriteString(numVal)
            }
            // fmt.Println(result.String())
        }
    }
    
    return nil
}


func rollDiceLowAlloc() (string, error) {
    var diceRolls [5]byte
    
    for i := range 5 {
        rd, err := rand.Int(rand.Reader, big.NewInt(6))
        if err != nil {
            return "", err
        }
        diceRolls[i] = byte(rd.Int64() + 1 + '0') 
    }
    
    return string(diceRolls[:]), nil
}


func rollDice() (string, error) {
    var idx string
    for range 5 {
        rd, err := rand.Int(rand.Reader, big.NewInt(6))
        if err != nil {
            return "", err
        }
        diceValue := rd.Int64() + 1
        idx += strconv.FormatInt(diceValue, 10)
    }
    return idx, nil
}


func Csprng(length int) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*|/"
	password := make([]byte, length)

	for i := range length {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func applyTransformations(words []string,opts DicewareOptions)([]string,error){
	special := "!@#$%^&*/|"
	if opts.Uppercase {
        randomWordIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
        if err != nil {
            return []string{}, err
        }
        words[randomWordIndex.Int64()] = strings.ToUpper(words[randomWordIndex.Int64()])
    }
	if opts.AddSymbols{
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(special))))
        if err != nil {
            return []string{}, err
        }
        words = append(words, string(special[idx.Int64()]))
	}
	if opts.AddNumbers{
		idx, err := rand.Int(rand.Reader,big.NewInt(int64(696969)))
		if err != nil{
			return []string{},err
		}
		words = append(words,strconv.Itoa(int(idx.Int64())))
	}
    
	if !opts.Uppercase && opts.Capitalize {
		for i := range words {
			a := words[i]
			a = caser.String(a)
			words[i] = a
		}
		
	}
	
	return words,nil
}

func Diceware(opts DicewareOptions) (string, error) {
    
    words := []string{}
    
    for  range opts.WordCount {
        diceRoll, err := rollDice()
        if err != nil {
            return "", err
        }
        
        word, exists := wordList[diceRoll]
        if !exists {
            return "", fmt.Errorf("word not found for dice roll: %s", diceRoll)
        }
        // words[i] = word
        words = append(words,word)
    }
    
    words, err := applyTransformations(words, opts)
    if err != nil {
        return "", err
    }
    
    return strings.Join(words, opts.Separator), nil
}

func CsprngAnalysis(length int, result *strings.Builder) (*PasswordAnalysis, error) {
	err := CsprngLowAlloc(length, result)
	if err != nil {
		return nil ,err
	}
	password := result.String()
	complexity := AnalyzePassword(password)
	return &PasswordAnalysis{
        Password:   password,
        Length:     len(password),
        Complexity: complexity,
        Generated:  time.Now().UTC().Format(time.RFC3339),
    }, nil
	
}

func DicewareAnalysis(opts DicewareOptions, result *strings.Builder) (*PasswordAnalysis, error) {
    err := DicewareLowAlloc(opts, result)
    if err != nil {
        return nil, err
    }
    
    password := result.String()
    complexity := AnalyzePassword(password)
    
    return &PasswordAnalysis{
        Password:   password,
        Length:     len(password),
        Complexity: complexity,
        Generated:  time.Now().UTC().Format(time.RFC3339),
    }, nil
}
