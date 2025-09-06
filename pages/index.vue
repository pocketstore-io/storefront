<template>
  <section class="page">
    <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
      <HomepageHero />
      <CatalogProductGrid identifier="bowling" />
    </div>
  </section>
</template>

<script setup lang="ts">
import config from './../pocketstore.json'
import {addBreadcrumb} from "~/util/breadcrumbs";
const { t } = useI18n();
useHead({
  title: t('general.title'),
});

onMounted(() => {
  addBreadcrumb({
    label: 'test',
    link: 'test123',
    icon: 'question'
  })
})
</script>