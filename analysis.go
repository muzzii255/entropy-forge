package entropyforge
import(
	"strings"
	
	"math"
		"fmt"
		)

type ComplexityScore struct {
    Score           int      `json:"score"`            // 0-100
    Entropy         float64  `json:"entropy_bits"`
    Strength        string   `json:"strength"`         // "Very Weak", "Weak", "Fair", "Good", "Strong", "Very Strong"
    CrackTime       string   `json:"crack_time"`       // "2.3 years", "centuries"
    Weaknesses      []string `json:"weaknesses"`
    Suggestions     []string `json:"suggestions"`
    CharacterTypes  int      `json:"character_types"`  // Number of different char types used
    PatternScore    int      `json:"pattern_score"`    // Detection of common patterns
}

type PasswordAnalysis struct {
    Password   string          `json:"password"`
    Length     int             `json:"length"`
    Complexity ComplexityScore `json:"complexity"`
    Generated  string          `json:"generated_at"`
}


func AnalyzePassword(password string) ComplexityScore {
    score := ComplexityScore{
        Weaknesses:  []string{},
        Suggestions: []string{},
    }

    score.Entropy = calculateEntropy(password)
    score.CharacterTypes = countCharacterTypes(password)
    score.PatternScore = analyzePatterns(password)

    score.Score = calculateOverallScore(password, score.Entropy, score.CharacterTypes, score.PatternScore)

    score.Strength = getStrengthLevel(score.Score)

    score.CrackTime = calculateCrackTime(score.Entropy)

    score.Weaknesses, score.Suggestions = analyzeWeaknesses(password, score)

    return score
}

func calculateEntropy(password string) float64 {
    if len(password) == 0 {
        return 0
    }

    freq := make(map[rune]int)
    for _, char := range password {
        freq[char]++
    }

    var entropy float64
    length := float64(len(password))

    for _, count := range freq {
        p := float64(count) / length
        if p > 0 {
            entropy -= p * math.Log2(p)
        }
    }

    return entropy * length
}

func countCharacterTypes(password string) int {
    var hasLower, hasUpper, hasDigit, hasSymbol bool

    for _, char := range password {
        switch {
        case char >= 'a' && char <= 'z':
            hasLower = true
        case char >= 'A' && char <= 'Z':
            hasUpper = true
        case char >= '0' && char <= '9':
            hasDigit = true
        default:
            hasSymbol = true
        }
    }

    count := 0
    if hasLower { count++ }
    if hasUpper { count++ }
    if hasDigit { count++ }
    if hasSymbol { count++ }

    return count
}

func analyzePatterns(password string) int {
    score := 100

    patterns := []string{
        "123", "abc", "qwe", "asd", "zxc",
        "password", "admin", "user", "test",
        "000", "111", "222", "999",
    }

    lowerPass := strings.ToLower(password)
    for _, pattern := range patterns {
        if strings.Contains(lowerPass, pattern) {
            score -= 20
        }
    }

    if containsKeyboardPattern(password) {
        score -= 15
    }

    if hasRepeatedChars(password) {
        score -= 10
    }

    if hasSequentialChars(password) {
        score -= 10
    }

    if score < 0 {
        score = 0
    }

    return score
}
func clamp(value, min, max int) int {
    if value < min { return min }
    if value > max { return max }
    return value
}

func calculateOverallScore(password string, entropy float64, charTypes, patternScore int) int {
    entropyScore := int(math.Min(entropy*2, 50))

    lengthScore := int(math.Min(float64(len(password)), 20))

    typeScore := charTypes * 5

    patternPenalty := (100 - patternScore) / 10

    totalScore := entropyScore + lengthScore + typeScore - patternPenalty

    totalScore = clamp(totalScore,0,100)

    return totalScore
}

func getStrengthLevel(score int) string {
    switch {
    case score >= 90:
        return "Very Strong"
    case score >= 75:
        return "Strong"
    case score >= 60:
        return "Good"
    case score >= 40:
        return "Fair"
    case score >= 20:
        return "Weak"
    default:
        return "Very Weak"
    }
}

func calculateCrackTime(entropy float64) string {
    guessesPerSecond := 1e9

    attempts := math.Pow(2, entropy) / 2
    seconds := attempts / guessesPerSecond

    return formatTime(seconds)
}

func formatTime(seconds float64) string {
    if seconds < 1 {
        return "Instant"
    } else if seconds < 60 {
        return fmt.Sprintf("%.0f seconds", seconds)
    } else if seconds < 3600 {
        return fmt.Sprintf("%.0f minutes", seconds/60)
    } else if seconds < 86400 {
        return fmt.Sprintf("%.0f hours", seconds/3600)
    } else if seconds < 31536000 {
        return fmt.Sprintf("%.0f days", seconds/86400)
    } else if seconds < 31536000000 {
        return fmt.Sprintf("%.1f years", seconds/31536000)
    } else {
        return "Centuries"
    }
}

func analyzeWeaknesses(password string, score ComplexityScore) ([]string, []string) {
    var weaknesses, suggestions []string

    if len(password) < 8 {
        weaknesses = append(weaknesses, "Password too short")
        suggestions = append(suggestions, "Use at least 12 characters")
    }

    if score.CharacterTypes < 3 {
        weaknesses = append(weaknesses, "Limited character variety")
        suggestions = append(suggestions, "Include uppercase, lowercase, numbers, and symbols")
    }

    if score.Entropy < 50 {
        weaknesses = append(weaknesses, "Low entropy")
        suggestions = append(suggestions, "Use more random characters or longer passphrase")
    }

    if score.PatternScore < 80 {
        weaknesses = append(weaknesses, "Contains common patterns")
        suggestions = append(suggestions, "Avoid dictionary words and common patterns")
    }

    if hasRepeatedChars(password) {
        weaknesses = append(weaknesses, "Contains repeated characters")
        suggestions = append(suggestions, "Reduce character repetition")
    }

    return weaknesses, suggestions
}

func containsKeyboardPattern(password string) bool {
    patterns := []string{"qwerty", "asdf", "zxcv", "1234", "abcd"}
    lower := strings.ToLower(password)
    for _, pattern := range patterns {
        if strings.Contains(lower, pattern) {
            return true
        }
    }
    return false
}

func hasRepeatedChars(password string) bool {
    for i := range len(password)-2 {
        if password[i] == password[i+1] && password[i+1] == password[i+2] {
            return true
        }
    }
    return false
}

func hasSequentialChars(password string) bool {
    for i := range len(password)-2 {
        if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
            return true
        }
    }
    return false
}
