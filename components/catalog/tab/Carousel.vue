<template>
  <div class="carousel w-full">
    <div id="slide1" class="carousel-item relative w-full">
      <img src="https://place-hold.it/1920x1080" class="w-full" />
      <div
        class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between"
      >
        <a href="#slide4" class="btn btn-circle">❮</a>
        <a href="#slide2" class="btn btn-circle">❯</a>
      </div>
    </div>
    <div id="slide2" class="carousel-item relative w-full">
      <img src="https://place-hold.it/1920x1080" class="w-full" />
      <div
        class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between"
      >
        <a href="#slide1" class="btn btn-circle">❮</a>
        <a href="#slide3" class="btn btn-circle">❯</a>
      </div>
    </div>
    <div id="slide3" class="carousel-item relative w-full">
      <img src="https://place-hold.it/1920x1080" class="w-full" />
      <div
        class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between"
      >
        <a href="#slide2" class="btn btn-circle">❮</a>
        <a href="#slide4" class="btn btn-circle">❯</a>
      </div>
    </div>
    <div id="slide4" class="carousel-item relative w-full">
      <img src="https://place-hold.it/1920x1080" class="w-full" />
      <div
        class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between"
      >
        <a href="#slide3" class="btn btn-circle">❮</a>
        <a href="#slide1" class="btn btn-circle">❯</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style>
</style>