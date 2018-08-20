class ValidParentheses {
    public boolean isValid(String s) {
        if(s.length() % 2 == 1) return false;
        if(s.length() == 0) return true;
        
        Stack<Character> stack = new Stack();
        
        for(int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if(isClosingParen(c)) {
                if(stack.empty() || !isOk(stack.pop(), c)) {
                    return false;
                }
            } else {
                stack.push(c);
            }
        }
    
        return stack.empty();
    }
    
    private boolean isOk(char a, char b) {
        if ((a == '(' && b == ')') || (a == '{' && b == '}') || (a == '[' && b == ']')) {
            return true;
        } else {
            return false;
        }
    }
    
    private boolean isClosingParen(char c) {
        return c == ')' || c == ']' || c == '}';
    }
    

}