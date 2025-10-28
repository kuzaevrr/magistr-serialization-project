package ram.ka.ru.models.pojos;// BoardingPass.java
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.Accessors;

import java.time.LocalDateTime;
import java.util.List;

@Getter
@Setter
@Accessors(chain = true)
public class BoardingPass {

    private String ticketNo;
    private Integer flightId;
    private String seatNo;
    private Integer boardingNo;
    private LocalDateTime boardingTime;

}