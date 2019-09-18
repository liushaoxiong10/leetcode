## 题目

[两数之和](https://leetcode-cn.com/problems/two-sum)

难度：**简单**

给定一个整数数组 **nums** 和一个目标值 **target** ，请你在该数组中找出和为目标值的那 **两个** 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

## 示例输出

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

## 解析

我们最简单的思路就是两层嵌套循环，遍历每个元素，寻找它之后的元素是否有与它的和等于 targe 的，这种暴力解析的复杂度为 O(n^2)

伪代码如下：

```
for i=0; i<length(nums); i++:
    for j=i+1 ; j < length(nums; j++):
       if nums[i] + nums[j] == target:
           return [i,j]
```

第二种思路是使用一个 hash 表，记录没有被匹配的元素复杂度为 O(n)

伪代码如下：

```
   hashmap type {int:int}
   for i=0;i<length(nums);i++{
       result = target - nums[i]
       if result in hashmap{
           return [hashmap[result], i]
       }
       hashmap[nums[i]] = i
   }
```
