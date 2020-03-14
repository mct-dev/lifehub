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
export default class CmdInput extends Vue {
  @Prop() isOpen!: boolean;
  cmd = ''

  handleKeydown = (ev: KeyboardEvent) => {
    const alphaNumChars = ' abcdefghijklmnopqrstuvwxyz0123456789'
    const isAlphaNumeric = alphaNumChars.includes(ev.key)

    if (isAlphaNumeric) {
      this.cmd += ev.key
      return
    }

    switch (ev.key) {
      case 'Escape':
        this.$store.commit(SET_SHOW_COMMAND_INPUT, false)
        break
      case 'Enter':
        this.$store.commit(SET_SHOW_COMMAND_INPUT, false);
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
  background: rgb(255, 255, 255);
  border-radius: 10px;
  display: flex;
  height: 100vh;
  justify-content: center;
  left: 0;
  position: fixed;
  top: 0;
  width: 100vw;
  opacity: 0;
  transition: opacity .3s ease-out;

  input {
    background: transparent;
    border-radius: 5px;
    border: 3px solid #eee;
    font-size: 2.5rem;
    outline: none;
    padding: 30px 20px;
    width: 80%;
  }

  &.open {
    opacity: 1;
    transition: opacity .2s ease-in;
  }
}
</style>
