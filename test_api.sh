#!/bin/bash

# Curl settings
URL="http://localhost:3000/api/v1/stopword"
API_KEY="3c05efa8-14c3-4413-834e-e2dfc982fbf9"

# Test array
ENGLISH_SENTENCES=(
  "Max Lesage is a great person, he lives in Paris at the moment"
  "I'm doing some tests currently"
  "I have a lot of problem right now, some with Max and other dealing with the fact I leave in Niort ahah"
  "Tests aren't bad yet, I think this might be good for english speakers"
  "Let's try this in prod at the Paris game's week tournament"
)

# A burst of tests
for s in "${ENGLISH_SENTENCES[@]}"; do
  BEFORE=$(date +"%s.%3N")
  curl -X POST $URL \
    -H "X-API-KEY: $API_KEY" \
    -H "Content-Type: application/json" \
    -d '{"text": "'"$s"'"}'
  echo
  AFTER=$(date +"%s.%3N")
  ELAPSED=$(echo "$AFTER - $BEFORE" | bc)
  echo "L'API a répondu en 0$ELAPSED secondes"
  echo
done
