<template>
    <div class="bg-gray-800 rounded-lg p-4 relative">
      <div class="absolute -left-4 top-1/2 transform -translate-y-1/2 w-8 h-8 bg-gray-700 rounded-full flex items-center justify-center text-xs">
        {{ formatDate(commit.commit.author.date) }}
      </div>
      <h3 class="text-xl font-semibold mb-2">{{ commit.commit.message }}</h3>
      <p class="text-gray-400">Commit SHA: {{ commit.sha.substring(0, 7) }}</p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, type PropType } from 'vue'
  
  interface Commit {
    sha: string
    commit: {
      message: string
      author: {
        date: string
      }
    }
  }
  
  export default defineComponent({
    name: 'CommitItem',
    props: {
      commit: {
        type: Object as PropType<Commit>,
        required: true,
      },
    },
    methods: {
      formatDate(dateString: string): string {
        const date = new Date(dateString)
        return `${date.getDate()}/${date.getMonth() + 1}`
      },
    },
  })
  </script>