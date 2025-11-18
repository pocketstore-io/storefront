// plugins/i18n.client.ts
import { createI18n } from 'vue-i18n'

export default defineNuxtPlugin(nuxtApp => {
    const i18n = createI18n({ legacy: false, locale: 'en', messages: {} })
    nuxtApp.vueApp.use(i18n)
})