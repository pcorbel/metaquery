<template>
  <v-card>
    <v-toolbar card dense color="transparent">
      <v-toolbar-title> <h4>Fields</h4> </v-toolbar-title>
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
            <td class="text-xs-left">{{ props.item.name }}</td>
            <td class="text-xs-left">{{ props.item.type }}</td>
            <td class="text-xs-left">{{ props.item.mode }}</td>
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
    }
  },
  data: () => ({
    headers: [
      {
        text: "Name",
        value: "name",
        sortable: false
      },
      {
        text: "Type",
        value: "type",
        sortable: false
      },
      {
        text: "Mode",
        value: "mode",
        sortable: false
      },
      {
        text: "Description",
        value: "description",
        sortable: false
      }
    ]
  })
};
</script>
