<template>
  <v-app id="project">
    <v-container grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <mini-statistic
            icon="fa fa-cloud"
            :title="project.full_id"
            sub-title="Project"
            color="green"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg6 xl6>
          <mini-statistic
            icon="fa fa-database"
            :title="datasetCountFormatted"
            sub-title="Datasets"
            color="blue"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg6 xl6>
          <mini-statistic
            icon="fa fa-table"
            :title="tableCountFormatted"
            sub-title="Tables"
            color="blue"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg3 xl3>
          <mini-statistic
            icon="fa fa-bars"
            :title="rowCountFormatted"
            sub-title="Rows"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg3 xl3>
          <mini-statistic
            icon="fab fa-buromobelexperte"
            :title="byteCountFormatted"
            sub-title="Size"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg3 xl3>
          <mini-statistic
            icon="far fa-clock"
            :title="partitionCountFormatted"
            sub-title="Partitions"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg3 xl3>
          <mini-statistic
            icon="far fa-file-code"
            :title="fieldCountFormatted"
            sub-title="Fields"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <dataset-list :items="datasets" :project-id="this.$props.projectId" />
        </v-flex>
        <v-flex xs12 sm12 md12 lg6 xl6>
          <mini-statistic
            icon="far fa-clock"
            :title="latestPartitionDateFormatted"
            sub-title="Latest Partition"
            color="red"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg6 xl6>
          <mini-statistic
            icon="fa fa-bars"
            :title="latestPartitionRowCountFormatted"
            sub-title="Latest Partition Row Count"
            color="red"
          />
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script>
import MiniStatistic from "@/components/MiniStatistic";
import DatasetList from "@/components/DatasetList";
import { FormatterMixin } from "@/mixins/Formatter";

export default {
  components: {
    MiniStatistic,
    DatasetList
  },
  mixins: [FormatterMixin],
  props: {
    projectId: {
      type: String,
      default: ""
    }
  },
  data: () => ({
    project: {},
    datasets: []
  }),
  computed: {
    byteCountFormatted: function() {
      return this.byteFormat(this.project.byte_count);
    },
    rowCountFormatted: function() {
      return this.countFormat(this.project.row_count);
    },
    partitionCountFormatted: function() {
      return this.separatorFormat(this.project.partition_count);
    },
    fieldCountFormatted: function() {
      return this.separatorFormat(this.project.field_count);
    },
    datasetCountFormatted: function() {
      return String(this.project.dataset_count);
    },
    tableCountFormatted: function() {
      return String(this.project.table_count);
    },
    latestPartitionDateFormatted: function() {
      return this.dateFormat(this.project.latest_partition);
    },
    latestPartitionRowCountFormatted: function() {
      return this.separatorFormat(this.project.latest_partition_row_count);
    }
  },
  created: function() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      let response = await this.$http.get(
        process.env.VUE_APP_ROOT_API + "/projects/" + this.$props.projectId
      );
      this.project = response.data ? response.data : {};
      this.$store.commit("setUpdatedAt", response.data._updated_at);
      this.$store.commit("setProject", this.$props.projectId);
      this.$store.commit("setDataset", undefined);
      this.$store.commit("setTable", undefined);

      response = await this.$http.get(
        process.env.VUE_APP_ROOT_API +
          "/projects/" +
          this.$props.projectId +
          "/datasets"
      );
      this.datasets = response.data ? response.data : [];
    }
  }
};
</script>
