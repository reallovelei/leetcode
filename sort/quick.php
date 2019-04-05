<?php
//function findmid($start, $end)

function quicksort(&$arr, $start, $end) {
    if ($start > $end) return ;
    $mid = intval(($start + $end)/2);
    $i = $start;
    $j = $end;
    $mv = $arr[$mid]; // 注意 这个值需要 先定下来，不能在下面的判断条件里写 > $arr[$mid]  因为在midindex位置上的值 可能会被换掉。

    while($i < $j) {
        echo "start:{$start} {$arr[$start]} mid:{$mid} {$mv} end:{$end} {$arr[$end]}\n";
        while($i < $j && $arr[$i] < $mv) {
            $i++;
        }

        while($i < $j && $arr[$j] > $mv) {
            $j--;
        }

        //if ($arr[$i] > $arr[$j]) {
            echo "index:{$i} {$j} val:{$arr[$i]} {$arr[$j]}  \n";
            $tmp = $arr[$i];
            $arr[$i] = $arr[$j];
            $arr[$j] = $tmp;
            echo implode(',', $arr). "\n\n";
        //}
    }
    //echo "\n i:".$i." j:".$j."  s:".$start." e:".$end;
    //var_dump($arr);

    if ($i -1 > $start)
    quicksort($arr, $start, $i -1);
    if ($end > $i+1)
    quicksort($arr, $i + 1, $end);
}
$arr = [2,3,9,5,7,4,1,10,6,8];
$arr = quicksort($arr, 0, count($arr) -1);
var_dump($arr);
die;

