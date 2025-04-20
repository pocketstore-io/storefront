<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
const { proxy } = useScriptPlausibleAnalytics()

if (navigator && 'serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker
      .register('/service-worker.js')
      .then((registration) => {
        console.log('Service Worker registered with scope:', registration.scope);
      })
      .catch((error) => {
        console.error('Service Worker registration failed:', error);
      });
  });
}

onMounted(() => {
  useScriptPlausibleAnalytics({
    scriptInput: {
      src: 'https://tracking.jmse.cloud/js/script.outbound-links.pageview-props.revenue.tagged-events.js'
    }
  })
});
</script>