export default {
    async trackPage() {
        if (typeof window !== 'undefined' && window.plausible) {
            window.plausible('CustomEventName', {
              props: {
                key1: 'value1',
                key2: 'value2'
              }
            });
          } else {
            console.warn('Plausible is not initialized');
          }
    }
}