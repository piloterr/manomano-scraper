#!/usr/bin/env bash
# ManoMano Product Offer API – cURL example
# Usage: PILOTERR_API_KEY=your-key ./product_offer.sh
set -euo pipefail

curl -sS -G "https://api.piloterr.com/v2/manomano/product/offer" \
  -H "x-api-key: ${PILOTERR_API_KEY}" \
  --data-urlencode query=28910102 \
  --data-urlencode region=FR
