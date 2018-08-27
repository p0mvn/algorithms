class PalindromeNumber {
    public boolean isPalindrome(int x) {
        if(x < 0) return false;
        int y = 0;
        int t = x;
        while(t > 0) {
            y = y * 10 + t % 10;
            t /= 10;
        }
        return y == x;        
    }
}