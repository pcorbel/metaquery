<template>
  <v-card>
    <v-toolbar card dense color="transparent">
      <v-toolbar-title> <h4>Tables</h4> </v-toolbar-title>
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
                :href="
                  '#/project/' +
                    projectId +
                    '/dataset/' +
                    datasetId +
                    '/table/' +
                    props.item.name
                "
              >
                {{ props.item.name }}
              </a>
            </td>
            <td class="text-xs-left">
              <v-chip label small color="green" text-color="white">
                {{ formatLastModificationTime(props.item.last_modified_time) }}
              </v-chip>
            </td>
            <td class="text-xs-left">{{ props.item.description }}</td>
            <td class="text-xs-left">
              <v-chip
                label
                small
                :color="
                  getColorByTimePartitionning(props.item.time_partitioning)
                "
                text-color="white"
              >
                {{ props.item.time_partitioning }}
              </v-chip>
            </td>
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
    },
    datasetId: {
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
        text: "Last Modified",
        value: "last_modified_time",
        sortable: false
      },
      {
        text: "Description",
        value: "description",
        sortable: false
      },
      {
        text: "Partitionning",
        value: "time_partitioning",
        sortable: false
      }
    ]
  }),
  methods: {
    getColorByTimePartitionning(timePartitionning) {
      if (timePartitionning === "DAY") {
        return "blue";
      } else {
        return "red";
      }
    },
    formatLastModificationTime(time) {
      let d = new Date(time);
      let month = String(d.getMonth() + 1);
      let day = String(d.getDate());
      let year = d.getFullYear();

      if (month.length < 2) month = "0" + month;
      if (day.length < 2) day = "0" + day;

      return [year, month, day].join("-");
    }
  }
};
</script>
