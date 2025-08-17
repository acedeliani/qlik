#!/usr/bin/env bash

DATA_DIR=testdata
DATA=$(cat "$DATA_DIR/order.json")
BASE_URL="localhost:8080"

# List by customer
echo "Listing by customer"
CUSTOMER_ID="01"
curl -X POST -d "$DATA" "$BASE_URL/items/$CUSTOMER_ID" 2>/dev/null | jq .
echo -e "\n------------------------\n"

# List summary
echo "Listing summary"
curl -X POST -d "$DATA" "$BASE_URL/items/summary" 2>/dev/null | jq .
echo -e "\n------------------------"
