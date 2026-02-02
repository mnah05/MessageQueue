#!/usr/bin/env bash

URL="http://localhost:8080/new-user"
CONTENT_TYPE="Content-Type: application/json"

for i in $(seq 1 100); do
  (
    curl -s -X POST "$URL" \
      -H "$CONTENT_TYPE" \
      -d "{\"email\":\"user$i@example.com\"}" \
      -w "request $i → status %{http_code}\n"
  ) &
done

wait
echo "All requests sent"
