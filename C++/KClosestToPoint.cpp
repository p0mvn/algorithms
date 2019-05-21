class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int K) 
    {
        multimap<int, int> map;
        for(int i = 0; i < points.size(); ++i)
        {
            map.insert(std::pair<int, int>(
                points[i][0] * points[i][0] + points[i][1]*points[i][1],
                i));
        }
        
        vector<vector<int>> result;
        
        for(multimap<int, int>::iterator itr = map.begin();
            K > 0;
           ++itr, K--)
        {
            result.push_back(points[itr->second]);
        }
        
        return result;
        
    }
};