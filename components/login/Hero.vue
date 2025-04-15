<template>
  <div class="overflow-hidden bg-white py-6 rounded-lg grid grid-cols-6">
    <LoginCardHero2 />
    <div class="col-span-6 md:col-span-3 px-8">
      <LoginCardRegisterForm v-if="form == 'register'" />
      <LoginCardLoginForm v-else />
    </div>
  </div>
</template>
