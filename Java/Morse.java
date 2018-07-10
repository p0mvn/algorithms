class Morse {
    
    private String[] morse = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
  
  public int uniqueMorseRepresentations(String[] words) {
      
      
      
      HashSet<String> uniqueRepresentations = new HashSet();
      
      for(String w: words) {
          uniqueRepresentations.add(convertToMorse(w));
      }
      return uniqueRepresentations.size();
      
  }
  
  private String convertToMorse(String word) {
      String result = "";
      for(int i = 0; i < word.length(); i++) {
          result += morse[word.charAt(i) - 'a'];
      }
      return result.toString();
  }
}