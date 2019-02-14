export const FormatterMixin = {
  methods: {
    byteFormat: function(input) {
      let fileSizeInBytes = input;
      let i = -1;
      let byteUnits = [" kB", " MB", " GB", " TB", "PB", "EB", "ZB", "YB"];
      do {
        fileSizeInBytes = fileSizeInBytes / 1024;
        i++;
      } while (fileSizeInBytes > 1024);
      return Math.max(fileSizeInBytes, 0.1).toFixed(1) + byteUnits[i];
    },
    countFormat: function(input) {
      let value = input;
      let min = 1e3;
      if (value >= min) {
        let units = [" k", " M", " B", " T"];
        let order = Math.floor(Math.log(value) / Math.log(1000));
        let unitname = units[order - 1];
        let num = Math.floor(value / 1000 ** order);
        return num + unitname;
      }
      return value;
    },
    separatorFormat: function(input) {
      let value = input;
      let separator = " ";
      value = String(value);
      let x = value.split(".");
      let x1 = x[0];
      let x2 = x.length > 1 ? "." + x[1] : "";
      let rgx = /(\d+)(\d{3})/;
      while (rgx.test(x1)) {
        x1 = x1.replace(rgx, "$1" + separator + "$2");
      }
      return x1 + x2;
    },
    dateFormat: function(input) {
      let d = new Date(input);
      let month = String(d.getMonth() + 1);
      let day = String(d.getDate());
      let year = d.getFullYear();

      if (month.length < 2) month = "0" + month;
      if (day.length < 2) day = "0" + day;

      return [year, month, day].join("-");
    }
  }
};
