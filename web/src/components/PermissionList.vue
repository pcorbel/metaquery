<template>
  <v-card>
    <v-toolbar card dense color="transparent">
      <v-toolbar-title> <h4>Permissions</h4> </v-toolbar-title>
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
              <v-icon
                dark
                :color="getColorByType(props.item.type)"
                :title="props.item.type"
              >
                {{ getAvatarByType(props.item.type) }}
              </v-icon>
            </td>
            <td>
              <v-chip
                label
                small
                :color="getColorByType(props.item.type)"
                text-color="white"
              >
                {{ props.item.type }}
              </v-chip>
            </td>
            <td>{{ props.item.name }}</td>
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
    }
  },
  data: () => ({
    headers: [
      {
        text: "",
        value: "",
        sortable: false
      },
      {
        text: "Type",
        value: "type",
        sortable: false
      },
      {
        text: "Name",
        value: "name",
        sortable: false
      }
    ]
  }),
  methods: {
    getAvatarByType: function(type) {
      if (type === "owner") {
        return "fab fa-superpowers";
      } else if (type === "writer") {
        return "fas fa-pencil-alt";
      } else if (type === "reader") {
        return "fas fa-book-reader";
      }
    },
    getColorByType: function(type) {
      if (type === "owner") {
        return "blue";
      } else if (type === "writer") {
        return "green";
      } else if (type === "reader") {
        return "red";
      }
    }
  }
};
</script>
