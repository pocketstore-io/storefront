<template>
    <section class="bg-gray-300 min-h-screen">
        <header>
            <Header />
        </header>
        <main class="mx-auto max-w-6xl min-h-screen px-3 md:px-0">
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
import '~/main.css';
import Footer from '~/components/Footer.vue';
import Header from '~/components/Header.vue';
import BackToTop from '~/components/BackToTop.vue';
</script>

<style>
@import url('layout.css');
</style>