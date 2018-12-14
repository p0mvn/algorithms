class SortArrayByParity {
    public int[] sortArrayByParity(int[] A) {
        int pointer = 0; 
        for(int i = 0; i < A.length; i++) {
            if (A[i] % 2 == 0) {
                int temp = A[pointer];
                A[pointer] = A[i];
                A[i] = temp;
                pointer++;
            }
        }
        return A;
    }
}