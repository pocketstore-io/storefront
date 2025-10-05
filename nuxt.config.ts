import config from "./pocketstore.json";
// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";
export default defineNuxtConfig({
    devtools: { enabled: true },
    css: ["@fortawesome/fontawesome-svg-core/styles.css"],
    modules: ["@nuxtjs/i18n"],
    plausible: {
        // Prevent tracking on localhost
        ignoredHostnames: [
            // 'localhost'
        ],
        apiHost: "https://tracking.jmse.cloud",
    },
    i18n: {
        strategy: "prefix",
        locales: config.language.languages,
        defaultLocale: "de",
    },
    nitro: {
        routeRules: {
            "/product/**": { static: true, redirect: (to) => `${to}.html` },
            "/category/**": { static: true, redirect: (to) => `${to}.html` },
            "/brand/**": { static: true, redirect: (to) => `${to}.html` },
        },
    },
    app: {
        head: {
            link: [
                {
                    rel: "manifest",
                    href: "/manifest.webmanifest",
                },
                {
                    rel: "icon",
                    type: "image/svg+xml",
                    href: "/pocketstore-favicon.svg",
                },
            ],
        },
    },
    vite: {
        plugins: [tailwindcss()],
    },
    compatibilityDate: "2025-06-21",
});
