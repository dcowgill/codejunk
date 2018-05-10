namespace go thrifttest

enum OrderType {
  DELIVERY = 1,
  PICKUP = 2,
  REDELIVERY = 3
}

enum OrderStatus {
  PENDING = 1,
  AUTHORIZED = 2,
  NEW = 3,
  REJECTED = 4,
  ACCEPTED = 5,
  CONFIRMED = 6,
  ALMOST_READY = 7,
  READY = 8,
  IN_TRANSIT = 9,
  DELIVERED = 10,
  PICKED_UP_BY_CUSTOMER = 11,
  COULD_NOT_DELIVER = 12,
  CANCELED = 13,
  FAILED = 14,
  ABANDONED = 15
}

// struct Amount {
//   1: i64 fractional,
//   2: string currencyCode
// }

struct Order {
  1: i64 id,
  2: i64 restaurant_id,
  3: i64 user_id,
  4: OrderType order_type,
  5: OrderStatus status,
  7: i64 address_id,
  8: i16 estimated_prep_time,
  9: string order_target_date,
  10: string created_at,
  11: string acknowledged_at,
  12: string almost_ready_at,
  13: string ready_at,
  14: string delivered_at,
  15: i64 subtotal,
  16: i64 delivery_fee,
  17: i64 tip,
  18: i64 tax,
  19: i64 card_fee,
  20: i64 total,
  21: string submitted_at,
  22: i64 adjusted_subtotal,
  23: i64 zone_id,
  25: string completion_target,
  26: string confirmed_at,
  27: i64 estimated_travel_time,
  28: string updated_at,
  29: i64 payment_token_id,
  30: i16 drivers_needed,
  31: i16 estimated_loading_time,
  32: string target_confirmed_by,
  33: string target_send_driver_at,
  34: string target_ready_at,
  35: string target_delivered_at,
  36: string project_code,
  37: i64 corporate_group_id,
  38: bool manual_assignment,
  39: bool contains_alcohol,
  40: i64 fee,
  41: i64 surcharge,
  42: double estimated_distance,
  43: double estimated_raw_travel_time
}
