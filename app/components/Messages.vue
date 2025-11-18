<template>
  <Disclaimer />
  <section class="messages space-y-3 mb-3">
    <section v-for="message in messages">
      <section v-if="time < message.timer" class="alert alert-success text-white">
        {{ message.message }}
      </section>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import type { ToastMessage } from "~/utils/toast";
const messages = useLocalStorage("messages", [], {});
const time = ref(Math.floor(Date.now() / 1000));

onMounted(() => {
    setInterval(() => {
        time.value = Math.floor(Date.now() / 1000 + 1);
        messages.value.map((item: ToastMessage, index) => {
            if (time.value < item.timer) {
                messages.value = messages.value.splice(index, 1);
            }
        });
    }, 1000);
});
</script>