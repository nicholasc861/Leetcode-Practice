import "strings"

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    string_parts := strings.Split(s, " ")
    newString := make([]string, 0)

    for i := len(string_parts)-1; i >= 0; i-- {
        if strings.TrimSpace(string_parts[i]) != ""{
            newString = append(newString, string_parts[i])
        } 
    }
    
    newS := strings.Join(newString, " ")

    return newS
}