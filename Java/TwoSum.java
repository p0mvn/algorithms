class TwoSum {
    
    
    public int[] twoSum(int[] nums, int target) {
        
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        
        for(int i = 0; i < nums.length; i++) {
            int diff = target - nums[i];
            if(map.containsKey(diff)) {
                return new int[] {map.get(diff), i};
            } else {
                map.put(nums[i], i);
            }
        }
        
        return new int[0];
    }
    
    
    
    public int[] twoSumFirstAttempt(int[] nums, int target) {
        
        
        
        int[] res = new int[2];
        
        for(int i = 0; i < nums.length; i++) {
            for(int j = 0; j < nums.length; j++) {
                // System.out.println(i + " " + j);
                if(nums[i] + nums[j] == target && i != j) {
                    res[0] = i;
                    res[1] = j;
                    return res;
                }
            }
        }
        
        return res;
        
        
        
    }
}