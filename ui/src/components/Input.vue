<template>
  <div :class="`input ${isOpen ? 'open' : ''}`">
    <input
      ref="commandInput"
      @keydown="handleKeydown"
      v-model="cmd"
      placeholder="Enter command..."
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { SET_SHOW_COMMAND_INPUT } from '../store/mutation-types'

@Component({})
export default class Input extends Vue {
  @Prop() isOpen!: boolean;
  cmd = ''

  handleKeydown = (ev: KeyboardEvent) => {
    const alphaChars = ' abcdefghijklmnopqrstuvwxyz0123456789'
    const isAlphaNumeric = alphaChars.includes(ev.key)

    if (isAlphaNumeric) {
      this.cmd += ev.key
      return
    }

    switch (ev.key) {
      case 'Escape':
        this.$store.commit(SET_SHOW_COMMAND_INPUT, false)
        break
      case 'Enter':
        console.log('command is: ', this.cmd);
        (this.$refs.commandInput as HTMLInputElement).value = ''
        this.cmd = ''
        break
      default:
        break
    }
  }

  beforeUpdate () {
    if (this.isOpen) {
      (this.$refs.commandInput as HTMLInputElement).focus()
    } else {
      (this.$refs.commandInput as HTMLInputElement).blur()
      this.cmd = ''
    }
  }
}
</script>
<style scoped lang="scss">
.input {
  align-items: center;
  background: rgb(243, 243, 243);
  border-radius: 10px;
  display: flex;
  height: 100vh;
  justify-content: center;
  left: 0;
  position: fixed;
  top: 0;
  width: 100vw;
  opacity: 0;

  input {
    padding: 20px;
    width: 80%;
    outline: none;
    border: 1px solid #eee;
    border-radius: 2px;
    font-size: 2.5rem;
  }

  &.open {
    opacity: 1;
    transition: opacity .3s ease-in;
  }
}
</style>
