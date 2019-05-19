class Solution {
public:
    void reverseString(vector<char>& s) 
    {
        char temp;
        int end = s.size() - 1;
        int start = 0;
        while(start < end)
        {
            temp = s[start];
            s[start] = s[end];
            s[end] = temp;
            ++start;
            --end;
        }
    }
};