/*
 * @lc app=leetcode.cn id=42 lang=c
 *
 * [42] 接雨水
 */

// @lc code=start


int trap(int* height, int heightSize){
    int *left = &height[0];
    int *right = &height[heightSize - 1];

    int *cl = left + 1; // 当前左指针
    int *cr = right - 1; // 当前右指针
    
    int waterCnt = 0;   // 结果
    while (cr >= left && cl <= right) {
        // 如果最右边较矮，将当前右指针向左移动(其实这里就是看先移动 比较左边的还是先比较右边的)
        if (*right < *left) {
            // 如果当前右指针cr 比 最右指针还高的时候代表他们之间可以形成洼地。
            if (*cr >= *right) {
                // （*right） 高度   *  宽度 (right - cr - 1)
                waterCnt += *right * (right - cr - 1);
                // 将最右边界 设置到当前位置。
                right = cr;
            } else {
                // 减去当前自身所占用的面积。
                waterCnt -= *cr;
            }
            cr--;
        } else { // 左边较矮,将当前左指针右移动

            if (*cl >= *left) {
                waterCnt += *left * (cl - left - 1);
                left = cl;
            } else {
                waterCnt -= *cl;
            }
            cl++;
        }
    }
    return waterCnt;
}


 
// @lc code=end

