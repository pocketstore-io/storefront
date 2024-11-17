import { defineNuxtPlugin, useRoute } from '#app';

export default defineNuxtPlugin((nuxtApp) => {
    const route = useRoute();
    nuxtApp.provide('layout', route);
});