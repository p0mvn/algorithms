class Solution {
public:
    double myPow(double x, int n) 
    {        
        if(n == 0 || x == 1)
        {
            return 1;
        } else if (n == INT_MIN)
        {
            return x == -1;
        }
        else if(n < 0)
        {
            n = -n;
            x = 1 / x;
        }
        
        return (n % 2 == 0) ? myPow(x * x, n / 2) : x * myPow(x * x, n / 2);   
    }
};