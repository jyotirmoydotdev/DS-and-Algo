func diagonalDifference(arr [][]int32) int32 { 
    // Write your code here
    var SumPrimaryDiagonal int32
    var SumSecondaryDiagonal int32
    for i:=0;i<len(arr);i++{
        for j:=0; j<len(arr);j++{
            if i == j{
				SumPrimaryDiagonal += arr[i][j]
			}
			if i+j == len(arr)-1{
				fmt.Printf("%v ",arr[i][j])
				SumSecondaryDiagonal += arr[i][j]
			}
        }
    }
	var diff int32 = SumPrimaryDiagonal - SumSecondaryDiagonal
	if diff<0{
		diff = diff * -1
		return diff
	} else{
		return diff
	}
}