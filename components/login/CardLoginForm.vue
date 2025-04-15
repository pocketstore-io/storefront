<template>
    <section class="card bg-base-300 px-3 py-3">
        <div class="form">
            <h2 class="font-bold text-lg mb-3">{{ $t('form.login-via') }}</h2>
            <SocialLogin />
            <div class="divider divider-primary font-bold my-9">{{ $t('form.or') }}</div>
            <h2 class="font-bold text-lg mb-3">{{ $t('form.form') }}</h2>
            <FormRegisterAndLogin />
        </div>
    </section>
</template>

<script setup lang="ts">
import FormRegisterAndLogin from './FormRegisterAndLogin.vue';
import SocialLogin from './SocialLogin.vue';
</script>