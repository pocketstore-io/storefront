import pocketstoreConfig from "../pocketstore.json";
export default defineNuxtPlugin(() => {
    useHead({
        script: [
            {
                async: true,
                defer: true,
                src:
                    "https://consent.cookiefirst.com/sites/pocketstore-" +
                    pocketstoreConfig.integrations.cookiefirst.id +
                    "/consent.js",
            },
        ],
    });
});
