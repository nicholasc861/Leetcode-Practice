package main

func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        difference := target - nums[i]
        
        value, ok := numbers[difference]
        if ok {
            return []int{value, i} 
        }
        numbers[nums[i]] = i
    }
     return nil
    
}