// ManoMano Product API – Node.js example (Node 18+, native fetch, ES module)
// Usage: PILOTERR_API_KEY=your-key node product.mjs

const params = new URLSearchParams({
  "query": "https://www.manomano.fr/p/pergola-retractable-298l-x-213l-x-222h-m-structure-metal-epoxy-anticorrosion-toile-polyester-haute-densite-180-g-m-incluse-gris-17786586?model_id=17784588",
});

const response = await fetch(`https://api.piloterr.com/v2/manomano/product?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY },
  signal: AbortSignal.timeout(120_000),
});

if (!response.ok) {
  console.error(`Piloterr error ${response.status}: ${await response.text()}`);
  process.exit(1);
}

console.log(JSON.stringify(await response.json(), null, 2));
