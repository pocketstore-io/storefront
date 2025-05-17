<template>
  <section class="page bg-red-400">
    <div class="hero bg-base-200 py-44">
      <div class="hero-content text-center">
        <div class="max-w-xl">
          <h1 class="text-2xl font-bold">Herzlichen Dank für ihre Bestelleung</h1>
          <p class="py-6">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugiat voluptates ducimus molestias eaque maiores
            consectetur dicta! Dolorum iste accusantium repudiandae quaerat perferendis, iusto reiciendis tempore. Ipsam
            ea tenetur molestias quae.
          </p>
          <section class="actions flex justify-center space-x-3">
            <a href="/en" class="btn btn-primary">Weiter einkaufen</a>
            <a :href="'/customer/invoice/' + route.query.order" class="btn btn-primary">Zur Rechnung</a>
          </section>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();

onMounted(() => {
    if (typeof route.query.order == "undefined") {
        router.push("/de");
    }
});
</script>
