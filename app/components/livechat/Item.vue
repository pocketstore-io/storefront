<template>
  <div class="chat" :class="{ 'chat-start': props.item.from == 'from', 'chat-end': props.item.from != 'from' }">
    <div class="chat-bubble"
      :class="{ 'chat-bubble-primary': props.item.from == 'from', 'chat-bubble-neutral': props.item.from != 'from' }"
      @click="beep()">
      {{ props.item.message }}
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
    item: { type: Object, required: true },
});

const beep = () => {
    const snd = new Audio("/sounds/bling.mp3");
    snd.play();
};
</script>