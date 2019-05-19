/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* swapPairs(ListNode* head) 
    {
        return head != NULL ? swap(head, head->next) : NULL;
    }
    
    ListNode* swap(ListNode* first, ListNode* second)
    {
        if(second != NULL)
        {
            first->next = second->next != NULL ? swap(second->next, second->next->next) : NULL;
            second->next = first;
    
            return second;        
        } 
        
        return first;
    }
};