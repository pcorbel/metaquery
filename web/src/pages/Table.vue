<template>
  <v-app id="table">
    <v-container grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm12 md12 lg12 xl12>
          <mini-statistic
            icon="fa fa-table"
            :title="table.name"
            :sub-title="table.description"
            color="green"
          />
        </v-flex>
        <v-flex v-if="table.time_partitioning === 'DAY'" xs12 sm12 md12 lg12 xl12>
          <v-card dark color="blue">
            <v-card-title>
              <div class="layout row ma-0">
                <div class="subheading">Partitions</div>
              </div>
            </v-card-title>
            <v-responsive class="white--text">
              <e-chart
                v-if="rowPerPartitionFormatted.length > 1"
                :path-option="[
                  ['dataset.source', rowPerPartitionFormatted],
                  ['grid.bottom', '10%'],
                  ['grid.top', '5%'],
                  ['xAxis.show', false],
                  ['yAxis.show', false],
                  ['series[0].type', 'bar'],
                  ['series[0].barGap', '-100%'],
                  ['series[0].itemStyle.normal.color', 'rgba(0,0,0,0.1)'],
                  ['series[0].barWidth', '50%'],
                  ['series[1].barWidth', '50%'],
                  ['series[1].type', 'bar'],
                  ['series[1].itemStyle.normal.color', color.shades.white]
                ]"
                height="200px"
                width="100%"
              />
            </v-responsive>
          </v-card>
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
            icon="far fa-file-code"
            :title="fieldCountFormatted"
            sub-title="Fields"
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
          <field-list :items="fields" />
        </v-flex>
        <v-flex xs12 sm12 md12 lg3 xl3>
          <mini-statistic
            icon="fa fa-globe"
            :title="'' + table.legacySQL"
            sub-title="LegacySQL"
            color="orange"
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
import FieldList from "@/components/FieldList";
import { FormatterMixin } from "@/mixins/Formatter";
import { PartitionParserMixin } from "@/mixins/PartitionParser";
import EChart from "@/components/chart/echart";
import Material from "vuetify/es5/util/colors";

export default {
  components: {
    MiniStatistic,
    FieldList,
    EChart
  },
  mixins: [FormatterMixin, PartitionParserMixin],
  props: {
    projectId: {
      type: String,
      default: ""
    },
    datasetId: {
      type: String,
      default: ""
    },
    tableId: {
      type: String,
      default: ""
    }
  },
  data: () => ({
    dataset: {},
    table: {},
    fields: [],
    partitions: [],
    color: Material
  }),
  computed: {
    byteCountFormatted: function() {
      return this.byteFormat(this.table.byte_count);
    },
    rowCountFormatted: function() {
      return this.countFormat(this.table.row_count);
    },
    partitionCountFormatted: function() {
      return this.separatorFormat(this.table.partition_count);
    },
    fieldCountFormatted: function() {
      return this.separatorFormat(this.table.field_count);
    },
    latestPartitionDateFormatted: function() {
      return this.dateFormat(this.table.latest_partition);
    },
    creationTimeFormatted: function() {
      return this.dateFormat(this.table.creation_time);
    },
    lastModificationTimeFormatted: function() {
      return this.dateFormat(this.table.last_modified_time);
    },
    latestPartitionRowCountFormatted: function() {
      return this.separatorFormat(this.table.latest_partition_row_count);
    },
    rowPerPartitionFormatted() {
      return this.parsePartitions(this.partitions);
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
          this.$props.datasetId +
          "/tables/" +
          this.$props.tableId
      );
      this.table = response.data ? response.data : {};
      this.$store.commit("setUpdatedAt", response.data._updated_at);
      this.$store.commit("setProject", this.$props.projectId);
      this.$store.commit("setDataset", this.$props.datasetId);
      this.$store.commit("setTable", this.$props.tableId);

      response = await this.$http.get(
        process.env.VUE_APP_ROOT_API +
          "/projects/" +
          this.$props.projectId +
          "/datasets/" +
          this.$props.datasetId +
          "/tables/" +
          this.$props.tableId +
          "/partitions"
      );
      this.partitions = response.data ? response.data : [];

      response = await this.$http.get(
        process.env.VUE_APP_ROOT_API +
          "/projects/" +
          this.$props.projectId +
          "/datasets/" +
          this.$props.datasetId +
          "/tables/" +
          this.$props.tableId +
          "/fields"
      );
      this.fields = response.data ? response.data : [];
    }
  }
};
</script>
