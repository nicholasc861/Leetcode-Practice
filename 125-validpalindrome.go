import (
    "strings"
	"regexp"
)

func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {

	}
    s = reg.ReplaceAllString(s, "")
    
    if len(strings.TrimSpace(s)) != 0 {
        s = strings.ToLower(s)
        characters := strings.Split(s, "")
        end := len(characters)-1
        begin := 0
        
        for begin < end {
            if characters[begin] != characters[end] {
                return false
            }
            begin++
            end--
        }
    }
    return true
}
