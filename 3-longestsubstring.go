import "strings"

func lengthOfLongestSubstring(s string) int {
    characters := strings.Split(s, "")
    length := 0
    temp := 0
    substring := ""
    
    for i := 0; i < len(characters)-1; i++ {
        if strings.ContainsAny(substring, characters[i]){
            temp = 0
            substring = ""
        } 
        
        temp++
  
        if temp > length {
            length = temp
        }
            
        substring += characters[i]
    }
    
    return length
}