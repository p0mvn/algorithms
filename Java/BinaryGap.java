class BinaryGap {
    public int binaryGap(int N) {
        
        String bin = toBinary(N, "");
        System.out.println(bin);
        
        return helper(bin, 0, 0);
            
        
    }
    
    private int helper(String bin, int max, int count) {
        char cur = bin.charAt(0);

        if(bin.length() == 1) {
            if ( cur == '0') {
                return max;
            } else {
                return Math.max(max, count);
            }
            
        }
        if(cur == '0' && count == 0) {
            return helper(bin.substring(1), max, count);
        } else if(cur == '1' && count > 0){
            return helper(bin.substring(1), Math.max(count, max), 1);
            
        } else {
            count++;
            return helper(bin.substring(1), max, count);
        } 
        
    }

    
        
        
       
    private String toBinary(int N, String acc) {
        if (N <= 0) {
            return acc;
        } 
        if(N % 2 == 0) {
            return toBinary(N / 2, "0" + acc);
        } else {
            return toBinary(N / 2, "1" + acc);
        }
    }
    
}