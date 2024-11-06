import PocketBase from 'pocketbase';

// TODO get pocketbase url from Nuxt Config ?
const pb = new PocketBase('https://admin.pocketstore.io/');

export default {
    async trackPage() {
        if (location.host != 'localhost:3000') {
            let lc = localStorage.getItem('identifier');

            if (lc === null) {
                function randomString(length) {
                    const chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz'.split('');

                    if (!length) {
                        length = Math.floor(Math.random() * chars.length);
                    }

                    let str = '';
                    for (let i = 0; i < length; i++) {
                        str += chars[Math.floor(Math.random() * chars.length)];
                    }
                    return str;
                }

                lc = randomString(8);
                localStorage.setItem('identifier', lc)
            }

            await pb.collection('tracking_pages').create({
                uri: location.pathname,
                customer: pb.authStore.model?.id,
                identifier: lc,
                data: {
                    platform: navigator.platform,
                    referer: document.referrer
                }
            });
        }

    }
}