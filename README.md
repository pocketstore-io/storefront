# PocketStore.io - Storefront

## Setup

Make sure to install the dependencies:

```bash
# npm
npm install

# bun
bun install
```

## Development Server

Start the development server on `http://localhost:3000`:

```bash
# npm
npm run dev

# bun
bun run dev
```

## Production

Build the application for production:

```bash
# npm
npm run build
npx pm2 start ecosystem.config.cjs

# bun
bun run build
bun x pm2 start ecosystem.config.cjs
```

Available under localhost:4000 or your Domain

