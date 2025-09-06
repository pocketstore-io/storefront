<template>
  <ModalBackToTop />
  <transition>
    <section v-if="show"
      class="back-to-top-icon fixed bottom-0 right-0 mr-2 mb-3 bg-black text-white rounded-full h-16 w-16 flex items-center justify-center z-10"
      @click="scrollToTop()">
      <Fa :icon="faArrowUp" size="2x" />
    </section>
  </transition>
</template>

<script lang="ts" setup>
import { faArrowUp } from "@fortawesome/free-solid-svg-icons";
import { useLocalStorage } from "@vueuse/core";

const open = useLocalStorage("open", false, {});
const show = ref(false);
const cookie = useLocalStorage("cookie", null, {});

const scrollToTop = () => {
    window.scrollTo({ top: 0, behavior: "smooth" });
};

onMounted(() => {
    if (cookie.value === null) {
        open.value = true;
    }
    window.addEventListener("scroll", () => {
        if (document.documentElement.scrollTop > 250) {
            show.value = true;
        } else {
            show.value = false;
        }
    });
});
</script>