// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['@fortawesome/fontawesome-svg-core/styles.css'],
  modules: ['@pinia/nuxt', '@nuxtjs/i18n','nuxt-vue3-google-signin'],
  googleSignIn: {
    clientId: '897877024261-d84fc3v63gqvcujvb1b8alqhv40i14en.apps.googleusercontent.com',
  },
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
      '/product/**': { static: true, redirect: (to) => `${to}.html`}
    }
  },
  compatibilityDate: '2024-10-19',
})