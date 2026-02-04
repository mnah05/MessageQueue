#!/bin/bash

API_URL="http://localhost:8080"

ORDER_ID="ORD-$(date +%s)"
CUSTOMER_ID="CUST-001"
EMAIL="customer@example.com"
AMOUNT=199.99

echo "Creating order..."
echo "Order ID: $ORDER_ID"

curl -X POST "$API_URL/orders" \
  -H "Content-Type: application/json" \
  -d "{
    \"order_id\": \"$ORDER_ID\",
    \"customer_id\": \"$CUSTOMER_ID\",
    \"email\": \"$EMAIL\",
    \"amount\": $AMOUNT,
    \"items\": [\"item1\", \"item2\", \"item3\"]
  }"

echo
echo "Done."
