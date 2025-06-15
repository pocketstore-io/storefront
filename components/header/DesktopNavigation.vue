<template>
  <nav class="hidden md:flex items-center text-white">
    <ul class="flex space-x-6 text-sm font-bold items-center mr-6">
      <li>Startseite</li>
      <li>
        <a href="https://www.pocketstore.io/" target="_blank">Docs</a>
      </li>
      <li>
        <a href="https://www.pocketstore.io/shops/"> Shops </a>
      </li>
    </ul>
    <ul class="flex space-x-3">
      <li>
        <a
          href="https://github.com/pocketstore-io/demo"
          target="_blank"
          class="flex justify-center items-center btn btn-primary text-white"
        >
          <FontAwesomeIcon size="2x" class="w-5" :icon="faGithub" />
        </a>
      </li>
      <li>
        <a
          href="/de/checkout"
          class="flex justify-center items-center btn btn-neutral"
        >
          <FontAwesomeIcon size="1x" :icon="faShoppingCart" />
          <span class="ml-2">0</span>
        </a>
      </li>
      <li>
        <button class="flex justify-center items-center btn btn-neutral">
          <FontAwesomeIcon size="1x" :icon="faHeart" color="red" />
        </button>
      </li>
    </ul>
  </nav>
</template>


<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faShoppingCart, faHeart } from "@fortawesome/free-solid-svg-icons";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
</script>