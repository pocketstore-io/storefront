<template>
  <div>
    <div v-if="!props.quantity">
      <section class="alert alert-error flex justify-end">
        <span>{{  $t('product.stock.empty') }}: </span>
        <b>0</b>
      </section>
    </div>
    <div v-else-if="props.quantity <= 10">
      <section class="alert alert-warning flex justify-end">
        <span>{{ $t('product.stock.warning') }}</span>:
        <b>{{ props.quantity }}</b>
      </section>
    </div>
    <div v-else>
      <section class="alert alert-success flex justify-end">
         <span>{{ $t('product.stock.full') }}</span>:
        <b>{{ props.quantity }}</b>
      </section>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  quantity: {
    required:true,
    type: Number
  },
});
</script>