package leetcode

import (
	"reflect"
	"testing"
)

// 799.香槟塔
// https://leetcode-cn.com/problems/champagne-tower
func champagneTower(poured int, query_row int, query_glass int) float64 {
	var dp [][]float64
	dp[0][0] = float64(poured)

	for i := 1; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			// 如果是第一列 只要它正上方的那一杯的 1/2的流量
			if j == 0 {
				dp[i][j] = float64(max(0.0, dp[i-1][j]-1)/2)
				continue
			}
			// 最后一列
			if j == i {
				dp[i][j] = float64(max(0.0, dp[i-1][j-1]-1)/2)
				continue
			}

			// 在中间的列
			dp[i][j] = float64(max(0.0, dp[i-1][j]-1)/2 + max(0.0, dp[i-1][j-1]-1)/2)
		}
	}
	return dp[query_row][query_glass]
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want 
	}{
		{
            want: ,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := 
			if !reflect.DeepEqual(tC.want,get){
				t.Errorf("input: %+v get: %v\n",tC,get)
			}
		})
	}
}
