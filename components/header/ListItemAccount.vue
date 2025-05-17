<template>
  <li v-if="isLoggedIn" class="mt-1">
    <a href="/user/logout">
      <Fa :icon="faRightFromBracket" class="text-white mx-2 mt-1" />
    </a>
  </li>
  <li v-else class="mt-1">
    <a href="/login">
      <Fa :icon="faUser" class="text-white mx-2 mt-1" />
    </a>
  </li>
</template>

<script lang="ts" setup>
import { faRightFromBracket, faUser } from "@fortawesome/free-solid-svg-icons";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const isLoggedIn = computed(() => {
    return pb.authStore.isValid == true;
});
</script>