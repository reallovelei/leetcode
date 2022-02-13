/*
 * @lc app=leetcode.cn id=70 lang=c
 *
 * [70] 爬楼梯
 */


int climbStairs(int n){
    int dp[n+3];
    dp[1] = 1;
    dp[2] = 2;
    for (int i = 3; i <= n; i++) {
        dp[i] = dp[i-2] + dp[i-1];
    }
    return dp[n];
}



