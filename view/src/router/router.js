import VueRouter from 'vue-router';
import Stock from '../pages/stock.vue';


const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Stock
    },
  ]
});

export default router;
