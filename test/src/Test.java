import java.text.ParseException;
import java.util.ArrayList;
import java.util.Currency;
import java.util.List;

public class Test {

    public static void main(String[] args) throws ParseException {
        List s = new ArrayList<>();
        s.add("s");
        s.add("e");

        System.out.println("..." + Currency.getInstance("CNY"));
        System.out.println("..." + Currency.getInstance("CNH"));

    }


}
