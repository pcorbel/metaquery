<template>
  <v-layout row class="align-center layout px-4 pt-4 app--page-header">
    <div class="page-header-left">
      <h3 class="pr-3">{{ title }}</h3>
    </div>
    <v-breadcrumbs divider=">">
      <v-breadcrumbs-item
        :href="'/#/project/' + this.$store.getters.getProject"
      >
        <v-icon larg> home </v-icon>
      </v-breadcrumbs-item>
      <v-breadcrumbs-item
        v-if="this.$store.getters.getProject"
        :href="getBreadcrumbUrl('project')"
      >
        {{ this.$store.getters.getProject }}
      </v-breadcrumbs-item>
      <v-breadcrumbs-item
        v-if="this.$store.getters.getDataset"
        :href="getBreadcrumbUrl('dataset')"
      >
        {{ this.$store.getters.getDataset }}
      </v-breadcrumbs-item>
      <v-breadcrumbs-item
        v-if="this.$store.getters.getTable"
        :href="getBreadcrumbUrl('table')"
      >
        {{ this.$store.getters.getTable }}
      </v-breadcrumbs-item>
    </v-breadcrumbs>
    <v-spacer />
    <div class="page-header-right">
      <a :href="getBigqueryUrl()"> <img src="/static/bigquery.png" /> </a>
    </div>
  </v-layout>
</template>

<script>
export default {
  data: () => ({
    title: ""
  }),
  computed: {
    breadcrumbs: function() {
      let breadcrumbs = ["Project"];
      if (this.$route.name === "Dataset") {
        breadcrumbs.push("Dataset");
      }
      if (this.$route.name === "Table") {
        breadcrumbs.push("Dataset");
        breadcrumbs.push("Table");
      }
      return breadcrumbs;
    }
  },
  methods: {
    getBreadcrumbUrl(stop) {
      let url = "";
      if (stop === "project") {
        url = "/#/project/" + this.$store.getters.getProject;
      }
      if (stop === "dataset") {
        url =
          "/#/project/" +
          this.$store.getters.getProject +
          "/dataset/" +
          this.$store.getters.getDataset;
      }
      if (stop === "table") {
        url =
          "/#/project/" +
          this.$store.getters.getProject +
          "/dataset/" +
          this.$store.getters.getDataset +
          "/table/" +
          this.$store.getters.getTable;
      }
      return url;
    },
    getBigqueryUrl() {
      let url = "https://console.cloud.google.com/bigquery?project=";
      let page = "project";
      if (this.$store.getters.getProject) {
        url +=
          this.$store.getters.getProject +
          "&p=" +
          this.$store.getters.getProject;
      }
      if (this.$store.getters.getDataset) {
        url += "&d=" + this.$store.getters.getDataset;
        page = "dataset";
      }
      if (this.$store.getters.getTable) {
        url += "&t=" + this.$store.getters.getTable;
        page = "table";
      }
      return url + "&page=" + page;
    }
  }
};
</script>
