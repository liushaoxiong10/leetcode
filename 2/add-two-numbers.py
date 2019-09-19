# Definition for singly-linked list.


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

    def AddNext(self, x):
        node = ListNode(x)
        self.next = node
        return self.next


def CreaterNewListNode(vals):
    if len(vals) == 0:
        return listNode(0)
    node = ListNode(0)
    lastnode = node
    for i in vals:
        lastnode = lastnode.AddNext(i)
    return node.next


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        node = ListNode(0)
        lastnode = node
        add = 0
        while True:
            sum_number = add
            add = 0
            if l1:
                sum_number += l1.val
                l1 = l1.next
            if l2:
                sum_number += l2.val
                l2 = l2.next
            if sum_number >= 10:
                sum_number -= 10
                add += 1
            lastnode.next = ListNode(sum_number)
            lastnode = lastnode.next
            if not l1 and not l2 and add == 0:
                break
        return node.next


if __name__ == "__main__":
    print("begin")
    l1 = CreaterNewListNode([1, 2, 3])
    l2 = CreaterNewListNode([9, 7])
    s = Solution()
    l3 = s.addTwoNumbers(l1, l2)
    while l3:
        print(l3.val)
        l3 = l3.next
