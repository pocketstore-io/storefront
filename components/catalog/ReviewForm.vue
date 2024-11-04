<template>
  <section class="form">
    <form @submit.prevent="save()" class="w-full bg-red-400">
      <div class="grid grid-cols-6 gap-3">
        <div class="col-span-6 md:col-span-2">
          customer preview <br>
          name <br>
          avatar <br>
        </div>
        <div class="col-span-6 md:col-span-4">
          date<br>
          <div class="grid grid-cols-6 gap-3">
            <div class="col-span-6 md:col-span-2">
              <input type="date" v-model="date" readonly class="input input-primary input-bordered w-full">
            </div>
            <div class="col-span-6 md:col-span-4">
              <input type="text" placeholder="Title" class="input input-primary input-bordered w-full">
            </div>
            <div class="col-span-6 md:col-span-4">
              <textarea name="" id="" class="textarea w-full"></textarea>
            </div>
            <div class="col-span-6 md:col-span-2">
              <input type="date" v-model="date" readonly class="input input-primary input-bordered w-full">
            </div>
          </div>
          description<br>
          headline<br>
          stars<br>
        </div>
      </div>
    </form>
  </section>
</template>

<script lang="ts" setup>

const date = computed(()=> {
  return new Date().toLocaleDateString('en');
})

const save = () => {
  console.log('saved');
}
</script>
