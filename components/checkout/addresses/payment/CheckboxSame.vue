<template>
  <div class="alert border-black alert-outline mb-3">
    <label class="label" for="same">
      <span>{{ $t('checkout.same') }}</span>
      <input v-model="same" id="same" type="checkbox" checked="checked" class="checkbox" />
    </label>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const same = useLocalStorage("same", true, {});
</script>