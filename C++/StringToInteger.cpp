class Solution {
public:
    //"9223372036854775808"
    // "0000000000854775808"
    int myAtoi(string str) 
    {
        int flag = 1;
        int start = 0;
        while(start < str.size() && str[start] == ' ')
        {
            ++start;
        }
        
        if(str[start] == '-')
        {
            flag = -1;
            ++start;
        } else if (str[start] == '+')
        {
            ++start;
        }
        
        int end = start;
        long result = 0;
        
        while(start < str.size() && str[start] - 48 >= 0 && str[start] - 48 <= 9)
        {
            result = result * 10 + (str[start] - 48);
            if(result * flag > INT_MAX || result * flag < INT_MIN)
            {
                return flag > 0 ? INT_MAX : INT_MIN;   
            }
            
            ++start;
        }
               
        return static_cast<int>(result * flag);
    }
};