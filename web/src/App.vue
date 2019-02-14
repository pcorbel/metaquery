<template>
  <div id="appRoot">
    <template v-if="!$route.meta.public">
      <v-app
        id="inspire"
        class="app"
      >
        <app-toolbar class="app--toolbar" />
        <v-content>
          <!-- Page Header -->
          <page-header v-if="$route.meta.breadcrumb" />
          <div class="page-wrapper">
            <router-view :key="$route.fullPath" />
          </div>
          <!-- App Footer -->
          <page-footer />
        </v-content>
        <!-- Go to top -->
        <app-fab />
      </v-app>
    </template>
    <template v-else>
      <transition>
        <keep-alive> <router-view /> </keep-alive>
      </transition>
    </template>
  </div>
</template>

<script>
  import AppToolbar from '@/components/AppToolbar'
  import AppFab from '@/components/AppFab'
  import PageHeader from '@/components/PageHeader'
  import PageFooter from '@/components/PageFooter'

  export default {
    components: {
      AppToolbar,
      AppFab,
      PageHeader,
      PageFooter
    },
    data: () => ({
      expanded: true
    }),

    computed: {},

    created () {
      window.getApp = this
    },
    methods: {}
  }
</script>
