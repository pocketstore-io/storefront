<template>
  <section class="bg-gray-600">
    <div
      class="grid grid-cols-7 px-3 py-3 text-center mx-auto max-w-5xl text-white"
    >
      <div class="col-span-7 md:col-span-2">
        <NuxtLink :to="'/de/impressum'" class="text-blue-400">{{
          $t("footer.impressum")
        }}</NuxtLink>
        &
        <a href="/de/data-privacy" class="text-blue-400">{{
          $t("footer.data-privacy")
        }}</a>
      </div>
      <div class="col-span-7 md:col-span-2">
        {{ $t("footer.made-with") }}
        <a
          :href="'https://' + url.value"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Fa :icon="faHeart" class="text-red-400" />
        </a>
        by
        <a
          class="text-primary"
          target="_blank"
          href="https://www.jonathan-martz.de"
          >Jonathan</a
        >
      </div>
      <div class="col-span-7 md:col-span-3">
        Copyright by
        <a
          class="text-primary"
          target="_blank"
          href="https://www.jonathan-martz.de"
          >Jonathan Martz Softwareentwicklung</a
        >
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { usePocketBaseUrl } from "~/util/pocketbase";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";
import { faHeart } from "@fortawesome/free-solid-svg-icons";

const url = usePocketBaseUrl();
</script>
