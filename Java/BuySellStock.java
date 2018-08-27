class BuySellStock {
    public int maxProfit(int[] prices) {
        
        int min = Integer.MAX_VALUE;
        int max = Integer.MIN_VALUE;
        int maxDiff = 0;
        for(int i = 0; i < prices.length; i++) {
            if(prices[i] < min) {
                min = prices[i];
                max = prices[i];
            } else {
                if(prices[i] > max) {
                    max = prices[i];
                }
            }
            if(maxDiff < (max - min)) {
                    
                    maxDiff = max - min; 
                }
        }
        
        return maxDiff;
    }
}