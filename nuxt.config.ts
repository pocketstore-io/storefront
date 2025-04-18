import config from './pocketstore.json'
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  css: ['@fortawesome/fontawesome-svg-core/styles.css'],

  // TODO remove pinia
  modules: ['@nuxtjs/i18n', '@vite-pwa/nuxt'],
  i18n: {
    strategy: 'prefix',
    // TODO allow custom lang
    locales: [
      {
        code: 'de',
        file: 'de.json'
      }
    ],
    defaultLocale: 'de'
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  nitro: {
    routeRules: {
      // TODO custom rules by pocketstore
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