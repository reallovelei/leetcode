#include <stdio.h>

int majorityElement(int* nums, int numsSize){
    int count = 0;
    int rs = 0;
    for (int i = 0; i < numsSize; i++) {
        if (rs == nums[i]) {
            count++;
        } else 
        {
            if (count > 0){
                count--;
            } else {
                rs = nums[i];
                count = 1;
            }
        } 
    }
    return rs;
}

int main() {
    int rs = 0;
    int arr[7] = {2,2,1,1,1,2,2};
    rs = majorityElement(arr, 7);
    printf("rs is:%d \n", rs);
}
