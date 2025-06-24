<template>
  <div class="grid grid-cols-7 gap-3">
    <div class="col-span-7 md:col-span-4 bg-white px-3 py-3">
      <p>
        Faucibus commodo massa rhoncus, volutpat. Dignissim sed eget risus enim.
        Mattis mauris semper sed amet vitae sed turpis id. Id dolor praesent
        donec est. Odio penatibus risus viverra tellus varius sit neque erat
        velit. Faucibus commodo massa rhoncus, volutpat. Dignissim sed eget
        risus enim. Mattis mauris semper sed amet vitae sed turpis id.
      </p>
      <ul role="list" class="mt-8 space-y-8 text-gray-600">
        <li class="flex gap-x-3">
          <FontAwesomeIcon :icon="faLock" />
          <span
            ><strong class="font-semibold text-gray-900"
              >Push to deploy.</strong
            >
            Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores
            impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis
            ratione.</span
          >
        </li>
        <li class="flex gap-x-3">
          <FontAwesomeIcon :icon="faLock" />
          <span
            ><strong class="font-semibold text-gray-900"
              >Push to deploy.</strong
            >
            Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores
            impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis
            ratione.</span
          >
        </li>
        <li class="flex gap-x-3">
          <FontAwesomeIcon :icon="faLock" />
          <span
            ><strong class="font-semibold text-gray-900"
              >Push to deploy.</strong
            >
            Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores
            impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis
            ratione.</span
          >
        </li>
      </ul>
    </div>
    <div class="col-span-7 md:col-span-3 bg-white px-3 py-6">
      <h3 class="font-bold text-lg mt-5 mb-12">Jetzt Registrieren</h3>
      <form @submit.prevent="register()" ref="form-ref" class="space-y-5">
        <label class="floating-label">
          <span>Your Name</span>
          <input
            type="text"
            v-model="form.name"
            class="input input-bordered input-primary w-full"
          />
        </label>
        <label class="floating-label">
          <span>Your Email</span>
          <input
            v-model="form.email"
            type="email"
            class="input input-bordered input-primary w-full"
          />
        </label>
        <label class="floating-label">
          <span>Your Password</span>
          <input
            type="password"
            v-model="form.password"
            class="input input-bordered input-primary w-full"
          />
        </label>
        <label class="floating-label">
          <span>Your Password</span>
          <input
            type="password"
            required
            v-model="form.passwordConfirm"
            class="input input-bordered input-primary w-full"
          />
        </label>

        <section class="actions flex justify-end">
          <button type="submit" :disabled="!valid" class="btn btn-primary btn-sm">
            Registrieren
          </button>
        </section>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faLock } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useRouter } from "vue-router";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const router = useRouter();
const formRef = useTemplateRef('form-ref');

const form = ref({
  name: "",
  email: "",
  password: "",
  passwordConfirm: "",
});

const register = async () => {
  await pb.collection("customers").create(form.value);
  await pb.collection('customers').authWithPassword(form.value.email,form.value.password);
  if (pb.authStore.isValid) {
    router.push("/");
  }
};

const valid = computed(()=>{
  return formRef.value.checkValidity()
});

onMounted(() => {
  if (pb.authStore.isValid) {
    router.push("/");
  }
});
</script>
