package ram.ka.ru.models.pojos;// Booking.java
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.Accessors;

import java.math.BigDecimal;
import java.time.LocalDateTime;

@Getter
@Setter
@Accessors(chain = true)
public class Booking {

    private String bookRef;
    private LocalDateTime bookDate;
    private BigDecimal totalAmount;

}