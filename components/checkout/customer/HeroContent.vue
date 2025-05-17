<template>
  <div v-if="!pb.authStore.isValid" class="text-center lg:text-left">
    <h1 class="text-5xl font-bold">Login now!</h1>
    <p class="py-6">
      Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
      quasi. In deleniti eaque aut repudiandae et a id nisi.
    </p>
  </div>
  <div v-else class="text-center lg:text-left">
    <h1 class="text-2xl font-bold">Deine Vorteile als Kunde</h1>
    <p class="py-6">
      Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
      quasi. In deleniti eaque aut repudiandae et a id nisi.
    </p>
  </div>
</template>

<script lang="ts" setup>
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
</script>
