class IsIsomorphic {
    public boolean isIsomorphic(String s, String t) {
        
        Map<Character, Character> map = new HashMap();        
        
        for(int i = 0; i < s.length(); i++) {
            Character sAti = s.charAt(i);
            Character tAti = t.charAt(i);
            
            if(map.containsKey(sAti)) {
                if (map.get(sAti) != tAti) {
                    return false;
                }
                
            } else if (map.containsValue(tAti)) {
                
                return false;
                
            } else {
                map.put(sAti, tAti);

            }
        }
        
        return true; 
    }
}