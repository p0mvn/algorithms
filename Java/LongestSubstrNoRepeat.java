

class LongestSubstrNoRepeat {
    
    public int lengthOfLongestSubstring(String s) {
        
        
        Map<Character, Integer> map = new HashMap();
        int max = 0;
        
        int start = 0;
        int i = 0;
        
        while(i < s.length()) {
            
            char c = s.charAt(i);
            if (map.containsKey(c) && start <= map.get(s.charAt(i))) {
                start = map.get(c) + 1;
            } else {
                max = Math.max(max, i - start + 1);
            }
            
            map.put(s.charAt(i), i);
            
            i++;
        }
        

        return max;
    }
    
}