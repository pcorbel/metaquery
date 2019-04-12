<script>
export default {
  created: function() {
    this.redirect();
  },
  methods: {
    async redirect() {
      let response = await this.$http.get(
        process.env.VUE_APP_ROOT_API + "/projects"
      );
      this.$router.push({
        path: "/project/" + response.data[0].full_id
      });
    }
  }
};
</script>
