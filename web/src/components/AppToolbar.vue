<template>
  <v-toolbar color="primary" fixed dark app>
    <v-toolbar-title class="ml-0 pl-3"> Metaquery </v-toolbar-title>
    <v-spacer />
    <v-autocomplete
      v-model="model"
      :items="items"
      :loading="isLoading"
      :search-input.sync="search"
      color="white"
      item-text="name"
      item-value="name"
      label="Search"
      placeholder="Start typing to search"
      prepend-icon="search"
      flat
      solo-inverted
      return-object
      hide-no-data
      hide-selected
      @change="onChange()"
    >
      <template slot="item" slot-scope="{ item }">
        <v-list-tile-avatar v-if="item.type == 'project'" color="white">
          <v-icon> fa-cloud </v-icon>
        </v-list-tile-avatar>
        <v-list-tile-avatar v-if="item.type == 'dataset'" color="white">
          <v-icon> fa-database </v-icon>
        </v-list-tile-avatar>
        <v-list-tile-avatar v-if="item.type == 'table'" color="white">
          <v-icon> fa-table </v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title v-text="item.name" />
          <v-list-tile-sub-title v-text="item.full_id" />
        </v-list-tile-content>
      </template>
    </v-autocomplete>
    <v-spacer />
  </v-toolbar>
</template>

<script>
export default {
  data() {
    return {
      descriptionLimit: 60,
      entries: [],
      isLoading: false,
      model: null,
      search: null,
      count: 0
    };
  },

  computed: {
    fields() {
      if (!this.model) return [];
      return Object.keys(this.model).map(key => {
        return { key, value: this.model[key] || "n/a" };
      });
    },
    items() {
      return this.entries.map(entry => {
        const fullID = entry.full_id;
        const parsedFullId = fullID.split(/\:|\./);
        const name = parsedFullId[parsedFullId.length - 1];
        let path = "/project/" + parsedFullId[0];
        if (entry.type === "dataset") {
          path += "/dataset/" + parsedFullId[1];
        }
        if (entry.type === "table") {
          path += "/dataset/" + parsedFullId[1] + "/table/" + parsedFullId[2];
        }
        return Object.assign({}, entry, { fullID, name, path });
      });
    }
  },

  watch: {
    search() {
      this.fetchData();
    }
  },

  created: function() {
    this.fetchData();
  },

  methods: {
    async fetchData() {
      // Items have already been loaded
      if (this.items.length > 0) return;

      // Items have already been requested
      if (this.isLoading) return;

      // Load items
      this.isLoading = true;
      let response = await this.$http.get(
        process.env.VUE_APP_ROOT_API + "/entries"
      );
      this.count = response.data.length;
      this.entries = response.data;
      this.isLoading = false;
    },
    onChange() {
      this.$router.push({ path: this.model.path });
    }
  }
};
</script>

<style>
.v-input__slot {
  margin-bottom: 0px;
}
</style>
