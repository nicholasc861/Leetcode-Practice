func divide(dividend int, divisor int) int {
    temp := divisor
    temp2 := dividend
    var nDividend bool
    var nDivisor bool
    var result int
    
    if divisor < 0 {
        divisor = temp - temp - temp
        temp = divisor
        nDivisor = true
    } 
    
    if dividend < 0 {
        dividend = temp2 - temp2 - temp2
        temp2 = dividend
        nDividend = true
    }
    
    for ; temp <= dividend; temp += divisor{
        result++
    }
    
    
    if nDivisor && nDividend {
        result = result
    } else if nDividend || nDivisor {
        result = result - result - result
    }  
    
    if result < -2147483648 || result > 2147483647 {
        return 2147483647
    }
    
    
    return result
}