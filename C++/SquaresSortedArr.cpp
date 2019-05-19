class Solution {
public:
    vector<int> sortedSquares(vector<int>& A) 
    {
        vector<int> result;
        result.resize(A.size());
        int start = 0;
        int end = A.size() - 1;
        int cur = A.size();
        
        while(start < end)
        {
            if(A[end] * A[end] > A[start] * A[start])
            {
                result[--cur] = A[end] * A[end];
                --end; 
            } else 
            {
                result[--cur] = A[start] * A[start];
                ++start;
            }   
        }
        
        if (cur != 0 && start == end)
        {
            result[--cur] = A[start] * A[start];
        }
        
        return result;
    }
};