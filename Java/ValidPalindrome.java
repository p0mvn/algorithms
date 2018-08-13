class ValidPalinfrome {
    public boolean isPalindrome(String s) {
        int indexLeft = 0;
        int indexRight = s.length() - 1;

        while (indexLeft < indexRight) {
            char left = s.charAt(indexLeft);
            char right = s.charAt(indexRight);
            boolean leftGood = Character.isLetterOrDigit(left);
            boolean rightGood = Character.isLetterOrDigit(right);
           
            if (leftGood && rightGood) {
                if(Character.toLowerCase(left) != Character.toLowerCase(right)) {
                    return false;
                }
                indexLeft++;
                indexRight--;
            } else {
                if(!leftGood) {
                    indexLeft++;
                }
                if(!rightGood) {
                    indexRight--;
                }
            }
        }
        return true;
    }
}