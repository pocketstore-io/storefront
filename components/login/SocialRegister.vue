<template>
  <section class="actions">
    <div class="grid grid-cols-6 gap-3">
      <div class="col-span-6 md:col-span-3">
        <button class="btn btn-block" @click="loginViaGoogle()">
          <Fa :icon="faGoogle" />
          <span>{{ $t('company.google') }}</span>
        </button>
      </div>
      <div class="col-span-6 md:col-span-3">
        <button class="btn btn-block" @click="loginViaGithub()">
          <Fa :icon="faGithub" />
          <span>{{ $t('company.github') }}</span>
        </button>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { faGithub, faGoogle } from "@fortawesome/free-brands-svg-icons";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();

const loginViaGoogle = async function () {
    const authData = await pb
        .collection("customers")
        .authWithOAuth2({ provider: "google" });
    console.log(authData);
};
const loginViaGithub = async function () {
    const authData = await pb
        .collection("customers")
        .authWithOAuth2({ provider: "github" });
    console.log(authData);
};
</script>