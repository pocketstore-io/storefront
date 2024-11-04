// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['@fortawesome/fontawesome-svg-core/styles.css'],
  modules: ['@pinia/nuxt', '@nuxtjs/i18n'],
  pinia: {
    storesDirs: ['./stores/**'],
  },
  i18n: {
    langDir: 'configuration/lang/',
    locales: [
      {
        code: 'de',
        file: 'lang_de.json'
      },
      {
        code: 'en',
        file: 'lang_en.json'
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