package mytesting

// import "fmt"

// func HomeworkBE4905186(input string) (bool, error) {
// 	countChars := len(input)
// 	if countChars < 3 {
// 		return false, fmt.Errorf("total character cant be less than 3")
// 	}

// 	if countChars > 15 {
// 		return false, fmt.Errorf("total character cant be greater than 15")
// 	}

//     for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
//         if input[i] != input[j] {
//             return false, nil
//         }
//     }
//     return true, nil
// }