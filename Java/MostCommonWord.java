class MostCommonWord {
    
    public String mostCommonWord(String paragraph, String[] banned) {
        
        paragraph = paragraph.replaceAll("[!?',;.]", "").toLowerCase();
        Map<String, Integer> map = new HashMap();
        Set<String> ban = new HashSet();
        
        for(String word: banned) {
            ban.add(word);
        }
        
        int max = 0;
        String maxWord = null;
        
        for(String word : paragraph.split(" ")) {
            
            if(!ban.contains(word)) {
                
                if(map.containsKey(word)) {
                    map.put(word, map.get(word) + 1);
                } else {
                    map.put(word, 1);
                }
                int count = map.get(word);
                if(count > max) {
                    max = count;
                    maxWord = word;
                }  
                
            }   
        }
        
        return maxWord;
        
    }
    
    
    
    
    public String mostCommonWordFirstTry(String paragraph, String[] banned) {
        
        paragraph = paragraph.replaceAll("[!?',;.]", "").toLowerCase();
        String[] words = paragraph.split(" ");
        
        Arrays.sort(words);
        
        String prev = "";
        int count = 0;
        int max = 0;
        String maxWord = "";
        Set<String> bannedSet = new HashSet(Arrays.asList(banned));
        for(String word : words) {
            if(!bannedSet.contains(word)) {

                if(word.equals(prev)) {
                    count++;
                } else {
                    count = 1;
                }


                if(count > max) {
                    max = count;
                    maxWord = word;
                }
                
            } else {
                count = 0;
            }
            prev = word;
        }

        
        return maxWord;
        
    }
}