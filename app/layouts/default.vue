<template>
  <section v-if="!config.maintance.enabled" class="page min-h-screen" data-theme="light">
    <header>
      <Header/>
    </header>
    <main class="mx-auto max-w-6xl min-h-screen px-3 2xl:px-0">
      <Messages/>
      <Breadcrumb/>
      <CookieBanner/>
      <BackToTop/>
      <BugReporting />
      <IssueReporting />
      <slot/>
    </main>
    <footer>
      <Footer/>
    </footer>
  </section>
  <section v-else class="bg-gray-400 min-h-screen w-full px-3 py-3 flex items-center justify-center">
    <div class="card shadow-sm max-w-md bg-gray-300">
      <figure>
        <img
            src="/unsplash/christopher-gower-m_HRfLhgABo-unsplash.jpg"
            alt="Shoes"/>
      </figure>
      <div class="card-body">
        <h2 class="card-title">Dieser Shop ist aktuell im Wartungsmodus</h2>
        <p>
          Wir warten gerade unserm Shop System von
          <a href="https://PocketStore.io" class="text-blue-400">PocketStore.io</a>
          für noch bessere Einkaufsmöglichkeiten und Support
          unserer Kunden.
        </p>
        <div class="card-actions justify-end mt-3">
          <p class="block text-center alert alert-info text-white">
            Wie gewohnt ab erreichbar ab {{ config.maintance.until }}
          </p>
          <button class="btn btn-block btn-secondary">Benachrichtigung per Mail</button>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import "../main.css";
import config from "~/pocketstore.json";
import Footer from "~/components/Footer.vue";
import Header from "~/components/Header.vue";
import BugReporting from "~/components/BugReporting.vue";
import IssueReporting from "~/components/IssueReporting.vue";
</script>

<style scoped>
.page {
  background: radial-gradient(
      circle at top,
      transparent 9%,
      #1EADEF 10%,
      #1EADEF 15%,
      transparent 16%
  ),
  radial-gradient(
      circle at bottom,
      transparent 19%,
      #1EADEF 20%,
      #1EADEF 25%,
      transparent 26%
  ),
  radial-gradient(
      circle at right,
      transparent 19%,
      #1EADEF 20%,
      #1EADEF 25%,
      transparent 26%
  ),
  radial-gradient(
      circle at left,
      transparent 9%,
      #1EADEF 10%,
      #1EADEF 15%,
      transparent 16%
  );
  background-size: 3em 3em;
  background-color: #ccc;
  opacity: 1;
}
</style>

