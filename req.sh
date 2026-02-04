#!/bin/bash

API_URL="http://localhost:8080"

echo "Sending 10 email-only requests..."

for i in {1..10}; do
  curl -s -X POST "$API_URL/new-user" \
    -H "Content-Type: application/json" \
    -d "{
      \"email\": \"user$i@example.com\"
    }" > /dev/null

  echo "Email request $i sent"
done

echo
echo "Sending 10 order requests..."

for i in {1..10}; do
  ORDER_ID="ORD-$(date +%s)-$i"

  curl -s -X POST "$API_URL/orders" \
    -H "Content-Type: application/json" \
    -d "{
      \"order_id\": \"$ORDER_ID\",
      \"customer_id\": \"CUST-$i\",
      \"email\": \"customer$i@example.com\",
      \"amount\": $((100 + i * 10)),
      \"items\": [\"item1\", \"item2\"]
    }" > /dev/null

  echo "Order request $i sent ($ORDER_ID)"
done

echo
echo "Done."
