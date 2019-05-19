class Solution {
public:
    string reverseStr(string s, int k) 
    {
        if(s.size() < 2)
        {
            return s;
        }
        
        char temp;
        
        int curStart = 0;
        int curEnd = k - 1;
        int startTemp;
        int endTemp;
        
        if (curEnd >= s.size() && curStart < s.size())
        {
           curEnd = s.size() - 1;
        }
        
        while(curEnd < s.size())
        {
            startTemp = curStart;
            endTemp = curEnd;
            
            while(startTemp < endTemp)
            {
                temp = s[startTemp];
                s[startTemp++] = s[endTemp];
                s[endTemp--] = temp;
            }
            
            curStart += 2 * k;
            curEnd += 2 * k;
            
            if (curEnd >= s.size() && curStart < s.size())
            {
                curEnd = s.size() - 1;
            }   
        }
        
        return s;
    }
};