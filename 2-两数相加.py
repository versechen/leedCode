# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
node_container = []

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        flag = True
        add_1 = []
        add_2 = []
        node = l1
        while flag:
            add_1.append(node.val)
            if node.next:
                node = node.next
            else:
                flag = False
                
        flag = True
        node = l2
        while flag:
            add_2.append(node.val)
            if node.next:
                node = node.next
            else:
                flag = False

        add_1 = add_1[::-1]
        add_2 = add_2[::-1]
        total = []
        add_1_len = len(add_1)
        add_2_len = len(add_2)
        max_len = max(add_1_len, add_2_len)
        if add_1_len < add_2_len:
            temp =  add_1
            add_1 = add_2
            add_2 = temp

        for index, _ in enumerate(add_1):
            if index + 1 <= add_2_len:
                total.append(add_1[index] + add_2[index])
            else:
                total.append(add_1[index])
        #进位
        for index in range(0, max_len):
            num = total[index]
            a = num % 10
            b = num / 10
            if b !=0:
                total[index] = a
                if index + 1 < max_len:
                    total[index+1] = total[index+1] + 1
                else:
                    total.append(1)
        
        # 输出序列
        revere_node = total[::-1]
        for index, num in enumerate(revere_node):
                node_container.append(ListNode(num))

        # 赋值next
        num_len = len(revere_node)
        for index, _ in enumerate(node_container):
            if index+1< num_len:
                node_container[index].next = node_container[index+1]

        return node_container[0]

