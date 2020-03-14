<template>
  <div class="home__container flex-center--column">
    <Info />
    <CmdInput :isOpen="showCommandInput" />
  </div>
</template>

<script>
import { Component, Vue } from 'vue-property-decorator'
import CmdInput from '@/components/CmdInput.vue'
import Info from '@/components/Info.vue'
import { mapState } from 'vuex'

@Component({
  components: {
    CmdInput,
    Info
  },
  computed: {
    ...mapState(['showCommandInput'])
  }
})
export default class Home extends Vue {}
</script>
