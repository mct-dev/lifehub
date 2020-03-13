<template>
  <div id="app">
    <CommandHandler />
    <router-view />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import CommandHandler from '@/components/CommandHandler.vue'

@Component({
  components: {
    CommandHandler
  }
})
export default class App extends Vue {}
</script>

<style lang="scss">
html, body {
  padding: 0;
  margin: 0;
  font-size: 10px;
}

// ---- UTILITY ---- //
.flex {
  &-center {
    display: flex;
    justify-content: center;
    &--column {
      display: flex;
      align-items: center;
      flex-direction: column;
    }
  }

}

.w-full {
  width: 100%;
}

.h-full {
  height: 100%;
}

// ---- APP ---- //
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;

  .home {
    &__container {
      height: 100vh;
    }
  }
}
</style>
