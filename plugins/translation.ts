/* eslint-disable */
import config from '../pocketstore.json';
import PocketBase from 'pocketbase'

export default defineNuxtPlugin((nuxtApp) => {
    // nuxtApp.$i18n.setMissingHandler(async (locale, message) => {
    //     const pb = new PocketBase(config.domain);
    //     await pb.collection('translations').create({
    //         lang: locale,
    //         key: message,
    //         translated: 'Missing Translation'
    //     });
    // });
})
