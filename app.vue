<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
onMounted(() => {
    if (navigator && "serviceWorker" in navigator) {
        window.addEventListener("load", () => {
            navigator.serviceWorker
                .register("/service-worker.js")
                .then((registration) => {
                    console.log(
                        "Service Worker registered with scope:",
                        registration.scope,
                    );
                })
                .catch((error) => {
                    console.error("Service Worker registration failed:", error);
                });
        });
    }
});
</script>