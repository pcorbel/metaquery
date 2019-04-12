export const PermissionParserMixin = {
  methods: {
    parsePermission: function(input) {
      let response = input;
      let permissions = [];
      if (response.data.owners != null) {
        for (let i = 0; i < response.data.owners.length; i++) {
          permissions.push({
            type: "owner",
            name: response.data.owners[i]
          });
        }
      }
      if (response.data.writers != null) {
        for (let i = 0; i < response.data.writers.length; i++) {
          permissions.push({
            type: "writer",
            name: response.data.writers[i]
          });
        }
      }
      if (response.data.readers != null) {
        for (let i = 0; i < response.data.readers.length; i++) {
          permissions.push({
            type: "reader",
            name: response.data.readers[i]
          });
        }
      }
      return permissions;
    }
  }
};
