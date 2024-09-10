<template>
  <div class="flex h-screen">
    <div class="w-1/2">
      <LeftPanel />
    </div>
    <div class="w-1/2 overflow-y-auto">
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LeftPanel from './components/LeftPanel.vue'

export default defineComponent({
  name: 'App',
  components: {
    LeftPanel,
  },
})
</script>