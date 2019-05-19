class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) 
    {
        unordered_map<string, std::vector<string> > tempResult;
        string cur;
        string curKey;
        for(int i = 0; i < strs.size(); ++i)
        {
            //calculate code
            cur = strs[i];
            
            sort(cur.begin(), cur.end());
            
            tempResult[cur].push_back(strs[i]);
        }
        
        vector<vector<string>> result;
        for(unordered_map<string, std::vector<string> >::iterator itr = tempResult.begin();
           itr != tempResult.end();
           ++itr)
        {
            result.push_back(itr->second);
        }
        
        return result;
    }
};