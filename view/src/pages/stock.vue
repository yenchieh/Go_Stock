<template>
  <div class="stockPage">
    <h1>Stock Page should have these component</h1>
    <div class="container">
        <md-autocomplete @md-selected="selected" ref="searchText" v-model="stockKeyword" :md-options="searchResult" @md-changed="search" md-dense>
          <label>Search Stock</label>
          <template slot="md-autocomplete-item" slot-scope="{ item }">{{ item.name }}</template>
        </md-autocomplete>

      <div class="stockCard-container">
        <md-card v-for="stock in savedStocks" class="stockCard">
          <md-card-header>
            <div class="md-title">{{ stock.name }}({{ stock.symbol }})</div>
            <div class="md-subhead">{{ stock.industry }}</div>
          </md-card-header>

          <md-card-content class="stockDetail">
            <div>
              <div>IPO Year</div>
              <div>{{ stock.ipo_year }}</div>
            </div>
            <div>
              <div>Last Sale</div>
              <div>{{ stock.last_sale }}</div>
            </div>
            <div>
              <div>Market Cap</div>
              <div>{{ stock.market_cap }}</div>
            </div>
            <div>
              <div>Sector</div>
              <div>{{ stock.sector }}</div>
            </div>
          </md-card-content>

          <md-card-actions>
            <md-button>Add To List</md-button>
            <md-button>Remove</md-button>
          </md-card-actions>
        </md-card>
      </div>

    </div>

  </div>
</template>

<script>
  import {mapGetters} from 'vuex';
  export default {
    name: 'stock',
    data () {
      return {
        stockKeyword: '',
        requestTimeout: null,
      }
    },

    computed: {
      ...mapGetters([
        'searchResult',
        'savedStocks',
      ])
    },

    methods: {
      search() {
        if(this.requestTimeout){
          clearTimeout(this.requestTimeout);
          this.requestTimeout = null;
        }

        this.requestTimeout = setTimeout(() => {
          this.$store.dispatch('searchStock', this.$refs.searchText.searchTerm);
          this.requestTimeout = null;
        }, 500);
      },
      selected(item) {
        this.$store.dispatch('addSavedStock', item).then(() => {
          this.$refs.searchText.searchTerm = "";
          this.$store.dispatch('searchStock', "");
        });

      }
    }
  }
</script>

<style lang="scss" scoped>
  .stockPage {
    .container {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }

    .md-field {
      max-width: 300px;
    }

    .stockCard {
      width: 300px;
      margin: 15px;

      &-container {
        display: flex;
        max-width: 100%;
        flex-wrap: wrap;
      }
    }

    .stockDetail {
      > div {
        display: flex;
        justify-content: space-between;
        border-bottom: 1px solid rgba(0,0,0,.3);
        margin-bottom: 10px;
        padding-bottom: 5px;
      }
    }
  }
</style>
