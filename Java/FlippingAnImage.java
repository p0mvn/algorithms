class FlippingAnImage {
    public int[][] flipAndInvertImage(int[][] A) {
        
        for(int[] x : A) {
            // reverse
            reverseArray(x); // pass by reference 
            
            
            //flip
            flipArray(x); // again by reference 
            
            
        }
        
        return A;
        
    }
    
    
    private void reverseArray(int[] x) {
        int start = 0;
        int end = x.length - 1;
        while(start < end) {
            int temp = x[start];
            x[start] = x[end];
            x[end] = temp;
            start++;
            end--;
        }
    }
    
    private void flipArray(int[] x) {
        for(int i = 0; i < x.length; i++) {
            if (x[i] == 1) {
                x[i] = 0;
            } else {
                x[i] = 1;
            }
        }
    }
}