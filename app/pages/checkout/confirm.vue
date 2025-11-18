<template>
  <section class="page bg-red-400">
    <div class="hero bg-base-200 py-44">
      <div class="hero-content text-center">
        <div class="max-w-xl">
          <h1 class="text-2xl font-bold">Herzlichen Dank für ihre Bestellung</h1>
          <p class="py-6">
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
            Hier könnte dein BlaBla stehen in dem Fall das du einen PocketStore für dich aufsetzt.
          </p>
          <section class="actions flex justify-center space-x-3">
            <a href="/de" class="btn btn-primary">Weiter einkaufen</a>
            <a :href="'/customer/invoice/' + route.query.order" class="btn btn-secondary">Zur Rechnung</a>
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
