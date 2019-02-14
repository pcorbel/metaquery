<template>
  <v-app id="dataset">
    <v-container grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <mini-statistic
            icon="fa fa-database"
            :title="dataset.name"
            :sub-title="dataset.description"
            color="green"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <mini-statistic
            icon="fa fa-table"
            :title="tableCountFormatted"
            sub-title="Tables"
            color="blue"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg4 xl4>
          <mini-statistic
            icon="far fa-clock"
            :title="creationTimeFormatted"
            sub-title="Creation Time"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg4 xl4>
          <mini-statistic
            icon="fa fa-globe"
            :title="dataset.location"
            sub-title="Location"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg4 xl4>
          <mini-statistic
            icon="far fa-clock"
            :title="lastModificationTimeFormatted"
            sub-title="Last Modification Time"
            color="orange"
          />
        </v-flex>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <table-list
            :items="tables"
            :project-id="projectId"
            :dataset-id="datasetId"
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
          <permission-list :items="permissions" />
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
import PermissionList from "@/components/PermissionList";
import TableList from "@/components/TableList";
import { FormatterMixin } from "@/mixins/Formatter";
import { PermissionParserMixin } from "@/mixins/PermissionParser";

export default {
  components: {
    MiniStatistic,
    PermissionList,
    TableList
  },
  mixins: [FormatterMixin, PermissionParserMixin],
  props: {
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
    dataset: {},
    tables: [],
    permissions: []
  }),
  computed: {
    byteCountFormatted: function() {
      return this.byteFormat(this.dataset.byte_count);
    },
    rowCountFormatted: function() {
      return this.countFormat(this.dataset.row_count);
    },
    partitionCountFormatted: function() {
      return this.separatorFormat(this.dataset.partition_count);
    },
    fieldCountFormatted: function() {
      return this.separatorFormat(this.dataset.field_count);
    },
    datasetCountFormatted: function() {
      return String(this.dataset.dataset_count);
    },
    tableCountFormatted: function() {
      return String(this.dataset.table_count);
    },
    latestPartitionDateFormatted: function() {
      return this.dateFormat(this.dataset.latest_partition);
    },
    creationTimeFormatted: function() {
      return this.dateFormat(this.dataset.creation_time);
    },
    lastModificationTimeFormatted: function() {
      return this.dateFormat(this.dataset.last_modified_time);
    },
    latestPartitionRowCountFormatted: function() {
      return this.separatorFormat(this.dataset.latest_partition_row_count);
    }
  },
  created: function() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      let response = await this.$http.get(
        process.env.VUE_APP_ROOT_API +
          "/projects/" +
          this.$props.projectId +
          "/datasets/" +
          this.$props.datasetId
      );
      this.dataset = response.data ? response.data : {};
      this.permissions = this.parsePermission(response);
      this.$store.commit("setUpdatedAt", response.data._updated_at);
      this.$store.commit("setProject", this.$props.projectId);
      this.$store.commit("setDataset", this.$props.datasetId);
      this.$store.commit("setTable", undefined);

      response = await this.$http.get(
        process.env.VUE_APP_ROOT_API +
          "/projects/" +
          this.$props.projectId +
          "/datasets/" +
          this.$props.datasetId +
          "/tables"
      );
      this.tables = response.data ? response.data : [];
    }
  }
};
</script>
