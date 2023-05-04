package mytesting

// import "fmt"


// func HomeworkBE4920086(numbers []int) (sorted []int, err error) {
// 	if len(numbers) < 5 {
// 		return sorted, fmt.Errorf("cant process if the input less than 5 numbers")
// 	}

// 	if len(numbers) > 30 {
// 		return sorted, fmt.Errorf("cant process if the input greater than 30 numbers")
// 	}

//     for i := 0; i < len(numbers)-1; i++ {
//         for j := i + 1; j < len(numbers); j++ {
//             if numbers[i] > numbers[j] {
//                 numbers[i], numbers[j] = numbers[j], numbers[i]
//             }
//         }
//     }

//     return numbers, nil
// }