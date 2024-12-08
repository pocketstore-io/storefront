<template>
  <div class="px-3">
    <section class="alert alert-error flex mb-3 xl:mx-auto">
      <Fa :icon="faInfoCircle" class="text-white" />
      <p class="text-sm font-bold w-full text-center">
        {{ $t('general.note') }}
      </p>
      <Fa :icon="faInfoCircle" class="text-white" />
    </section>
  </div>
</template>

<script lang="ts" setup>
import { faInfoCircle } from '@fortawesome/free-solid-svg-icons';
</script>