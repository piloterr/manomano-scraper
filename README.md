<p align="center">
  <a href="https://www.piloterr.com"><img src="images/banner.png" alt="ManoMano Scraper API by Piloterr – scrape ManoMano data as JSON" width="100%"></a>
</p>

# ManoMano Scraper API

[![Discord](https://img.shields.io/badge/Discord-Join%20the%20community-5865F2?logo=discord&logoColor=white)](https://discord.gg/CFj3ph8aHU) [![Get API key](https://img.shields.io/badge/Piloterr-Get%20a%20free%20API%20key-111827)](https://app.piloterr.com) [![API docs](https://img.shields.io/badge/Docs-OpenAPI-0EA5E9)](https://app.piloterr.com/openapi) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

The **Piloterr ManoMano Scraper API** lets you extract public data from ManoMano as **clean, structured JSON** — no proxies, headless browsers or HTML parsing to maintain. Endpoints available in production:

- **ManoMano Product API** — Extract full product details from a ManoMano product page, including price, seller, brand, images, variants, and stock.
- **ManoMano Product Offer API** — Get all marketplace seller offers for a ManoMano product by product ID, including pricing, shipping, and seller information.
- **ManoMano Search API** — Search for products on ManoMano by category or keyword URL, returning listings with prices, ratings, and seller information.

Use it for price monitoring, assortment tracking, competitor analysis, product research or catalog enrichment.

| Endpoint | Path | Credits | What you get |
|---|---|---|---|
| [ManoMano Product API](#manomano-product-api) | `/v2/manomano/product` | 1 | Extract full product details from a ManoMano product page, including price, seller, brand, images, variants, and stock. |
| [ManoMano Product Offer API](#manomano-product-offer-api) | `/v2/manomano/product/offer` | 1 | Get all marketplace seller offers for a ManoMano product by product ID, including pricing, shipping, and seller information. |
| [ManoMano Search API](#manomano-search-api) | `/v2/manomano/search` | 1 | Search for products on ManoMano by category or keyword URL, returning listings with prices, ratings, and seller information. |

## Table of contents

- [How it works](#how-it-works)
- [ManoMano Product API](#manomano-product-api)
- [ManoMano Product Offer API](#manomano-product-offer-api)
- [ManoMano Search API](#manomano-search-api)
- [Code examples in other languages](#code-examples-in-other-languages)
- [Practical use cases](#practical-use-cases)
- [Why Piloterr?](#why-piloterr)
- [FAQ](#faq)
- [Related scrapers](#related-scrapers)
- [Support](#support)

## How it works

1. **Create a free Piloterr account** and copy your API key from the [dashboard](https://app.piloterr.com).
2. **Send a `GET` (or `POST`) request** to `https://api.piloterr.com/v2/...` with the header `x-api-key: YOUR_API_KEY`.
3. **Receive structured JSON** — Piloterr handles rendering, proxies, retries and parsing.

```bash
export PILOTERR_API_KEY="your-api-key"
pip install requests        # Python
# TypeScript: Node.js 18+ (native fetch), run with `npx tsx file.ts`
```

> Every endpoint accepts both `GET` (query string) and `POST` (JSON body with the same fields).

## ManoMano Product API

Extract full product details from a ManoMano product page, including price, seller, brand, images, variants, and stock.

`GET | POST https://api.piloterr.com/v2/manomano/product` · **1 credit per request** · [Full reference](https://app.piloterr.com/openapi)

### Request parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `query` | string | ✅ | Full ManoMano product page URL. Optionally include the `model_id` query parameter to target a specific variant. |

### Python example

```python
import json
import os
import requests

response = requests.get(
    "https://api.piloterr.com/v2/manomano/product",
    headers={"x-api-key": os.environ["PILOTERR_API_KEY"]},
    params={
        "query": "https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588",
    },
    timeout=60,
)
response.raise_for_status()
data = response.json()

print(json.dumps(data, indent=2, ensure_ascii=False)[:2000])
```

### TypeScript example

```typescript
type ManomanoProduct = {
  ean?: string;
  brand?: {
    id?: number;
    url?: string;
    name?: string;
  };
  media?: {
    type?: string;
    large_url?: string;
    regular_url?: string;
    thumbnail_url?: string;
  }[];
  mm_id?: number;
  price?: {
    primary_price?: {
      amount_with_vat?: number;
      measurement_unit?: string;
      amount_without_vat?: number;
    };
    discount_percentage?: number;
    eco_participation_with_vat?: number;
  };
  stock?: {
    minimum_value?: number;
    increase_quantity?: number;
  };
  title?: string;
  seller?: {
    id?: number;
    url?: string;
    name?: string;
    country?: string;
  };
  currency?: string;
  delivery?: {
    weight?: number;
    fulfillment?: boolean;
  };
  model_id?: number;
  offer_id?: string;
  quantity?: {
    stock?: unknown;
    value?: number;
  };
  variants?: {
    list?: {
      name?: string;
      values?: Record<string, unknown>[];
    }[];
  };
};

const params = new URLSearchParams({
  "query": "https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588",
});

const response = await fetch(`https://api.piloterr.com/v2/manomano/product?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY! },
});
if (!response.ok) throw new Error(`Piloterr error ${response.status}: ${await response.text()}`);

const data = (await response.json()) as ManomanoProduct;

console.log(JSON.stringify(data, null, 2).slice(0, 2000));
```

<details>
<summary><b>POST variant</b></summary>

```http
POST https://api.piloterr.com/v2/manomano/product
x-api-key: YOUR_API_KEY
Content-Type: application/json

{"query": "https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588"}
```
</details>

### Output example

Example response from the API reference, abridged — real values depend on the page you scrape.

```json
{
  "ean": "3662970077672",
  "brand": {
    "id": 6502,
    "url": "/marque/outsunny-6502",
    "name": "OUTSUNNY"
  },
  "media": [
    {
      "type": "image",
      "large_url": "https://cdn.manomano.com/images/images_products/6132097/L/28916742_1.jpg",
      "regular_url": "https://cdn.manomano.com/images/images_products/6132097/P/28916742_1.jpg",
      "thumbnail_url": "https://cdn.manomano.com/images/images_products/6132097/T/28916742_1.jpg"
    }
  ],
  "mm_id": 127463740566,
  "price": {
    "primary_price": {
      "amount_with_vat": 335.9,
      "measurement_unit": "",
      "amount_without_vat": 279.92
    },
    "discount_percentage": 0,
    "eco_participation_with_vat": 5.25
  },
  "stock": {
    "minimum_value": 1,
    "increase_quantity": 1
  },
  "title": "Pergola rétractable OUTSUNNY - 2,98L x 2,13l x 2,22H m - Structure métal époxy anticorrosion",
  "seller": {
    "id": 6132097,
    "url": "/marchand-6132097",
    "name": "Aosom France MF",
    "country": "FR"
  },
  "currency": "EUR",
  "delivery": {
    "weight": 40,
    "fulfillment": true
  },
  "model_id": 17784588,
  "offer_id": "afa4d772-b372-4f8a-b3d0-0f37ee205b58",
  "quantity": {
    "stock": null,
    "value": 1
  },
  "variants": {
    "list": [
      {
        "name": "couleur",
        "values": [
          {}
        ]
      }
    ]
  }
}
```

### Response fields

| Field | Type | Description |
|---|---|---|
| `ean` | string |  |
| `brand` | object |  |
| `brand.id` | integer |  |
| `brand.url` | string |  |
| `brand.name` | string |  |
| `media[]` | object[] |  |
| `media[].type` | string |  |
| `media[].large_url` | string |  |
| `media[].regular_url` | string |  |
| `media[].thumbnail_url` | string |  |
| `mm_id` | integer |  |
| `price` | object |  |
| `price.primary_price` | object |  |
| `price.discount_percentage` | integer |  |
| `price.eco_participation_with_vat` | number |  |
| `stock` | object |  |
| `stock.minimum_value` | integer |  |
| `stock.increase_quantity` | integer |  |
| `title` | string |  |
| `seller` | object |  |
| `seller.id` | integer |  |
| `seller.url` | string |  |
| `seller.name` | string |  |
| `seller.country` | string |  |
| `currency` | string |  |

## ManoMano Product Offer API

Get all marketplace seller offers for a ManoMano product by product ID, including pricing, shipping, and seller information.

`GET | POST https://api.piloterr.com/v2/manomano/product/offer` · **1 credit per request** · [Full reference](https://app.piloterr.com/openapi)

### Request parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `query` | string | ✅ | ManoMano product ID. This is the numeric `product_id` returned by the ManoMano Product endpoint. |
| `region` | string | ✅ | Region code for the ManoMano platform. Must match the original product's domain. |

### Python example

```python
import json
import os
import requests

response = requests.get(
    "https://api.piloterr.com/v2/manomano/product/offer",
    headers={"x-api-key": os.environ["PILOTERR_API_KEY"]},
    params={
        "query": "28910102",
        "region": "FR",
    },
    timeout=60,
)
response.raise_for_status()
data = response.json()

print(json.dumps(data, indent=2, ensure_ascii=False)[:2000])
```

### TypeScript example

```typescript
type ManomanoProductOffer = {
  count?: number;
  results?: {
    sku?: {
      title?: string;
      master_product?: Record<string, unknown>;
    };
    market?: string;
    seller?: {
      is_m_f?: boolean;
      contract_id?: string;
      contract_name?: string;
      country_location?: string;
    };
    pricing?: {
      sell_price?: Record<string, unknown>;
      retail_price?: unknown;
      eco_participation?: Record<string, unknown>;
    };
    packaging?: {
      increment?: string;
      minimum_quantity?: string;
    };
    product_id?: string;
    is_sellable?: boolean;
    delivery_offers?: {
      estimate_delivery?: Record<string, unknown>;
      delivery_countries?: string[];
    }[];
  }[];
};

const params = new URLSearchParams({
  "query": "28910102",
  "region": "FR",
});

const response = await fetch(`https://api.piloterr.com/v2/manomano/product/offer?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY! },
});
if (!response.ok) throw new Error(`Piloterr error ${response.status}: ${await response.text()}`);

const data = (await response.json()) as ManomanoProductOffer;

console.log(JSON.stringify(data, null, 2).slice(0, 2000));
```

<details>
<summary><b>POST variant</b></summary>

```http
POST https://api.piloterr.com/v2/manomano/product/offer
x-api-key: YOUR_API_KEY
Content-Type: application/json

{"query": "28910102", "region": "FR"}
```
</details>

### Output example

Example response from the API reference, abridged — real values depend on the page you scrape.

```json
{
  "count": 2,
  "results": [
    {
      "sku": {
        "title": "Modern single basin sink mixer tap slotted drain curved spout chrome bathroom",
        "master_product": {}
      },
      "market": "B2C",
      "seller": {
        "is_m_f": false,
        "contract_id": "36769168",
        "contract_name": "COOCHEER",
        "country_location": "FR"
      },
      "pricing": {
        "sell_price": {},
        "retail_price": null,
        "eco_participation": {}
      },
      "packaging": {
        "increment": "1",
        "minimum_quantity": "1"
      },
      "product_id": "137641413",
      "is_sellable": true,
      "delivery_offers": [
        {}
      ]
    }
  ]
}
```

### Response fields

| Field | Type | Description |
|---|---|---|
| `count` | integer |  |
| `results[]` | object[] |  |
| `results[].sku` | object |  |
| `results[].market` | string |  |
| `results[].seller` | object |  |
| `results[].pricing` | object |  |
| `results[].packaging` | object |  |
| `results[].product_id` | string |  |
| `results[].is_sellable` | boolean |  |
| `results[].delivery_offers[]` | object[] |  |

## ManoMano Search API

Search for products on ManoMano by category or keyword URL, returning listings with prices, ratings, and seller information.

`GET | POST https://api.piloterr.com/v2/manomano/search` · **1 credit per request** · [Full reference](https://app.piloterr.com/openapi)

### Request parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `query` | string | ✅ | Full ManoMano search or category page URL. |

### Python example

```python
import json
import os
import requests

response = requests.get(
    "https://api.piloterr.com/v2/manomano/search",
    headers={"x-api-key": os.environ["PILOTERR_API_KEY"]},
    params={
        "query": "https://www.manomano.fr/perceuse-1146",
    },
    timeout=60,
)
response.raise_for_status()
data = response.json()

print(json.dumps(data, indent=2, ensure_ascii=False)[:2000])
```

### TypeScript example

```typescript
type ManomanoSearch = {
  results?: {
    cpc?: number;
    url?: string;
    bulk?: boolean;
    title?: string;
    is_b2b?: boolean;
    is_mmf?: boolean;
    prices?: {
      main?: number;
      unit?: unknown;
      retail?: number;
      discount?: string;
      secondary?: unknown;
    };
    rating?: number;
    brand_id?: number;
    is_local?: boolean;
    model_id?: number;
    offer_id?: string;
    position?: number;
    seller_id?: number;
    article_id?: number;
  }[];
};

const params = new URLSearchParams({
  "query": "https://www.manomano.fr/perceuse-1146",
});

const response = await fetch(`https://api.piloterr.com/v2/manomano/search?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY! },
});
if (!response.ok) throw new Error(`Piloterr error ${response.status}: ${await response.text()}`);

const data = (await response.json()) as ManomanoSearch;

console.log(JSON.stringify(data, null, 2).slice(0, 2000));
```

<details>
<summary><b>POST variant</b></summary>

```http
POST https://api.piloterr.com/v2/manomano/search
x-api-key: YOUR_API_KEY
Content-Type: application/json

{"query": "https://www.manomano.fr/perceuse-1146"}
```
</details>

### Output example

Example response from the API reference, abridged — real values depend on the page you scrape.

```json
{
  "results": [
    {
      "cpc": 0.1,
      "url": "https://www.manomano.fr/makita-perceuse-visseuse-18v-3x3ah-ddf484rf3j",
      "bulk": false,
      "title": "Perceuse visseuse 18V Ø 13 mm - MAKITA - Avec 2 batteries 18V Li-Ion 5.0 Ah, chargeur, MakPak- DDF484RTJ",
      "is_b2b": false,
      "is_mmf": true,
      "prices": {
        "main": 277.2,
        "unit": null,
        "retail": 569.16,
        "discount": "51",
        "secondary": null
      },
      "rating": 4.589109,
      "brand_id": 16,
      "is_local": false,
      "model_id": 54973303,
      "offer_id": "d63853b8-314b-4de0-8cd8-fc9c993fe6f6",
      "position": 1,
      "seller_id": 5307819,
      "article_id": 78566961
    },
    {
      "url": "https://www.manomano.fr/dewalt-perceuse-visseuse-18v-dcd778d2t",
      "bulk": false,
      "title": "Perceuse visseuse sans fil 18V - DEWALT - DCD778D2T",
      "is_b2b": false,
      "is_mmf": false,
      "prices": {
        "main": 149.99,
        "unit": null,
        "retail": 199.99,
        "discount": "25",
        "secondary": null
      },
      "rating": 4.72,
      "model_id": 12345678,
      "offer_id": "abc123-def456",
      "position": 2,
      "article_id": 12345678
    }
  ]
}
```

### Response fields

| Field | Type | Description |
|---|---|---|
| `results[]` | object[] |  |
| `results[].cpc` | number |  |
| `results[].url` | string |  |
| `results[].bulk` | boolean |  |
| `results[].title` | string |  |
| `results[].is_b2b` | boolean |  |
| `results[].is_mmf` | boolean |  |
| `results[].prices` | object |  |
| `results[].rating` | number |  |
| `results[].brand_id` | integer |  |
| `results[].is_local` | boolean |  |
| `results[].model_id` | integer |  |
| `results[].offer_id` | string |  |
| `results[].position` | integer |  |
| `results[].seller_id` | integer |  |
| `results[].article_id` | integer |  |

## Code examples in other languages

Ready-to-run examples live in the [`examples/`](examples) folder. Each one reads your key from the `PILOTERR_API_KEY` environment variable.

| Language | Product | Product Offer | Search |
|---|---|---|---|
| cURL (Bash) | [`product`](examples/curl/product.sh) | [`product-offer`](examples/curl/product_offer.sh) | [`search`](examples/curl/search.sh) |
| PHP | [`product`](examples/php/product.php) | [`product-offer`](examples/php/product_offer.php) | [`search`](examples/php/search.php) |
| C# | [`product`](examples/csharp/Product.cs) | [`product-offer`](examples/csharp/ProductOffer.cs) | [`search`](examples/csharp/Search.cs) |
| Go | [`product`](examples/go/product.go) | [`product-offer`](examples/go/product_offer.go) | [`search`](examples/go/search.go) |
| Java | [`product`](examples/java/ProductExample.java) | [`product-offer`](examples/java/ProductOfferExample.java) | [`search`](examples/java/SearchExample.java) |
| Ruby | [`product`](examples/ruby/product.rb) | [`product-offer`](examples/ruby/product_offer.rb) | [`search`](examples/ruby/search.rb) |
| Node.js | [`product`](examples/nodejs/product.mjs) | [`product-offer`](examples/nodejs/product_offer.mjs) | [`search`](examples/nodejs/search.mjs) |

## Practical use cases

- **Price monitoring** — Track prices, promotions and availability over time to react to competitors.
- **Assortment & catalog tracking** — Follow new products, delisted items and category changes.
- **Competitor analysis** — Benchmark ranges, brands and positioning across retailers.
- **Product research** — Spot best-sellers and trends from search results.
- **Catalog enrichment** — Pull titles, images and attributes into your PIM or marketplace listings.

## Why Piloterr?

- **No infrastructure** — no proxies, headless browsers or CAPTCHA solving to run yourself.
- **Structured JSON** — stable, documented fields instead of brittle HTML selectors.
- **Pay per request** — transparent credit cost per endpoint, failed requests are retried for you.
- **One API key, hundreds of sources** — ManoMano plus many other e-commerce, search and company data endpoints in the [Piloterr library](https://www.piloterr.com/library).
- **Any language** — plain HTTPS + JSON, with examples for Python, TypeScript, cURL, PHP, C#, Go, Java, Ruby and Node.js.

## FAQ

### Is it legal to scrape ManoMano?

Collecting publicly available data is generally allowed in many jurisdictions, but rules depend on your country, the data you collect and how you use it. Avoid personal data unless you have a lawful basis, respect the website's terms where they apply, and get legal advice for your specific use case. This is not legal advice.

### Do I need proxies or a headless browser?

No. Piloterr handles rendering, proxy rotation and retries on its side. You only send an HTTPS request with your `x-api-key` header and receive JSON.

### How much does a request cost?

Credits per call: `/v2/manomano/product` **1**, `/v2/manomano/product/offer` **1**, `/v2/manomano/search` **1**. See the [Piloterr website](https://www.piloterr.com) for plans.

### What data does the ManoMano API return?

Each endpoint returns structured JSON: **ManoMano Product** — `ean`, `brand`, `brand.id`, `brand.url`, `brand.name`, `media[]`; **ManoMano Product Offer** — `count`, `results[]`, `results[].sku`, `results[].market`, `results[].seller`, `results[].pricing`; **ManoMano Search** — `results[]`, `results[].cpc`, `results[].url`, `results[].bulk`, `results[].title`, `results[].is_b2b`.

### Can I target a specific country or language?

Yes, use `region` — see the parameter tables above for accepted values.

### Which languages are supported?

Any language that can make an HTTP request. This README shows Python and TypeScript; the [`examples/`](examples) folder contains cURL, PHP, C#, Go, Java, Ruby and Node.js.

## Related scrapers

- [About You Scraper API](https://github.com/piloterr/aboutyou-scraper)
- [Action Scraper API](https://github.com/piloterr/action-scraper)
- [Aldi Scraper API](https://github.com/piloterr/aldi-scraper)
- [Alibaba Scraper API](https://github.com/piloterr/alibaba-scraper)
- [AliExpress Scraper API](https://github.com/piloterr/aliexpress-scraper)
- [Allegro Scraper API](https://github.com/piloterr/allegro-scraper)
- [All Piloterr scrapers on GitHub](https://github.com/piloterr)

## Support

- 💬 Join the [Piloterr Discord](https://discord.gg/CFj3ph8aHU) to ask questions and share what you build.
- 📚 Browse the [API reference (OpenAPI)](https://app.piloterr.com/openapi) and the [Piloterr library](https://www.piloterr.com/library).
- 🔑 [Get your API key](https://app.piloterr.com).

---

<p align="center">
  <a href="https://www.piloterr.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="images/logo-dark.svg">
      <img src="images/logo.svg" alt="Piloterr" height="32">
    </picture>
  </a>
  <br>
  <sub>Built by <a href="https://www.piloterr.com">Piloterr</a> — web scraping APIs for e-commerce, search, real estate and more, one API key.</sub>
</p>

<sub>ManoMano is a trademark of its respective owner. This repository is not affiliated with or endorsed by ManoMano. Examples are released under the MIT license.</sub>
