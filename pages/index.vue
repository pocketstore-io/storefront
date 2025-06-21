<template>
  <section class="page">
    <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
      <HomepageHero />
      <CatalogProductGrid identifier="bowling" />
    </div>
  </section>
</template>

<script setup lang="ts">
import config from "./../pocketstore.json";
const { t } = useI18n();
useHead({
    title: t("general.title"),
});
</script>