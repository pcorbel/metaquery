<template>
  <v-card>
    <v-toolbar card dense color="transparent">
      <v-toolbar-title> <h4>Datasets</h4> </v-toolbar-title>
    </v-toolbar>
    <v-divider />
    <v-card-text class="pa-0">
      <template>
        <v-data-table
          :headers="headers"
          :items="items"
          hide-actions
          class="elevation-0 table-striped"
        >
          <template slot="items" slot-scope="props">
            <td>
              <a
                :href="'#/project/' + projectId + '/dataset/' + props.item.name"
              >
                {{ props.item.name }}
              </a>
            </td>
            <td class="text-xs-left">
              <v-chip
                label
                small
                :color="getColorByLocation(props.item.location)"
                text-color="white"
              >
                {{ props.item.location }}
              </v-chip>
            </td>
            <td class="text-xs-left">{{ props.item.description }}</td>
          </template>
        </v-data-table>
      </template>
      <v-divider />
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      default: () => []
    },
    projectId: {
      type: String,
      default: ""
    }
  },
  data: () => ({
    headers: [
      {
        text: "ID",
        value: "id",
        sortable: false
      },
      {
        text: "Location",
        value: "location",
        sortable: false
      },
      {
        text: "Description",
        value: "description",
        sortable: false
      }
    ]
  }),
  methods: {
    getColorByLocation(location) {
      if (location === "EU") {
        return "blue";
      } else {
        return "red";
      }
    }
  }
};
</script>
