func canCompleteCircuit(gas []int, cost []int) int {
    var totalGas int
    var totalCost int
    
    for i := 0; i < len(gas); i++ {
        totalGas += gas[i]
        totalCost += cost[i]
    }
    
    if totalGas < totalCost {
        return -1
    }
    
    var saveIndex int
    var currentGas int
    
     for i := 0; i < len(gas); i++ {
		 currentGas = 0
		 works := true
         if gas[i] >= cost[i] {
			 temp := i
             for j := i; j < len(gas); j++ {
                 if (currentGas + gas[j] - cost[j]) >= 0 {
					currentGas += (gas[j] - cost[j])
                 } else {
                     currentGas = -1
                     works = false
                     break
                 }
             }
             if works && temp != 0{
                 for k := 0; k < saveIndex; k++ {
                    if (currentGas + gas[k] - cost[k]) >= 0 {
                        currentGas += (gas[k] - cost[k])
                    } else {
                        currentGas = -1
                        break
                    }
                }
			 }
			 if currentGas >= 0 {
				 return temp
			 }
         }
    }
    return -1
}
