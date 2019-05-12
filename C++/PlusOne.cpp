    vector<int> plusOne(vector<int>& digits) 
    {
        int curN;
        int carry = 1;
        int cur;
        
        for(cur = digits.size() - 1; 
           cur >= 0;
           --cur)
        {
            curN = digits[cur] + carry;
            carry = 0;
            
            digits[cur] = curN % 10;
            
            if(curN < 10)
            {
                break;
            } else 
            {
                carry = 1;
            }
        }
        
        if(1 == carry && digits.size() > 0)
        {
            carry = digits[0];
            digits[0] = 1;
            for(cur = 1; cur < digits.size(); ++cur)
            {
                curN = digits[cur];
                digits[cur] = carry;
                carry = curN;
            }
            digits.push_back(carry);
        }
        
        return digits;
    }