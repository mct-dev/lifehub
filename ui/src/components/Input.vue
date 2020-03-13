<template>
  <div class="input">
    <input
      @keydown="handleKeydown"
      v-model="msg"
      placeholder="Enter command..."
      autofocus
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { SET_IS_ACCEPTING_COMMAND } from '../store/mutation-types'

@Component({})
export default class Input extends Vue {
  msg = 'wow'

  handleKeydown = (ev: KeyboardEvent) => {
    console.log('input key', ev.key)
    if (ev.key === 'Escape') {
      this.$store.commit(SET_IS_ACCEPTING_COMMAND, true)
    }
  }
}
</script>
<style scoped lang="scss">
.input {
  margin-top: 30vh;
  width: 80%;

  input {
    padding: 20px;
    width: 100%;
    outline: none;
    border: 1px solid #eee;
    border-radius: 2px;
    font-size: 2.5rem;
  }
}
</style>
