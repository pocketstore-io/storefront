import config from './pocketstore.json'
import tailwindcss from '@tailwindcss/vite'
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  css: ['@fortawesome/fontawesome-svg-core/styles.css'],
  plausible: {
    // Prevent tracking on localhost
    ignoredHostnames: [
      // 'localhost'
    ],
    apiHost: 'https://tracking.jmse.cloud'
  },
  // TODO remove pinia
  modules: ['@nuxtjs/i18n', '@vite-pwa/nuxt', '@nuxtjs/plausible'],
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
  },vite: {
    plugins: [
      tailwindcss(),
    ],
  },
  compatibilityDate: '2024-10-19',
})