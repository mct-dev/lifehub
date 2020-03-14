<template>
  <div class="home__container flex-center--column">
    <Info />
    <Input :isOpen="showCommandInput" />
  </div>
</template>

<script>
import { Component, Vue } from 'vue-property-decorator'
import Input from '@/components/Input.vue'
import Info from '@/components/Info.vue'
import { mapState } from 'vuex'

@Component({
  components: {
    Input,
    Info
  },
  computed: {
    ...mapState(['showCommandInput'])
  }
})
export default class Home extends Vue {}
</script>
