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
        lengthOfString := len(characters)
        
        if (lengthOfString % 2) == 0 {
            lengthOfString--
        }
        
        for i := 0; i <= lengthOfString/2;{
            for j := len(characters)-1; j >= lengthOfString/2; j-- {
                if characters[j] != characters[i] {
                    return false
			    }
			    i++
            }
        }  
    }
    return true
}
