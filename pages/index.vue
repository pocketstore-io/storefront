<template>
  <section class="page">
    <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
      <HomepageHero />
      <CatalogProductGrid identifier="bowling" />
    </div>
  </section>
</template>

<script setup lang="ts">
import {addBreadcrumb} from "~/util/breadcrumbs";
const { t } = useI18n();
useHead({
    title: t("general.title"),
    meta: [
      {
        name: 'description',
        content: 'PocketStore.io is a Open Source Shop Framework which is based on PocketBase and made with Nuxt/VueJs by Jonathan Martz.'
      }
    ]
});

onMounted(() => {
  addBreadcrumb({
    label: 'test',
    link: 'test123',
    icon: 'question'
  })
})
</script>