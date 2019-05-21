class Solution {
public:
    int fib(int N) 
    {
        if(N == 0) return 0;
        if(N == 1) return 1;
        
        std::vector<int> nums;
        nums.resize(N);
        int nMinTwo = 0;
        int nMinOne = 1;
        nums[nMinTwo] = 0;
        nums[nMinOne] = 1;
        
        while(nMinOne + 1 != N)
        {
            nums[nMinOne + 1] = nums[nMinOne] + nums[nMinTwo];
            ++nMinOne;
            ++nMinTwo;
        }
        
        return nums[nMinOne] + nums[nMinTwo];
    }
};