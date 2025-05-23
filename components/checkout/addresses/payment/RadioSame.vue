<template>
  <div class="alert border-black alert-outline mb-3">
    <label class="label" for="same">
      <input v-model="same" id="same" type="checkbox" checked="checked" class="checkbox input" />
      <span>{{ $t('checkout.same') }}</span>
    </label>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const same = useLocalStorage("same", true, {});
</script>