import Vue from 'vue'
import VueMaterial from 'vue-material'
import VueRouter from 'vue-router';
import 'vue-material/dist/vue-material.min.css'

import App from './App.vue'
import router from './router/router';

Vue.use(VueMaterial);
Vue.use(VueRouter);

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
