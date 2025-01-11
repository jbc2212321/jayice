public class Hello {

    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("hello my world!");
        } else {
            System.out.println("hello my world! " + args[0]);
        }
    }

}
