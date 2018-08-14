class UncommonWords {
    public String[] uncommonFromSentences(String A, String B) {
        
        HashMap<String, Integer> hmap = new HashMap<String, Integer>();
        
        String[] aWords = A.split("\\s+");
        String[] bWords = B.split("\\s+");
        
        int upperBound = 0;
        

        
        
         for(int i = 0; i < Math.max(bWords.length, aWords.length); i++) {
            if(i < aWords.length) {
                if(hmap.containsKey(aWords[i])) {
                    hmap.put(aWords[i], 2);
                } else {
                    hmap.put(aWords[i], 1);
                } 
            } 
             
            if(i < bWords.length) {
                if(hmap.containsKey(bWords[i])) {
                    hmap.put(bWords[i], 2);
                } else {
                    hmap.put(bWords[i], 1);
                } 
            }  
            
        }        
        
        
        Set<String> keySet = hmap.keySet();
        
        
        for(String s: keySet) {
            if(hmap.get(s) == 1) {
                upperBound++;
            }
        }
        String[] result = new String[Math.max(0,upperBound)];
        
        
        int j = 0;
        for(String s: keySet) {
            if(hmap.get(s) == 1) {

                result[j] = s;
                j++;
            }
        }
        
        
        return result;
        
        
        
        
    }
    
    
    
}