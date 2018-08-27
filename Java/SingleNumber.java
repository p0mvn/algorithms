class Solution {
    

    
    public int singleNumber(int[] nums) {
        
        Map<Integer, Boolean> map = new HashMap();
        
        for(int i: nums) {
            if(map.containsKey(i)) {
                map.put(i, false);
            } else {
                map.put(i, true);
            }
        }
        
        for(int key: map.keySet()) {
            if(map.get(key)) {
                return key;
            }
        }
        
        return -1;
        
    }
}