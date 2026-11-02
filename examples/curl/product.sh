#!/usr/bin/env bash
# ManoMano Product API – cURL example
# Usage: PILOTERR_API_KEY=your-key ./product.sh
set -euo pipefail

curl -sS --fail-with-body -G "https://api.piloterr.com/v2/manomano/product" \
  -H "x-api-key: ${PILOTERR_API_KEY}" \
  --data-urlencode 'query=https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588'
