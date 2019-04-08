#include <vector>

using namespace std;

int removeDuplicates(vector<int>& nums) 
{
    int numErase = 0;
    int i;
    int j;
    int temp;
    
    for(i = 0,
        j = 1; 
        j < nums.size(); 
        ++i,
        ++j)
    {
        if(nums[i] == nums[j])
        {
            temp = nums[i];
            nums[i] = nums[numErase];
            nums[numErase] = temp;
            ++numErase;
        }
    }
    
    nums.erase(nums.begin(), nums.begin() + numErase);
    
    std::sort(nums.begin(), nums.end());
    
    return nums.size();
}
