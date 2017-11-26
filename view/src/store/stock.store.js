import axios from 'axios';

const stock = {
  state: {
    searchResult: [],
    savedStocks: [],
  },
  getters: {
    searchResult: state => state.searchResult,
    savedStocks: state => state.savedStocks,
  },
  mutations: {
    setSearchResult(state, data) {
      state.searchResult.splice(0, state.searchResult.length, ...data);
    },
    addSavedStock(state, data) {
      state.savedStocks.push(data);
    }

  },
  actions: {
    searchStock({commit}, keyword) {
      if(keyword === '') {
        commit('setSearchResult', []);
        return;
      }

      return axios.get(`/searchStock?keyword=${keyword}`, {
      }).then((result) => {
        if (result.status === 200) {
          result.data = result.data || [];
          commit('setSearchResult', result.data);
        }
      })
    },
    addSavedStock({commit}, data) {
      commit('addSavedStock', data);
    }
  }
};
export default stock;
