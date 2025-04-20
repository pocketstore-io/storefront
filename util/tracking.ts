export default {
    async trackPage() {
        const { proxy } = useScriptPlausibleAnalytics({
            domain: 'demo.pocketstore.io',
            scriptInput: {
                src: 'https://tracking.jmse.cloud/js/script.outbound-links.pageview-props.revenue.tagged-events.js'
            }
        })
        proxy.plausible('event', { name: 'page load' })
    }
}