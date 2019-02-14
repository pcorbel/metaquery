export default [
  {
    path: "/",
    meta: {},
    name: "Root",
    component: () =>
      import(/* webpackChunkName: "routes" */
      /* webpackMode: "lazy" */
      `@/pages/Root.vue`)
  },
  {
    path: "/project/:projectId",
    meta: { breadcrumb: true },
    name: "Project",
    props: true,
    component: () =>
      import(/* webpackChunkName: "routes" */
      /* webpackMode: "lazy" */
      `@/pages/Project.vue`)
  },
  {
    path: "/project/:projectId/dataset/:datasetId",
    meta: { breadcrumb: true },
    name: "Dataset",
    props: true,
    component: () =>
      import(/* webpackChunkName: "routes" */
      /* webpackMode: "lazy" */
      `@/pages/Dataset.vue`)
  },
  {
    path: "/project/:projectId/dataset/:datasetId/table/:tableId",
    meta: { breadcrumb: true },
    name: "Table",
    props: true,
    component: () =>
      import(/* webpackChunkName: "routes" */
      /* webpackMode: "lazy" */
      `@/pages/Table.vue`)
  }
];
