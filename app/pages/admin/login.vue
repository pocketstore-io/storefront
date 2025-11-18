<template>
  <!-- Put this part before </body> tag -->
  <input type="checkbox" v-model="open" id="my_modal_7" class="modal-toggle" />
  <div class="modal" role="dialog">
    <div class="modal-box">
      <h3 class="text-lg font-bold">SuperAdmin Login?</h3>
      <form @submit.prevent="login()" class="py-3">
        <label class="floating-label mt-3">
          <span>Your Username</span>
          <input
            v-model="identity"
            type="email"
            class="input input-bordered input-primary w-full"
          />
        </label>
        <label class="floating-label mt-3">
          <span>Your Password</span>
          <input
            v-model="password"
            type="password"
            class="input input-bordered input-primary w-full"
          />
        </label>
        <section class="actions flex justify-end mt-3">
          <button type="button" class="btn btn-secondary btn-sm hidden">
            register
          </button>
          <button type="submit" class="btn btn-primary btn-sm">save</button>
        </section>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import "@/main.css";
import { usePocketBase } from "~/utils/pocketbase";

const open = ref(true);
const identity = ref("");
const password = ref("");
const pb = usePocketBase();
const router = useRouter();

const login = async () => {
    await pb
        .collection("_superusers")
        .authWithPassword(identity.value, password.value);
    if (pb.authStore.isSuperuser) {
        router.push("/de/admin");
    }
};

definePageMeta({
    layout: "login",
});
</script>