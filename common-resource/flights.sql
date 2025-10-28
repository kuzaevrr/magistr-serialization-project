SELECT
    json_build_object(
            'flightId', f.flight_id,
            'routeNo', f.route_no,
            'status', f.status,
            'scheduledDeparture', f.scheduled_departure,
            'scheduledArrival', f.scheduled_arrival,
            'actualDeparture', f.actual_departure,
            'actualArrival', f.actual_arrival,
            'segments', (
                SELECT json_agg(
                               json_build_object(
                                       'ticketNo', s.ticket_no,
                                       'flightId', s.flight_id,
                                       'fareConditions', s.fare_conditions,
                                       'price', s.price,
                                       'ticket', (
                                           SELECT json_build_object(
                                                          'ticketNo', t.ticket_no,
                                                          'passengerId', t.passenger_id,
                                                          'passengerName', t.passenger_name,
                                                          'outbound', t.outbound,
                                                          'booking', (
                                                              SELECT json_build_object(
                                                                             'bookRef', b.book_ref,
                                                                             'bookDate', b.book_date,
                                                                             'totalAmount', b.total_amount
                                                                     )
                                                              FROM bookings b
                                                              WHERE b.book_ref = t.book_ref
                                                          )
                                                  )
                                           FROM tickets t
                                           WHERE t.ticket_no = s.ticket_no
                                       ),
                                       'boardingPasses', (
                                           SELECT json_agg(
                                                          json_build_object(
                                                                  'seatNo', bp.seat_no,
                                                                  'boardingNo', bp.boarding_no,
                                                                  'boardingTime', bp.boarding_time
                                                          )
                                                  )
                                           FROM boarding_passes bp
                                           WHERE bp.ticket_no = s.ticket_no AND bp.flight_id = s.flight_id
                                       )
                               )
                       )
                FROM segments s
                WHERE s.flight_id = f.flight_id
            )
    ) as flight_data
FROM flights f
ORDER BY f.flight_id;