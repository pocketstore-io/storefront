/* eslint-disable */
import config from "../pocketstore.json";
import PocketBase from "pocketbase";

// TODO upload translation file and write with plugin to translation file

export default defineNuxtPlugin((nuxtApp) => {
    // nuxtApp.$i18n.setMissingHandler(async (locale, message) => {
    //     const pb = new PocketBase(config.domains.pocketbase);
    //     await pb.collection('translations').create({
    //         lang: locale,
    //         key: message,
    //         translated: 'Missing Translation'
    //     });
    // });
});
