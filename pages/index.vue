<script setup>
import { useBreadcrumbStore } from '~/stores/breadcrumb';

const storeBreadcrumb = useBreadcrumbStore();
storeBreadcrumb.clear();

useHead({
    script: [
        {
            async: true,
            defer: true,
            src: 'https://consent.cookiefirst.com/sites/demo.pocketstore.io-78a0f1a2-ac20-421b-81cf-5dc77f713b67/consent.js'
        },
    ],
});
</script>

<template>
    <section class="page">
        <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
            <HomepageHero />
            <CatalogProductGrid category-name="Welcome" />
        </div>
    </section>
</template>