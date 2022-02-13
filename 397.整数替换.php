/*
 * @lc app=leetcode.cn id=397 lang=php
 *
 * [397] 整数替换
 */
class Solution {

    public $map;
    /**
     * @param Integer $n
     * @return Integer
     */
    function integerReplacement($n) {
        /*
        if (isset($this->map[$n])) {
            return $this->map[$n] + 1;
        } else {
            $this->map[$n] = 
        }
        */
        if ($n > 1) {
                if ($n % 2 == 0) {
                    return $this->integerReplacement($n/2) + 1;
                } else {
                    return min($this->integerReplacement($n +1) +1, $this->integerReplacement($n-1) + 1);
                }
        } else {
            return 0;
        }
    }
}

