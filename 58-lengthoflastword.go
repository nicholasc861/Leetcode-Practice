import ("strings")

func lengthOfLastWord(s string) int {
    string_parts := strings.Split(s, " ")
    
    for i := (len(string_parts)-1); i >= 0; i-- {
        if string_parts[i] != ""{
            return len(string_parts[i])
        }
    } 
    return 0
}