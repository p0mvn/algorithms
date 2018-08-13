class ToLowerCase {
    
    public String toLowerCase(String str) {
        String low = "";
        for(int i = 0; i < str.length(); i++) {
            low += String.valueOf(str.charAt(i)).toLowerCase();
        }
        return low;
        
    }
}