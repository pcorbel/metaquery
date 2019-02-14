export const PartitionParserMixin = {
  methods: {
    parsePartitions: function(input) {
      if (!input.length) {
        return [{ count: 0 }];
      }
      let max = 0;
      for (var i = 0; i < input.length; i++) {
        if (input[i].count > max) {
          max = input[i].count;
        }
      }
      let data = input.map(input => {
        return {
          partitiontime: input.partitiontime,
          max: max,
          count: input.count
        };
      });
      data.sort(function(a, b) {
        var keyA = new Date(a.partitiontime);
        var keyB = new Date(b.partitiontime);
        if (keyA < keyB) return -1;
        if (keyA > keyB) return 1;
        return 0;
      });
      // Hack to display only one bar on the graph
      if (data.length === 1) {
        data.push({
          partitiontime: data[0].partitiontime,
          max: data[0].max,
          count: data[0].partitiontime + "" + data[0].count
        });
      }
      return data;
    }
  }
};
