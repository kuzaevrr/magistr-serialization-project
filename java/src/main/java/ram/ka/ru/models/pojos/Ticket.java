package ram.ka.ru.models.pojos;

import lombok.Getter;
import lombok.Setter;
import lombok.experimental.Accessors;

import java.time.LocalDateTime;

@Getter
@Setter
@Accessors(chain = true)
public class Ticket {

    private String ticketNo;
    private String bookRef;
    private String passengerId;
    private String passengerName;
    private Boolean outbound;
    private LocalDateTime departureTime;
    private Booking booking;

}