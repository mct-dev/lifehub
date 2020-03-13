<template>
  <div></div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import {
  SET_USER_COMMAND_KEYS,
  SET_IS_ACCEPTING_COMMAND
} from '../store/mutation-types'

@Component
export default class KeyHandler extends Vue {
  keyBuffer: string[] = []
  lastKeyDownTime = -1

  handleKeydown = (ev: KeyboardEvent) => {
    const allowedFirstKeys = ['Shift', ' ']

    if (ev.key === 'Escape') {
      this.keyBuffer = []
      this.$store.commit(SET_USER_COMMAND_KEYS, this.keyBuffer)
      this.$store.commit(SET_IS_ACCEPTING_COMMAND, true)
      return
    }

    if (!this.keyBuffer.length) {
      if (allowedFirstKeys.indexOf(ev.key) === -1) return
      this.keyBuffer.push(ev.key)
      return
    }

    if (ev.key === this.keyBuffer[this.keyBuffer.length - 1]) {
      return
    }

    this.keyBuffer.push(ev.key)
  }

  handleKeyup = () => {
    switch (this.keyBuffer.length) {
      case 0:
        this.$store.commit(SET_USER_COMMAND_KEYS, this.keyBuffer)
        return
      case 1:
        return
      case 2:
        this.$store.commit(SET_USER_COMMAND_KEYS, this.keyBuffer)
        this.keyBuffer = []
        return
      case 3:
        this.$store.commit(SET_USER_COMMAND_KEYS, this.keyBuffer)
        return
      default:
        this.keyBuffer = []
        break
    }
  }

  mounted () {
    document.addEventListener('keydown', this.handleKeydown)
    document.addEventListener('keyup', this.handleKeyup)
  }

  beforeDestroy () {
    document.removeEventListener('keydown', this.handleKeydown)
    document.removeEventListener('keyup', this.handleKeyup)
  }
}
</script>
