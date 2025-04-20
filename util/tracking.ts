export default {
    async trackPage() {
        const { proxy } = useScriptPlausibleAnalytics({domain: 'demo.pocketstore.io'})
        proxy.plausible('event', { name: 'page load' })
    }
}