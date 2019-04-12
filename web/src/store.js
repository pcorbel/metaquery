import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    updated_at: '1970-01-01T00:00:00Z',
    project: '',
    dataset: '',
    table: ''
  },
  getters: {
    getUpdatedAt: state => {
      return state.updated_at
    },
    getProject: state => {
      return state.project
    },
    getDataset: state => {
      return state.dataset
    },
    getTable: state => {
      return state.table
    }
  },
  mutations: {
    setUpdatedAt (state, updatedAt) {
      state.updated_at = updatedAt
    },
    setProject (state, project) {
      state.project = project
    },
    setDataset (state, dataset) {
      state.dataset = dataset
    },
    setTable (state, table) {
      state.table = table
    }
  }
})
