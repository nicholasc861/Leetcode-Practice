func reverse(x int) int {
    var reverse int
    
    negative := false
    
    if x < 0 {
        x *= -1
        negative = true
    }

    for ; x != 0; x /= 10 {
        digit := (x % 10)
        temp := reverse * 10 + digit
        reverse = temp
    }
    
    if negative {
        reverse *= -1
    }
    
     if reverse < -2147483647 || reverse > 2147483647 {
        return 0
    }
        
    return reverse
    
}