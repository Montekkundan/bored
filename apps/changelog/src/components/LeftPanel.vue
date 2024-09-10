<template>
    <div class="bg-blue-900 text-white h-screen p-8 flex flex-col">
      <h1 class="text-4xl font-bold mb-4">Commit</h1>
      <h2 class="text-3xl mb-4">hi hello</h2>
      <p class="mb-8">
        all bored commits
      </p>
      <div class="mt-auto">
        <input type="email" placeholder="Email address" class="bg-blue-800 text-white p-2 rounded w-full mb-4">
        <button class="bg-blue-700 text-white p-2 rounded w-full">Get updates →</button>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  
  export default defineComponent({
    name: 'LeftPanel',
  })
  </script>