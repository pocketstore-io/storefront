<template>
  <section class="bg-gray-300 min-h-screen">
    <header>
      <mainHeader />
      <div class="wave" />
    </header>
    <main class="mx-auto max-w-5xl min-h-screen px-3 md:px-0">
      <section class="alert alert-error flex mb-3">
        <Fa :icon="faInfoCircle" class="text-white" />
        <p class="text-sm font-bold w-full text-center">
          {{ $t('note') }}
        </p>
        <Fa :icon="faInfoCircle" class="text-white" />
      </section>
      <Messages />
      <Breadcrumb />
      <CookieBanner />
      <router-view />
    </main>
    <footer class="">
      <div class="zigzag" />
      <mainFooter />
    </footer>
  </section>
</template>

<script setup lang="ts">
import { faInfoCircle } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome';
import './main.css';
import tracking from '~/util/tracking';
import { useBreadcrumbStore } from '~/stores/breadcrumb';

const store = useBreadcrumbStore();

onMounted(() => {
  tracking.trackPage()
  store.clear()
});
</script>

<style>
.wave {
  --size: 40px;
  --R: calc(var(--size)*1.28);
  -webkit-mask:
    radial-gradient(var(--R) at 50% calc(1.8*var(--size)), #000 99%, #0000 101%) calc(50% - 2*var(--size)) 0/calc(4*var(--size)) 100%,
    radial-gradient(var(--R) at 50% calc(-.8*var(--size)), #0000 99%, #000 101%) 50% var(--size)/calc(4*var(--size)) 100% repeat-x;
  background: #1f6fed;
  height: 80px;
  rotate: 180deg;
}

.zigzag {
  height: 20px;
  width: 100%;
  background:
    linear-gradient(-135deg, transparent 5px, transparent 0) 0 5px,
    linear-gradient(135deg, transparent 5px, #fff 0) 0 5px;
  background-color: transparent;
  background-position: left bottom, left bottom;
  background-repeat: repeat-x;
  background-size: 10px 10px, 10px 10px;
}
</style>