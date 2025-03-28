// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  css: ['@fortawesome/fontawesome-svg-core/styles.css'],
  modules: ['@pinia/nuxt', '@nuxtjs/i18n', '@vite-pwa/nuxt'],
  pinia: {
    storesDirs: ['./stores/**'],
  },
  i18n: {
    strategy: 'prefix',
    locales: [
      {
        code: 'de',
        file: 'de.json'
      },
      {
        code: 'en',
        file: 'en.json'
      }
    ],
    defaultLocale: 'en'
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  nitro: {
    routeRules: {
      // Only apply .html suffix to routes starting with "product/"
      '/product/**': { static: true, redirect: (to) => `${to}.html` },
      '/category/**': { static: true, redirect: (to) => `${to}.html` }
    }
  },
  app: {
    head: {
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/android-chrome-192x192.png' }
      ]
    }
  },
  compatibilityDate: '2024-10-19',
})