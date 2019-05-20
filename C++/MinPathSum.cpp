class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) 
    {
        if(grid.size() == 0) return 0;
        int left;
        int up;
        for(int ver = 0; ver < grid.size(); ++ver)
        {
            for(int hor = 0; hor < grid[ver].size(); ++hor)
            {
                if(ver == 0 && hor == 0) continue;
                left = hor == 0 ? INT_MAX : grid[ver][hor - 1];
                up = ver == 0 ? INT_MAX : grid[ver - 1][hor];
                grid[ver][hor] += left < up ? left : up;
            }
        }
        
        return grid[grid.size() - 1][grid[0].size() - 1];
    }
};