<template>
  <h3 v-if="props.item.price" class="font-bold text-4xl">
    <div class="glass px-3 py-3 mt-6 mb-6 md:flex md:justify-between">
      <section class="text-center space-x-3">
        <button class="btn btn-warning hover:text-white">
          <Fa :icon="faStar" />
        </button>
        <button class="btn btn-warning hover:text-white">
          <Fa :icon="faStar" />
        </button>
        <button class="btn btn-warning hover:text-white">
          <Fa :icon="faStar" />
        </button>
        <button class="btn btn-warning hover:text-white">
          <Fa :icon="faStar" />
        </button>
        <button class="btn btn-warning hover:text-white">
          <Fa :icon="faStar" />
        </button>
      </section>
      <div class="mt-6 md:mt-2 text-right md:text-left">
        {{ props.item.price.toFixed(2) }} €
      </div>
    </div>
  </h3>
</template>

<script lang="ts" setup>
import { faStar } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";

const props = defineProps({
  item: {
    required: true,
    type: Object,
  },
});
</script>
