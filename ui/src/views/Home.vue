<template>
  <div class="home__container flex-center--column">
    <Info />
    <Input v-if="!isAcceptingCommand" />
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
    ...mapState(['isAcceptingCommand'])
  }
})
export default class Home extends Vue {}
</script>
