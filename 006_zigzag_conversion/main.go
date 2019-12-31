package leetcode006

//import "fmt"

// ABCD,2 -> ACBD
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	
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

		if i != 0 && numRows > 2 && marginCount % numRows == 0 {
			rowMargin = true
		}
	}

	zigzagString := ""
	for i:=0; i<numRows; i++ {
		zigzagString = zigzagString + zigzagMap[i]
	}
	
	return zigzagString
}
