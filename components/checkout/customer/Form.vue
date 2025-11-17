<template>
    <form class="card-body">
        <div class="form-control">
            <label class="label">
                <span class="label-text">Email</span>
            </label>
            <input v-model="email" type="email" placeholder="email" class="input input-bordered" required>
        </div>
        <div class="form-control">
            <label class="label">
                <span class="label-text">Password</span>
            </label>
            <input v-model="password" type="password" placeholder="password" class="input input-bordered" required>
            <label class="label">
                <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
            </label>
        </div>
        <div class="form-control mt-6">
            <button class="btn btn-primary" @click.prevent="login()">Login</button>
        </div>
    </form>
</template>

<script setup lang="ts">
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/utils/pocketbase";

const pb = usePocketBase();
const checkoutStep = useLocalStorage("checkoutStep", "cart", {});

const email = ref("");
const password = ref("");

const login = async () => {
    await pb
        .collection("customers")
        .authWithPassword(email.value, password.value);

    if (pb.authStore.isValid) {
        checkoutStep.value = "addresses";
    }
};
</script>