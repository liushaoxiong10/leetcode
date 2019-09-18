

class Solution:
    def twoSum(self, nums, target):
        for ii, i in enumerate(nums):
            for jj, j in enumerate(nums[ii+1:]):
                if (i + j) == target:
                    return [ii, jj + ii + 1]


if __name__ == "__main__":
    testExample = [0, 1, 2, 3]
    testResult = [1, 3]
    t1 = Solution().twoSum(testExample, 4)
    print(testResult == t1)
    print(t1)
