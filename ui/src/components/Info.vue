<template>
  <div>
    <p>{{ debugInfo || '...' }}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState } from 'vuex'

@Component({
  computed: mapState(['debugInfo'])
})
export default class Info extends Vue {
}
</script>

<style lang="scss" scoped>
  div {
    width: 100%;
    margin-top: 40px;
    p {
      font-size: 2.5rem;
    }
  }
</style>
