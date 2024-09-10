<template>
    <div class="bg-gray-900 text-white min-h-screen p-8">
      <h2 class="text-2xl font-bold mb-4">Changelog</h2>
      <div v-for="commit in commits" :key="commit.sha" class="mb-8">
        <CommitItem :commit="commit" />
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, onMounted, ref } from 'vue'
  import axios from 'axios'
  import CommitItem from '../components/CommitItem.vue'
  
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
    name: 'AppChangelog',
    components: {
      CommitItem,
    },
    setup() {
      const commits = ref<Commit[]>([])
  
      const fetchCommits = async () => {
        try {
          const response = await axios.get('https://api.github.com/repos/montekkundan/bored/commits')
          commits.value = response.data
        } catch (error) {
          console.error('Error fetching commits:', error)
        }
      }
  
      onMounted(fetchCommits)
  
      return {
        commits,
      }
    },
  })
  </script>