package ram.ka.ru.models.pojos;// Segment.java
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.Accessors;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;

@Getter
@Setter
@Accessors(chain = true)
public class Segment {

    private String ticketNo;
    private Integer flightId;
    private String fareConditions;
    private BigDecimal price;
    private Ticket ticket;
    private List<BoardingPass> boardingPasses;

}