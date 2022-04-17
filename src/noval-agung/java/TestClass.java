import java.util.*;

class TestClass {
    public static void main(String args[] ) throws Exception {
      
        
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
    
        String name = "";
        while(n >=0) {
            name = s.nextLine();  
            n--;
        }
        System.out.println("Hi, " + name + ".");

    }
}