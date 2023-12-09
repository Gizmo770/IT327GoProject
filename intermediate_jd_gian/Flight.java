public class Flight
{
    private String flightNumber;
    private String originCity;
    private String destinationCity;
    private String aircraftType;
    private int maxSeats;
    private int currentResSeats;

    public Flight(String flightNum, String oCity, String dCity, String aType)
    {
        flightNumber = flightNum;
        originCity = oCity;
        destinationCity = dCity;
        aircraftType = aType;
        currentResSeats = 0;
        calcMaximumSeats(aType);
    }

    /**
     * Method to convert object to a String format
     * 
     * @return Object in a String format
     */
    public String toString()
    {
        // String flightString = flightNumber + " " + originCity + " " + destinationCity
        // + " "
        // + getSeatsAvailable() + " seats available";
        String flightString = String.format("%-10s %-10s %-10s %-3d seats available", flightNumber, originCity,
                destinationCity, getSeatsAvailable());
        return flightString;
    }

    /**
     * Method that returns the number of seats available on a flight
     * 
     * @return Number of seats available
     */
    public int getSeatsAvailable()
    {
        return (maxSeats - currentResSeats);
    }

    /**
     * Method that determines whether or not it is possible to reserve given number
     * of seats. If possible, reserve seats and return true. If not, return false.
     * 
     * @param seats Number of seats
     * @return true or false based on seats available
     */
    public boolean reserve(int seats)
    {
        if (seats <= getSeatsAvailable())
        {
            currentResSeats += seats;
            return true;
        }
        else
        {
            return false;
        }
    }

    /**
     * Method that calulates the maximum number of seats on a flight based on
     * aircraft type
     * 
     * @param aType Aircraft type
     */
    private void calcMaximumSeats(String aType)
    {
        if (aType.equalsIgnoreCase("737-800"))
        {
            maxSeats = 160;
        }
        else if (aType.equalsIgnoreCase("A321"))
        {
            maxSeats = 196;
        }
        else if (aType.equalsIgnoreCase("CRJ-700"))
        {
            maxSeats = 63;
        }
        else
        {
            maxSeats = 0;
        }
    }
}
