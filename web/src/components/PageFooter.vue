<template>
  <v-footer height="auto" class="white pa-3 app--header">
    <span class="caption"> Last data refresh : {{ getUpdatedAt }} </span>
    <v-spacer />
    <a href="https://github.com/pcorbel/metaquery">
      <span class="caption mr-1"> Contribute on pcorbel/metaquery </span>
      <v-icon color="black" small> fab fa-github </v-icon>
    </a>
  </v-footer>
</template>

<script>
export default {
  computed: {
    getUpdatedAt() {
      return new Date(this.$store.getters.getUpdatedAt).toISOString();
    }
  }
};
</script>
