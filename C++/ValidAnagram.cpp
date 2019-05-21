class Solution {
public:
    bool isAnagram(string s, string t) 
    {
        if(s.size() != t.size()) return false;
        
        std::vector<int> alphabet;
        alphabet.resize(26);
        
        for(int i = 0; i < s.size(); ++i)
        {
            ++alphabet[static_cast<int>(s[i]) - 97];
            --alphabet[static_cast<int>(t[i]) - 97];
        }
        
        for(vector<int>::iterator itr = alphabet.begin(); 
            itr != alphabet.end(); ++itr)
        {
            if(*itr != 0)
            {
                return false;
            }
        }
        
        return true;
    }
};