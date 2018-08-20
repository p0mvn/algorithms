class MoveZeroes {
    public void moveZeroes(int[] nums) {
        
        int insertAt = 0;

        // remove all 0 till insertAt
        for(int x: nums) {
            if(x != 0) {
                nums[insertAt++] = x;
            }
        }
        // insert 0 from insertAt till end
        while(insertAt < nums.length) {
            nums[insertAt++] = 0;
        }
        
        
        
    }
}