public class Odd {
    public static void main(String[] args) {
        IsOdd id = new IsOdd();
        Boolean result = id.isOdd(10);
        System.out.println(result);
    }
}

class IsOdd {
    public boolean isOdd(int nums) {
        return (nums & 1) == 1;
    }
}