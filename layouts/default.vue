<template>
    <section class="page min-h-screen" data-theme="light">
        <header>
            <Header />
        </header>
        <main class="mx-auto max-w-6xl min-h-screen px-3 xl:px-0">
            <Messages />
            <Breadcrumb />
            <CookieBanner />
            <BackToTop />
            <slot />
        </main>
        <footer>
            <Footer />
        </footer>
    </section>
</template>
<script setup lang="ts">
import "../main.css";
import Footer from "~/components/Footer.vue";
import Header from "~/components/Header.vue";
</script>

<style scoped>
.page {
  background: radial-gradient(
      circle at top,
      transparent 9%,
      #1EADEF 10%,
      #1EADEF 15%,
      transparent 16%
    ),
    radial-gradient(
      circle at bottom,
      transparent 19%,
      #1EADEF 20%,
      #1EADEF 25%,
      transparent 26%
    ),
    radial-gradient(
      circle at right,
      transparent 19%,
      #1EADEF 20%,
      #1EADEF 25%,
      transparent 26%
    ),
    radial-gradient(
      circle at left,
      transparent 9%,
      #1EADEF 10%,
      #1EADEF 15%,
      transparent 16%
    );
  background-size: 3em 3em;
  background-color: #ccc;
  opacity: 1;
}
</style>

