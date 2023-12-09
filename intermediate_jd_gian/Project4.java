public class Project4
{

    // Declare objects for later use
    static InputValidator iv = new InputValidator();
    static Flight[] flights = new Flight[4];

    public static void main(String[] args)
    {
        // Initialize program and print menu
        addFlights();
        System.out.println("Welcome to Redbird Airlines!");
        System.out.println("Choose one of the following:");
        System.out.println("  L - list available flights");
        System.out.println("  R - reserve seats");
        System.out.println("  Q - quit");
        char c = iv.readChar("Enter your choice: ");
        System.out.println();
        // While the user doesn't enter "Q", run this code
        while (c != 'Q' && c != 'q')
        {
            // If input is "L", list flights with available seats
            if (c == 'L' || c == 'l')
            {
                listFlights();
                // Print menu again
                System.out.println();
                System.out.println("Choose one of the following:");
                System.out.println("  L - list available flights");
                System.out.println("  R - reserve seats");
                System.out.println("  Q - quit");
                c = iv.readChar("Enter your choice: ");
                System.out.println();
            }
            else if (c == 'R' || c == 'r')
            {
                // If input is "R", ask for flight and number of seats and reserve seats if
                // possible
                reserveSeats();

                // Print menu again
                System.out.println("Choose one of the following:");
                System.out.println("  L - list available flights");
                System.out.println("  R - reserve seats");
                System.out.println("  Q - quit");
                c = iv.readChar("Enter your choice: ");
                System.out.println();
            }
        }
        // Print goodbye message
        System.out.println("Goodbye!");
    }

    public static void addFlights()
    {
        // Add flight to each array element
        flights[0] = new Flight("BT7274", "Dallas", "Normal", "CRJ-700");
        flights[1] = new Flight("VS8156", "Fremont", "Portland", "A321");
        flights[2] = new Flight("FD5574", "Juneau", "Key West", "737-800");
        flights[3] = new Flight("GZ9601", "Sacramento", "Atlanta", "X");
    }

    public static void listFlights()
    {
        for (Flight f : flights)
        {
            // List flights using for each loop
            if (f.getSeatsAvailable() > 0)
            {
                System.out.println(f);
            }
        }
    }

    public static void reserveSeats()
    {
        // Ask which flight
        String resFlight = iv.readString("On which flight? ");
        // Ask for # of seats and reserve on given flight
        if (resFlight.equalsIgnoreCase("BT7274"))
        {
            int numSeats = iv.readInt("How many seats would you like? ", 0, Integer.MAX_VALUE);
            if (flights[0].reserve(numSeats))
            {
                System.out.println("Reservation successful.");
                System.out.println();
            }
            else
            {
                System.out.println("Sorry, not enough seats.");
                System.out.println();
            }
        }
        else if (resFlight.equalsIgnoreCase("VS8156"))
        {
            int numSeats = iv.readInt("How many seats would you like? ", 0, Integer.MAX_VALUE);
            if (flights[1].reserve(numSeats))
            {
                System.out.println("Reservation successful.");
                System.out.println();
            }
            else
            {
                System.out.println("Sorry, not enough seats.");
                System.out.println();
            }
        }
        else if (resFlight.equalsIgnoreCase("FD5574"))
        {
            int numSeats = iv.readInt("How many seats would you like? ", 0, Integer.MAX_VALUE);
            if (flights[2].reserve(numSeats))
            {
                System.out.println("Reservation successful.");
                System.out.println();
            }
            else
            {
                System.out.println("Sorry, not enough seats.");
                System.out.println();
            }
        }
        else if (resFlight.equalsIgnoreCase("GZ9601"))
        {
            int numSeats = iv.readInt("How many seats would you like? ", 0, Integer.MAX_VALUE);
            if (flights[3].reserve(numSeats))
            {
                System.out.println("Reservation successful.");
                System.out.println();
            }
            else
            {
                System.out.println("Sorry, not enough seats.");
                System.out.println();
            }
        }
        else
        {
            // Print if flight # not recognized
            System.out.println("Flight number not recognized.");
        }
    }
}
