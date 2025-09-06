<template>
  <div class="bg-gray-400 mx-auto max-w-md rounded-xl my-24">
    <form @submit.prevent="login()" class="w-full space-y-5 px-3 py-6">
      <h2 class="font-bold text-lg">Login</h2>
      <label class="floating-label">
        <span>Your Email</span>
        <input
          v-model="identity"
          type="email"
          required
          class="input w-full input-bordered input-primary"
        />
      </label>
      <label class="floating-label">
        <span>Your Password</span>
        <input
          type="password"
          v-model="password"
          required
          class="input w-full input-bordered input-primary"
        />
      </label>
      <seciton class="actions flex justify-end">
        <button type="submit" class="btn btn-primary btn-sm">login</button>
      </seciton>
    </form>
  </div>
</template>

<script lang="ts" setup>
import { usePocketBase } from "~/util/pocketbase";

const identity = ref("");
const password = ref("");

const pb = usePocketBase();
const router = useRouter();

const login = async () => {
  await pb
    .collection("customers")
    .authWithPassword(identity.value, password.value);

  if (pb.authStore.isValid) {
    router.push("/de");
  }
};

onMounted(() => {
  if (pb.authStore.isValid) {
    router.push("/de");
  }
});
</script>
