package main

import ("strings" )

func simplifyPath(path string) string {
    string_parts := strings.Split(path, "/")
    
    canonical := "/"
    newString := make([]string, 0)
    
    for i := 0; i < len(string_parts); i++ {
        switch string_parts[i] {
            case "/":
                continue
            case "..":
                if len(newString) > 0{
                    newString = newString[:len(newString)-1]
                }
            case ".":
                continue
			default:
				if len(string_parts[i]) > 0 {
					newString = append(newString, string_parts[i])
				}
        }
    }
  
	result := strings.Join(newString, "/")
    canonical += result
   
    
    return canonical
}
    
