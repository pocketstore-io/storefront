<template>
  <div class="grid grid-cols-6 gap-3">
    <div class="col-span-6 md:col-span-3">
      <FormLogin />
    </div>
    <div class="col-span-6">
      <LoginHero />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const router = useRouter();

onMounted(() => {
    if (pb.authStore.isValid) {
        router.push("/");
    }
});
</script>
