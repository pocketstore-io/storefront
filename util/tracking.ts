export default {
    async trackPage() {
        const { proxy } = useScriptPlausibleAnalytics()
        proxy.plausible('event', { name: 'page load' })
    }
}