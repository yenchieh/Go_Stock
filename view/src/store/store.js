import Vue from 'vue';
import Vuex from 'vuex';
import stock from "./stock.store";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    stock,
  }
});

export default store;
