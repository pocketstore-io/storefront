<template>
  <div class="alert border-black alert-outline mb-3 text-black">
    <label class="label" for="same">
      <input type="checkbox" v-model="same" checked="checked" class="toggle" />
      <span>{{ $t('checkout.same') }}</span>
    </label>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const same = useLocalStorage("same", true, {});
</script>