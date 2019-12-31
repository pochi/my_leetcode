package leetcode006

//import "fmt"

func convert(s string, numRows int) string {
	runes := []rune(s)
	zigzagMap := make(map[int]string, numRows)

	rowMargin := false
	marginCount := 0
	marginLength := numRows - 2
	margin := 1
	
	for i, c := range runes {
		if rowMargin == true {
			zigzagMap[numRows - margin - 1] = zigzagMap[numRows - margin - 1] + string(c)
			if margin == marginLength {
				margin = 1
				marginCount = 0
				rowMargin = false
				continue
			} else {
				margin += 1
			}
		} else {
			zigzagMap[marginCount % numRows] = zigzagMap[marginCount % numRows] + string(c)
		}
		
		marginCount += 1				

		if i != 0 && marginCount % numRows == 0 {
			rowMargin = true
		}
	}

	zigzagString := ""
	for _, value := range zigzagMap {
		zigzagString = zigzagString + value
	}
	
	return zigzagString
}
